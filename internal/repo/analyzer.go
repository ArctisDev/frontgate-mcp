package repo

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ArctisDev/frontgate/internal/parser"
	"github.com/ArctisDev/frontgate/internal/types"
)

func Analyze(ctx context.Context, projectPath, task string) (types.RepoContext, error) {
	var out types.RepoContext
	if projectPath == "" {
		return out, errors.New("project_path is required")
	}

	absPath, err := filepath.Abs(projectPath)
	if err != nil {
		return out, fmt.Errorf("resolve project path: %w", err)
	}
	info, err := os.Stat(absPath)
	if err != nil {
		return out, fmt.Errorf("stat project path: %w", err)
	}
	if !info.IsDir() {
		return out, fmt.Errorf("project_path must be a directory: %s", absPath)
	}

	out.ProjectPath = absPath
	out.DesignTokens = map[string]string{}
	out.DetectedScripts = map[string]string{}
	out.PrimitiveCatalog = types.PrimitiveCatalog{}

	pkg, err := parser.LoadPackageJSON(absPath)
	if err == nil {
		deps := mergeDeps(pkg.Dependencies, pkg.DevDependencies)
		out.DetectedDependencies = sortedKeys(deps)
		out.DetectedScripts = pkg.Scripts
		out.DetectedFramework = detectFramework(deps)
		out.UsesTailwind = hasAny(deps, "tailwindcss")
		out.UsesShadcn = hasAny(deps, "class-variance-authority", "tailwind-merge", "lucide-react")
		out.UsesRadix = hasPrefix(deps, "@radix-ui/")
		out.UsesDaisyUI = hasAny(deps, "daisyui")
	}

	appDir := filepath.Join(absPath, "app")
	if statIsDir(appDir) {
		out.UsesNextAppRouter = true
		out.RelevantDirectories = append(out.RelevantDirectories, rel(absPath, appDir))
	}
	for _, dir := range []string{"components", "components/ui", "lib", "styles"} {
		full := filepath.Join(absPath, dir)
		if statIsDir(full) {
			out.RelevantDirectories = append(out.RelevantDirectories, rel(absPath, full))
		}
	}

	spacingCounts := map[string]int{}
	paddingCounts := map[string]int{}
	marginCounts := map[string]int{}
	gapCounts := map[string]int{}
	baseComponents := map[string]struct{}{}
	layoutPatterns := map[string]struct{}{}
	shellPatterns := map[string]struct{}{}
	var arbitrary []string
	patternInventory := types.PatternInventory{}
	routeIndex := map[string]*types.RouteInfo{}

	err = filepath.WalkDir(absPath, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if d.IsDir() {
			name := d.Name()
			if name == "node_modules" || name == ".next" || name == ".git" {
				return filepath.SkipDir
			}
			return nil
		}

		relPath := rel(absPath, path)
		lower := strings.ToLower(relPath)
		if isRelevantFile(lower) {
			out.RelevantFiles = append(out.RelevantFiles, relPath)
		}

		if strings.HasSuffix(lower, "components.json") {
			out.UsesShadcn = true
		}

		if strings.HasSuffix(lower, ".css") {
			body, _ := os.ReadFile(path)
			for k, v := range parser.ExtractCSSVariables(string(body)) {
				out.DesignTokens[k] = v
			}
		}

		if !isSourceFile(lower) {
			return nil
		}

		bodyBytes, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		body := string(bodyBytes)

		for _, token := range parser.ExtractSpacingClasses(body) {
			spacingCounts[token]++
			switch {
			case strings.HasPrefix(token, "p"):
				paddingCounts[token]++
			case strings.HasPrefix(token, "m"):
				marginCounts[token]++
			case strings.HasPrefix(token, "gap"), strings.HasPrefix(token, "space-"):
				gapCounts[token]++
			}
			if strings.Contains(token, "[") {
				arbitrary = append(arbitrary, token)
			}
		}

		if strings.Contains(lower, "components/ui/") {
			componentName := strings.TrimSuffix(filepath.Base(lower), filepath.Ext(lower))
			baseComponents[componentName] = struct{}{}
			categorizePrimitive(componentName, &out.PrimitiveCatalog)
		}

		categorizePatterns(lower, body, &patternInventory, layoutPatterns)
		categorizeShell(lower, body, shellPatterns)
		updateRouteInfo(lower, body, routeIndex)
		return nil
	})
	if err != nil {
		return out, err
	}

	out.RelevantFiles = dedupeAndSort(out.RelevantFiles)
	out.RelevantDirectories = dedupeAndSort(out.RelevantDirectories)
	out.BaseComponents = sortedKeysFromSet(baseComponents)
	out.LayoutPatterns = sortedKeysFromSet(layoutPatterns)
	out.ShellPatterns = sortedKeysFromSet(shellPatterns)
	out.PatternInventory = patternInventory
	out.RouteInventory = flattenRoutes(routeIndex)
	out.ArbitrarySpacingValues = dedupeAndSort(arbitrary)
	out.SpacingPatterns = types.SpacingPatterns{
		MostCommonPadding: topCounts(paddingCounts, 6),
		MostCommonGap:     topCounts(gapCounts, 6),
		MostCommonMargin:  topCounts(marginCounts, 6),
		HasArbitrary:      len(out.ArbitrarySpacingValues) > 0,
	}

	out.TargetFileHints = suggestTargets(task, out)
	out.VisualRisks = detectRisks(out, spacingCounts)
	out.VisualSystemSummary = buildSummary(out)

	if out.DetectedFramework == "" && out.UsesNextAppRouter {
		out.DetectedFramework = "Next.js"
	}
	if out.DetectedFramework == "" {
		out.DetectedFramework = "Unknown"
	}

	return out, nil
}

func detectFramework(deps map[string]string) string {
	switch {
	case hasAny(deps, "next"):
		return "Next.js"
	case hasAny(deps, "react"):
		return "React"
	default:
		return ""
	}
}

func categorizePatterns(path, body string, inventory *types.PatternInventory, layouts map[string]struct{}) {
	name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	switch {
	case strings.Contains(path, "sidebar"):
		inventory.Sidebars = append(inventory.Sidebars, name)
		layouts["sidebar"] = struct{}{}
	case strings.Contains(path, "table"):
		inventory.Tables = append(inventory.Tables, name)
		layouts["table"] = struct{}{}
	case strings.Contains(path, "modal") || strings.Contains(path, "dialog"):
		inventory.Modals = append(inventory.Modals, name)
		layouts["modal"] = struct{}{}
	case strings.Contains(path, "form"):
		inventory.Forms = append(inventory.Forms, name)
		layouts["form"] = struct{}{}
	case strings.Contains(path, "tabs"):
		inventory.Tabs = append(inventory.Tabs, name)
		layouts["tabs"] = struct{}{}
	case strings.Contains(path, "card") || strings.Contains(body, "<Card"):
		inventory.Cards = append(inventory.Cards, name)
		layouts["cards"] = struct{}{}
	}
	if strings.Contains(body, "<aside") || strings.Contains(body, "Sidebar") {
		layouts["sidebar"] = struct{}{}
	}
	if strings.Contains(body, "<table") || strings.Contains(body, "<Table") {
		layouts["table"] = struct{}{}
	}
	if strings.Contains(body, "grid-cols") {
		layouts["grid"] = struct{}{}
	}
	if strings.Contains(body, "max-w-") {
		layouts["constrained-width"] = struct{}{}
	}
	if strings.Contains(body, "sticky top-") || strings.Contains(body, "position:sticky") {
		layouts["sticky-shell"] = struct{}{}
	}
	if strings.Contains(body, "overflow-y-auto") {
		layouts["scroll-region"] = struct{}{}
	}
}

func categorizePrimitive(componentName string, catalog *types.PrimitiveCatalog) {
	lower := strings.ToLower(componentName)
	switch {
	case strings.Contains(lower, "button"):
		catalog.Buttons = append(catalog.Buttons, componentName)
	case strings.Contains(lower, "input"), strings.Contains(lower, "select"), strings.Contains(lower, "textarea"), strings.Contains(lower, "checkbox"), strings.Contains(lower, "radio"), strings.Contains(lower, "switch"):
		catalog.Inputs = append(catalog.Inputs, componentName)
	case strings.Contains(lower, "dialog"), strings.Contains(lower, "modal"), strings.Contains(lower, "drawer"), strings.Contains(lower, "sheet"), strings.Contains(lower, "popover"):
		catalog.Dialogs = append(catalog.Dialogs, componentName)
	case strings.Contains(lower, "card"):
		catalog.Cards = append(catalog.Cards, componentName)
	case strings.Contains(lower, "table"), strings.Contains(lower, "datatable"):
		catalog.Tables = append(catalog.Tables, componentName)
	case strings.Contains(lower, "tab"):
		catalog.Tabs = append(catalog.Tabs, componentName)
	case strings.Contains(lower, "sidebar"):
		catalog.Sidebars = append(catalog.Sidebars, componentName)
	case strings.Contains(lower, "form"):
		catalog.Forms = append(catalog.Forms, componentName)
	case strings.Contains(lower, "nav"), strings.Contains(lower, "breadcrumb"), strings.Contains(lower, "menu"):
		catalog.Navigation = append(catalog.Navigation, componentName)
	case strings.Contains(lower, "alert"), strings.Contains(lower, "toast"), strings.Contains(lower, "badge"), strings.Contains(lower, "skeleton"), strings.Contains(lower, "empty"):
		catalog.Feedback = append(catalog.Feedback, componentName)
	}
}

func categorizeShell(path, body string, shells map[string]struct{}) {
	if strings.HasSuffix(path, "app/layout.tsx") {
		shells["root-layout"] = struct{}{}
	}
	if strings.Contains(path, "/layout.tsx") && strings.Contains(path, "app/") && !strings.HasSuffix(path, "app/layout.tsx") {
		shells["nested-layout"] = struct{}{}
	}
	if strings.Contains(body, "<Sidebar") || strings.Contains(body, "<aside") {
		shells["sidebar-shell"] = struct{}{}
	}
	if strings.Contains(body, "max-w-screen") || strings.Contains(body, "container mx-auto") {
		shells["centered-shell"] = struct{}{}
	}
	if strings.Contains(body, "min-h-screen") {
		shells["full-height-shell"] = struct{}{}
	}
	if strings.Contains(body, "sticky top-0") || strings.Contains(body, "<header") {
		shells["sticky-header"] = struct{}{}
	}
}

func updateRouteInfo(path, body string, routes map[string]*types.RouteInfo) {
	if !strings.Contains(path, "app/") {
		return
	}
	if !(strings.HasSuffix(path, "/page.tsx") || strings.HasSuffix(path, "/layout.tsx") || strings.HasSuffix(path, "/loading.tsx") || strings.HasSuffix(path, "/error.tsx")) {
		return
	}

	routeKey := routeFromPath(path)
	item := routes[routeKey]
	if item == nil {
		item = &types.RouteInfo{Route: routeKey}
		routes[routeKey] = item
	}

	switch {
	case strings.HasSuffix(path, "/page.tsx"):
		item.PageFile = path
	case strings.HasSuffix(path, "/layout.tsx"):
		item.LayoutFile = path
	case strings.HasSuffix(path, "/loading.tsx"):
		item.LoadingFile = path
	case strings.HasSuffix(path, "/error.tsx"):
		item.ErrorFile = path
	}

	if strings.Contains(body, `"use client"`) || strings.Contains(body, `'use client'`) {
		item.UsesClient = true
	}
	if strings.Contains(body, "<form") || strings.Contains(body, "<Form") {
		item.HasForm = true
	}
	if strings.Contains(body, "<table") || strings.Contains(body, "<Table") {
		item.HasTable = true
	}
	if strings.Contains(body, "<Sidebar") || strings.Contains(body, "<aside") {
		item.HasSidebar = true
	}
}

func detectRisks(ctx types.RepoContext, spacingCounts map[string]int) []string {
	var risks []string
	if len(ctx.ArbitrarySpacingValues) >= 4 {
		risks = append(risks, "Many arbitrary spacing values detected; likely inconsistent rhythm.")
	}
	if len(spacingCounts) > 28 {
		risks = append(risks, "Large variety of spacing utilities detected; padding and gap scale may be drifting.")
	}
	if len(ctx.DesignTokens) == 0 {
		risks = append(risks, "No CSS design tokens detected in scanned stylesheets.")
	}
	if len(ctx.BaseComponents) < 6 && ctx.UsesShadcn {
		risks = append(risks, "shadcn stack detected but few reusable primitives were found.")
	}
	if len(ctx.PatternInventory.Sidebars) > 2 {
		risks = append(risks, "Multiple sidebar implementations found; visual shell may be fragmented.")
	}
	if len(ctx.PatternInventory.Cards) > 6 {
		risks = append(risks, "Many card variants found; the product may be over-boxed or template-like.")
	}
	if len(ctx.RouteInventory) == 0 && ctx.UsesNextAppRouter {
		risks = append(risks, "Next.js App Router is detected but no route inventory could be assembled.")
	}
	if len(ctx.PrimitiveCatalog.Buttons) == 0 && len(ctx.BaseComponents) > 0 {
		risks = append(risks, "Base component library detected, but button primitives were not identified clearly.")
	}
	if len(ctx.ShellPatterns) == 0 {
		risks = append(risks, "No clear shell patterns were detected, which can make page-level refactors less disciplined.")
	}
	if !ctx.UsesTailwind {
		risks = append(risks, "Tailwind was not clearly detected, which is outside the MVP target stack.")
	}
	return dedupeAndSort(risks)
}

func buildSummary(ctx types.RepoContext) string {
	stack := []string{ctx.DetectedFramework}
	if ctx.UsesTailwind {
		stack = append(stack, "Tailwind")
	}
	if ctx.UsesShadcn {
		stack = append(stack, "shadcn/ui-style primitives")
	}
	if ctx.UsesRadix {
		stack = append(stack, "Radix")
	}
	return fmt.Sprintf(
		"Stack detected: %s. Reusable primitives: %d. Layout patterns: %s. Shell patterns: %s. Main visual risks: %s.",
		strings.Join(stack, ", "),
		len(ctx.BaseComponents),
		strings.Join(ctx.LayoutPatterns, ", "),
		strings.Join(ctx.ShellPatterns, ", "),
		strings.Join(ctx.VisualRisks, "; "),
	)
}

func suggestTargets(task string, ctx types.RepoContext) []string {
	var hints []string
	task = strings.ToLower(task)
	for _, file := range ctx.RelevantFiles {
		lower := strings.ToLower(file)
		switch {
		case strings.Contains(task, "deploy") && strings.Contains(lower, "deploy"):
			hints = append(hints, file)
		case strings.Contains(task, "dashboard") && strings.Contains(lower, "dashboard"):
			hints = append(hints, file)
		case strings.Contains(task, "sidebar") && strings.Contains(lower, "sidebar"):
			hints = append(hints, file)
		case strings.Contains(task, "table") && strings.Contains(lower, "table"):
			hints = append(hints, file)
		case strings.Contains(task, "form") && strings.Contains(lower, "form"):
			hints = append(hints, file)
		}
	}
	if len(hints) == 0 {
		hints = append(hints, ctx.PatternInventory.Sidebars...)
		hints = append(hints, ctx.PatternInventory.Cards...)
		for _, route := range ctx.RouteInventory {
			if route.PageFile != "" {
				hints = append(hints, route.PageFile)
			}
		}
	}
	return dedupeAndSort(hints)
}

func isRelevantFile(path string) bool {
	names := []string{
		"package.json", "components.json", "tailwind.config", "globals.css",
		"layout.tsx", "page.tsx", "utils.ts", "theme.ts",
	}
	for _, name := range names {
		if strings.Contains(path, name) {
			return true
		}
	}
	return strings.Contains(path, "components/ui/")
}

func isSourceFile(path string) bool {
	switch {
	case strings.HasSuffix(path, ".tsx"), strings.HasSuffix(path, ".ts"), strings.HasSuffix(path, ".jsx"), strings.HasSuffix(path, ".js"), strings.HasSuffix(path, ".css"):
		return true
	default:
		return false
	}
}

func topCounts(counts map[string]int, limit int) []types.MetricCount {
	items := make([]types.MetricCount, 0, len(counts))
	for name, count := range counts {
		items = append(items, types.MetricCount{Name: name, Count: count})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Count == items[j].Count {
			return items[i].Name < items[j].Name
		}
		return items[i].Count > items[j].Count
	})
	if len(items) > limit {
		items = items[:limit]
	}
	return items
}

func routeFromPath(path string) string {
	parts := strings.Split(filepath.ToSlash(path), "/")
	appIndex := -1
	for i, part := range parts {
		if part == "app" {
			appIndex = i
			break
		}
	}
	if appIndex == -1 || appIndex+1 >= len(parts) {
		return "/"
	}
	segments := parts[appIndex+1 : len(parts)-1]
	var route []string
	for _, seg := range segments {
		if seg == "" {
			continue
		}
		if strings.HasPrefix(seg, "(") && strings.HasSuffix(seg, ")") {
			continue
		}
		route = append(route, seg)
	}
	if len(route) == 0 {
		return "/"
	}
	return "/" + strings.Join(route, "/")
}

func flattenRoutes(routeIndex map[string]*types.RouteInfo) []types.RouteInfo {
	keys := make([]string, 0, len(routeIndex))
	for key := range routeIndex {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	out := make([]types.RouteInfo, 0, len(keys))
	for _, key := range keys {
		out = append(out, *routeIndex[key])
	}
	return out
}

func hasAny(deps map[string]string, names ...string) bool {
	for _, name := range names {
		if _, ok := deps[name]; ok {
			return true
		}
	}
	return false
}

func hasPrefix(deps map[string]string, prefix string) bool {
	for name := range deps {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	return false
}

func mergeDeps(a, b map[string]string) map[string]string {
	out := make(map[string]string, len(a)+len(b))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] = v
	}
	return out
}

func rel(root, path string) string {
	value, err := filepath.Rel(root, path)
	if err != nil {
		return path
	}
	return filepath.ToSlash(value)
}

func statIsDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func sortedKeys(m map[string]string) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func sortedKeysFromSet(m map[string]struct{}) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func dedupeAndSort(items []string) []string {
	set := make(map[string]struct{}, len(items))
	for _, item := range items {
		if item == "" {
			continue
		}
		set[item] = struct{}{}
	}
	out := make([]string, 0, len(set))
	for item := range set {
		out = append(out, item)
	}
	sort.Strings(out)
	return out
}
