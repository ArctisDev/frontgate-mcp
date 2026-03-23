package rules

import (
	"regexp"
	"sort"
	"strings"

	"github.com/ArctisDev/frontgate/internal/types"
)

var arbitrarySpacingRE = regexp.MustCompile(`\b(?:p|px|py|pt|pr|pb|pl|m|mx|my|mt|mr|mb|ml|gap|space-x|space-y)-\[[^\]]+\]`)
var genericWordsRE = regexp.MustCompile(`(?i)\b(insight|metric|analytics|overview|summary|widget|quick action|engagement)\b`)
var cardRE = regexp.MustCompile(`(?i)\bcard\b`)
var newComponentRE = regexp.MustCompile(`^\+\+\+|^diff|^\+.*(?:export\s+(?:default\s+)?function|const\s+[A-Z][A-Za-z0-9]+)\b`)
var classNameRE = regexp.MustCompile(`className\s*=\s*["'` + "`" + `]([^"'` + "`" + `]+)["'` + "`" + `]`)
var addedDivRE = regexp.MustCompile(`^\+\s*<div\b|^\+<div\b`)
var hardcodedColorRE = regexp.MustCompile(`\b(?:bg|text|border|ring|fill|stroke)-(?:slate|gray|zinc|neutral|stone|red|orange|amber|yellow|lime|green|emerald|teal|cyan|sky|blue|indigo|violet|purple|fuchsia|pink|rose)-[0-9]{2,3}\b`)
var headingRE = regexp.MustCompile(`(?i)<h[1-4]\b|text-(?:2xl|3xl|4xl)|font-(?:bold|semibold)`)
var sidebarWidthRE = regexp.MustCompile(`\b(?:w|basis|min-w|max-w)-(72|80|96|\[[0-9]{3,}px\])\b`)
var rawPrimitiveRE = regexp.MustCompile(`(?m)^\+.*<(button|input|select|textarea|dialog)\b`)

func ScoreFromIssues(report *types.CritiqueReport) {
	if report.Scores.SpacingConsistency == 0 {
		report.Scores.SpacingConsistency = 70
	}
	if report.Scores.TypographicHierarchy == 0 {
		report.Scores.TypographicHierarchy = 72
	}
	if report.Scores.LayoutBalance == 0 {
		report.Scores.LayoutBalance = 72
	}
	if report.Scores.ComponentReuse == 0 {
		report.Scores.ComponentReuse = 70
	}
	if report.Scores.TokenAdherence == 0 {
		report.Scores.TokenAdherence = 70
	}
	if report.Scores.DensityControl == 0 {
		report.Scores.DensityControl = 70
	}
	if report.Scores.VisualNoise == 0 {
		report.Scores.VisualNoise = 70
	}
	if report.Scores.TemplateLikeness == 0 {
		report.Scores.TemplateLikeness = 70
	}
	if report.Scores.UXClarity == 0 {
		report.Scores.UXClarity = 72
	}
	if report.Scores.ProductAlignment == 0 {
		report.Scores.ProductAlignment = 72
	}

	for _, issue := range report.Problems {
		penalty := severityPenalty(issue.Severity)
		switch issue.Category {
		case "spacing":
			report.Scores.SpacingConsistency -= penalty
			report.Scores.TokenAdherence -= penalty / 2
		case "hierarchy":
			report.Scores.TypographicHierarchy -= penalty
			report.Scores.UXClarity -= penalty / 2
		case "layout":
			report.Scores.LayoutBalance -= penalty
			report.Scores.DensityControl -= penalty / 2
		case "reuse":
			report.Scores.ComponentReuse -= penalty
		case "noise":
			report.Scores.VisualNoise -= penalty
			report.Scores.TemplateLikeness -= penalty / 2
		case "product":
			report.Scores.ProductAlignment -= penalty
			report.Scores.TemplateLikeness -= penalty / 2
		case "density":
			report.Scores.DensityControl -= penalty
			report.Scores.LayoutBalance -= penalty / 2
		}
	}

	clampScores(&report.Scores)
}

func FindArbitrarySpacing(diff string) []string {
	return uniqueStrings(arbitrarySpacingRE.FindAllString(diff, -1))
}

func CountGenericSaaSLanguage(text string) int {
	return len(genericWordsRE.FindAllString(text, -1))
}

func CountCardMentions(diff string) int {
	return len(cardRE.FindAllString(diff, -1))
}

func CountNewTopLevelComponents(diff string) int {
	lines := strings.Split(diff, "\n")
	count := 0
	for _, line := range lines {
		if strings.HasPrefix(line, "+") && newComponentRE.MatchString(line) {
			count++
		}
	}
	return count
}

func ExtractClassNames(diff string) []string {
	matches := classNameRE.FindAllStringSubmatch(diff, -1)
	var out []string
	for _, m := range matches {
		out = append(out, strings.Fields(m[1])...)
	}
	return uniqueStrings(out)
}

func CountAddedWrapperDivs(diff string) int {
	return len(addedDivRE.FindAllString(diff, -1))
}

func FindHardcodedPaletteClasses(diff string) []string {
	return uniqueStrings(hardcodedColorRE.FindAllString(diff, -1))
}

func HasClearHeadingStructure(diff string) bool {
	return headingRE.MatchString(diff)
}

func FindOversizedSidebarClasses(diff string) []string {
	return uniqueStrings(sidebarWidthRE.FindAllString(diff, -1))
}

func CountRawPrimitiveAdditions(diff string) int {
	return len(rawPrimitiveRE.FindAllString(diff, -1))
}

func CountExistingComponentReuse(diff string, components []string) int {
	count := 0
	for _, component := range components {
		if component == "" {
			continue
		}
		re := regexp.MustCompile(`(?i)<` + regexp.QuoteMeta(component) + `\b`)
		count += len(re.FindAllString(diff, -1))
	}
	return count
}

func uniqueStrings(items []string) []string {
	set := make(map[string]struct{}, len(items))
	for _, item := range items {
		set[item] = struct{}{}
	}
	out := make([]string, 0, len(set))
	for item := range set {
		out = append(out, item)
	}
	sort.Strings(out)
	return out
}

func severityPenalty(severity string) int {
	switch strings.ToLower(severity) {
	case "high":
		return 20
	case "medium":
		return 12
	default:
		return 7
	}
}

func clampScores(scores *types.CritiqueScores) {
	values := []*int{
		&scores.SpacingConsistency,
		&scores.TypographicHierarchy,
		&scores.LayoutBalance,
		&scores.ComponentReuse,
		&scores.TokenAdherence,
		&scores.DensityControl,
		&scores.VisualNoise,
		&scores.TemplateLikeness,
		&scores.UXClarity,
		&scores.ProductAlignment,
	}
	for _, value := range values {
		if *value < 0 {
			*value = 0
		}
		if *value > 100 {
			*value = 100
		}
	}
}
