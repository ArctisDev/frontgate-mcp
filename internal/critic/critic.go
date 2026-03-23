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
			Category: "spacing",
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
