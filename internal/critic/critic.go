package critic

import (
	"context"
	"fmt"
	"strings"

	"github.com/ArctisDev/frontgate/internal/playwright"
	"github.com/ArctisDev/frontgate/internal/rules"
	"github.com/ArctisDev/frontgate/internal/types"
)

func Critique(ctx context.Context, in types.CritiqueUIOutputInput) (types.CritiqueReport, error) {
	report := types.CritiqueReport{}
	diff := in.Diff
	if strings.TrimSpace(diff) == "" {
		return report, fmt.Errorf("diff is required")
	}

	if in.RenderRequest != nil {
		pwReport, err := playwright.Render(ctx, *in.RenderRequest)
		if err != nil {
			report.Problems = append(report.Problems, types.CritiqueIssue{
				Category: "layout",
				Severity: "low",
				Detail:   fmt.Sprintf("Playwright render could not run: %v", err),
				Action:   "Verify the local app is running and the playwright package is installed in the chosen working directory.",
			})
		} else {
			report.PlaywrightReport = pwReport
		}
	}

	arbitrarySpacing := rules.FindArbitrarySpacing(diff)
	if len(arbitrarySpacing) > 0 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "spacing",
			Severity: severityByCount(len(arbitrarySpacing)),
			Detail:   fmt.Sprintf("Arbitrary spacing utilities added in diff: %s", strings.Join(arbitrarySpacing, ", ")),
			Action:   "Replace arbitrary spacing with the project's existing Tailwind spacing scale unless there is a hard product reason.",
		})
	}

	if genericCount := rules.CountGenericSaaSLanguage(diff + "\n" + in.Task); genericCount >= 3 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "product",
			Severity: "medium",
			Detail:   "The change leans on generic SaaS language patterns instead of product-specific communication.",
			Action:   "Rewrite labels and support copy to reflect the actual workflow, state and user intent of the product.",
		})
	}

	if cards := rules.CountCardMentions(diff); cards >= 6 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "noise",
			Severity: "medium",
			Detail:   "The diff introduces many card references, which often correlates with over-boxed layouts.",
			Action:   "Flatten structure where possible and reserve card shells for true grouping or elevation changes.",
		})
	}

	if components := rules.CountNewTopLevelComponents(diff); components >= 3 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "reuse",
			Severity: "medium",
			Detail:   "Multiple new top-level component definitions were introduced in a single UI change.",
			Action:   "Check whether these abstractions should be replaced by existing primitives or existing screen sections.",
		})
	}

	classNames := rules.ExtractClassNames(diff)
	if len(classNames) >= 40 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "density",
			Severity: "medium",
			Detail:   "The diff carries a high number of utility classes, which may indicate wrapper-heavy composition.",
			Action:   "Collapse redundant wrappers and move repeated patterns into shared primitives where they already exist.",
		})
	}

	if wrappers := rules.CountAddedWrapperDivs(diff); wrappers >= 12 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "density",
			Severity: "medium",
			Detail:   fmt.Sprintf("Many wrapper divs were added in the diff (%d), which suggests structure inflation.", wrappers),
			Action:   "Reduce wrapper nesting and let existing layout primitives carry spacing and alignment where possible.",
		})
	}

	if !rules.HasClearHeadingStructure(diff) {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "hierarchy",
			Severity: "medium",
			Detail:   "The diff does not show a strong heading or emphasis structure for the target screen.",
			Action:   "Clarify the page hierarchy with stronger heading levels, supporting text and action grouping.",
		})
	}

	if hardcodedPalette := rules.FindHardcodedPaletteClasses(diff); len(hardcodedPalette) >= 3 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "token_adherence",
			Severity: "medium",
			Detail:   fmt.Sprintf("The diff introduces several hardcoded color utility classes: %s", strings.Join(hardcodedPalette, ", ")),
			Action:   "Prefer semantic tokens and existing color primitives over raw palette utilities when aligning with the product system.",
		})
	}

	if sidebarWidths := rules.FindOversizedSidebarClasses(diff); len(sidebarWidths) > 0 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "layout",
			Severity: "high",
			Detail:   fmt.Sprintf("The diff adds large fixed sidebar/container widths: %s", strings.Join(sidebarWidths, ", ")),
			Action:   "Avoid oversized side rails and reclaim space for the primary task area unless the product shell truly requires it.",
		})
	}

	rawPrimitives := rules.CountRawPrimitiveAdditions(diff)
	reusedComponents := rules.CountExistingComponentReuse(diff, in.Context.BaseComponents)
	if rawPrimitives >= 2 && reusedComponents == 0 && len(in.Context.BaseComponents) > 0 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "reuse",
			Severity: "medium",
			Detail:   "The diff adds raw form/action primitives without showing reuse of the repository's existing component primitives.",
			Action:   "Swap raw button, input and dialog markup for existing reusable primitives when they already exist in the repo.",
		})
	}

	if genericCTAs := rules.CountGenericSaaSCTAs(diff); len(genericCTAs) >= 2 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "product",
			Severity: "medium",
			Detail:   fmt.Sprintf("Generic SaaS CTA copy detected: %s. These phrases appear on every SaaS landing page and signal a lack of product identity.", strings.Join(genericCTAs, ", ")),
			Action:   "Replace generic CTAs with copy that reflects the product's actual value proposition and the user's specific intent at that point in the page.",
		})
	}

	if cdnRefs := rules.FindCDNFontReferences(diff); len(cdnRefs) > 0 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "token_adherence",
			Severity: "high",
			Detail:   fmt.Sprintf("External CDN font references detected: %s. These can fail to load (404) and cause FOUT/FOIT, degrading the visual experience.", strings.Join(cdnRefs, ", ")),
			Action:   "Use next/font with Google Fonts or self-hosted font files. Never rely on external CDN links for critical typography.",
		})
	}

	if gradients := rules.FindGradientOveruse(diff); len(gradients) >= 8 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "visual",
			Severity: "medium",
			Detail:   fmt.Sprintf("Excessive gradient usage detected (%d gradient-related classes). Overuse of gradients makes the design feel busy and undirected.", len(gradients)),
			Action:   "Use gradients sparingly for depth and accent. One strong gradient per section is enough. Let neutrals and solid colors do most of the work.",
		})
	}

	if shadows := rules.CountShadowUsage(diff); shadows >= 10 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "visual",
			Severity: "low",
			Detail:   fmt.Sprintf("Many shadow utilities detected (%d). Excessive shadows flatten the hierarchy and make everything feel elevated.", shadows),
			Action:   "Reserve shadows for true depth cues: elevated panels, dropdowns, modals. Remove shadows from cards and containers that don't need elevation.",
		})
	}

	if borderRadius := rules.CountBorderRadiusVariety(diff); borderRadius >= 6 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "visual",
			Severity: "low",
			Detail:   fmt.Sprintf("Too many border-radius variants detected (%d). Inconsistent rounding weakens the visual system.", borderRadius),
			Action:   "Standardize on 2-3 border-radius values (e.g., sm for inputs, lg for cards, full for avatars). Don't mix every available variant.",
		})
	}

	critiquePlaywright(&report)
	report.CorrectiveActions = collectActions(report.Problems)
	report.Summary = buildSummary(report)
	rules.ScoreFromIssues(&report)
	return report, nil
}

func critiquePlaywright(report *types.CritiqueReport) {
	pw := report.PlaywrightReport
	if pw == nil {
		return
	}

	if pw.OversizedSidebarPX > 360 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "layout",
			Severity: "high",
			Detail:   fmt.Sprintf("Sidebar width looks oversized at %dpx.", pw.OversizedSidebarPX),
			Action:   "Reduce the sidebar footprint or collapse non-essential chrome to preserve content space.",
		})
	}
	if pw.OverflowingElements > 0 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "layout",
			Severity: "high",
			Detail:   fmt.Sprintf("%d overflowing elements detected in the rendered page.", pw.OverflowingElements),
			Action:   "Fix width constraints, flex shrink behavior or breakpoints to eliminate horizontal overflow.",
		})
	}
	if pw.RepeatedWrappers >= 18 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "density",
			Severity: "medium",
			Detail:   fmt.Sprintf("High repeated wrapper count detected (%d).", pw.RepeatedWrappers),
			Action:   "Simplify nesting and remove presentation-only wrappers when an existing primitive can carry the layout.",
		})
	}
	if len(pw.UniqueFontSizes) >= 8 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "hierarchy",
			Severity: "medium",
			Detail:   fmt.Sprintf("Too many unique computed font sizes detected: %s", strings.Join(pw.UniqueFontSizes, ", ")),
			Action:   "Tighten typographic scale so headers, body and meta text use a smaller controlled set of sizes.",
		})
	}
	if len(pw.UniquePaddings) >= 12 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "spacing",
			Severity: "medium",
			Detail:   "Rendered DOM shows a wide spread of unique paddings, suggesting weak spacing consistency.",
			Action:   "Unify panel, section and card paddings around a tighter spacing scale.",
		})
	}

	if pw.MaxDepth > 20 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "density",
			Severity: "medium",
			Detail:   fmt.Sprintf("DOM tree is excessively deep (max depth: %d). Deep nesting increases render cost and makes layout harder to reason about.", pw.MaxDepth),
			Action:   "Flatten the component tree where possible and remove unnecessary intermediate wrappers.",
		})
	}
	if pw.FixedElements > 5 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "layout",
			Severity: "medium",
			Detail:   fmt.Sprintf("High number of fixed-position elements detected (%d). This may indicate overlay sprawl or stacking issues.", pw.FixedElements),
			Action:   "Reduce fixed-position elements. Use sticky for headers/sidebars and reserve fixed for true global overlays like toasts.",
		})
	}
	if pw.AbsoluteElements > 15 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "layout",
			Severity: "medium",
			Detail:   fmt.Sprintf("Excessive absolute-positioned elements (%d). Heavy use of absolute positioning suggests layout hacks instead of proper flex/grid flow.", pw.AbsoluteElements),
			Action:   "Replace absolute positioning with flex/grid layout where possible. Reserve absolute for tooltips, dropdowns and popovers.",
		})
	}
	if pw.ScrollContainers > 6 {
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "layout",
			Severity: "low",
			Detail:   fmt.Sprintf("Many scroll containers detected (%d). Nested scroll areas create a poor user experience.", pw.ScrollContainers),
			Action:   "Consolidate scroll areas. Use a single main scroll for the page and only one sidebar/panel scroll if needed.",
		})
	}

	for _, note := range pw.Notes {
		if strings.Contains(strings.ToLower(note), "oversized sidebar") {
			continue
		}
		if strings.Contains(strings.ToLower(note), "overflow") {
			continue
		}
		report.Problems = append(report.Problems, types.CritiqueIssue{
			Category: "layout",
			Severity: "low",
			Detail:   fmt.Sprintf("Playwright detected: %s", note),
			Action:   "Investigate the rendered DOM signal and determine if it indicates a structural issue.",
		})
	}
}

func collectActions(problems []types.CritiqueIssue) []string {
	var actions []string
	seen := map[string]struct{}{}
	for _, item := range problems {
		if _, ok := seen[item.Action]; ok {
			continue
		}
		seen[item.Action] = struct{}{}
		actions = append(actions, item.Action)
	}
	return actions
}

func buildSummary(report types.CritiqueReport) string {
	if len(report.Problems) == 0 {
		return "No critical UI issues were detected by the current deterministic heuristics."
	}
	var high int
	for _, item := range report.Problems {
		if item.Severity == "high" {
			high++
		}
	}
	return fmt.Sprintf(
		"Detected %d issues (%d high severity). Main problems center on spacing, hierarchy, layout efficiency and product alignment.",
		len(report.Problems),
		high,
	)
}

func severityByCount(count int) string {
	switch {
	case count >= 5:
		return "high"
	case count >= 2:
		return "medium"
	default:
		return "low"
	}
}
