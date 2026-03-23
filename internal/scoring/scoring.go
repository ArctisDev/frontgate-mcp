package scoring

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ArctisDev/frontgate/internal/types"
)

var defaultWeights = map[string]int{
	"spacing_consistency":   12,
	"typographic_hierarchy": 10,
	"layout_balance":        12,
	"component_reuse":       10,
	"token_adherence":       10,
	"density_control":       10,
	"visual_noise":          10,
	"template_likeness":     8,
	"ux_clarity":            10,
	"product_alignment":     8,
}

func Score(in types.ScoreUIQualityInput) types.QualityScore {
	weights := mergeWeights(in.Weights)
	axes := scoreAxes(in.CritiqueReport)
	totalWeight := 0
	totalScore := 0
	for _, axis := range axes {
		weight := weights[axis.Name]
		totalWeight += weight
		totalScore += axis.Score * weight
	}
	if totalWeight > 0 {
		totalScore = totalScore / totalWeight
	}
	return types.QualityScore{
		TotalScore:           totalScore,
		Axes:                 axes,
		RecommendedThreshold: recommendThreshold(totalScore, in.CritiqueReport),
	}
}

func Gate(in types.GateUIChangeInput) types.GateResult {
	threshold := in.Threshold
	if threshold == 0 {
		threshold = in.Score.RecommendedThreshold
		if threshold == 0 {
			threshold = 75
		}
	}

	approved := in.Score.TotalScore >= threshold && !hasBlockingIssues(in.Critique)
	var justification string
	if approved {
		justification = fmt.Sprintf("Approved: total score %d meets threshold %d and no blocking critique issue remains.", in.Score.TotalScore, threshold)
	} else {
		justification = fmt.Sprintf("Rejected: total score %d is below threshold %d or blocking issues remain.", in.Score.TotalScore, threshold)
	}

	changes := append([]string{}, in.Critique.CorrectiveActions...)
	if len(changes) == 0 && !approved {
		changes = append(changes, "Improve the lowest scoring axes before re-running the gate.")
	}

	return types.GateResult{
		Approved:        approved,
		Threshold:       threshold,
		Justification:   justification,
		RequiredChanges: changes,
	}
}

func scoreAxes(report types.CritiqueReport) []types.ScoreAxis {
	axes := []types.ScoreAxis{
		axis("spacing_consistency", report.Scores.SpacingConsistency, "Consistency of padding, gap and rhythm across the layout."),
		axis("typographic_hierarchy", report.Scores.TypographicHierarchy, "Clarity of headings, support text and information emphasis."),
		axis("layout_balance", report.Scores.LayoutBalance, "Efficiency of screen real estate, panel weight and shell footprint."),
		axis("component_reuse", report.Scores.ComponentReuse, "Reuse of existing primitives and avoidance of ad hoc markup."),
		axis("token_adherence", report.Scores.TokenAdherence, "Respect for current project tokens, scales and visual language."),
		axis("density_control", report.Scores.DensityControl, "Ability to show useful information without clutter or wasted space."),
		axis("visual_noise", report.Scores.VisualNoise, "Absence of decorative containers, redundant blocks and friction."),
		axis("template_likeness", report.Scores.TemplateLikeness, "Distance from generic SaaS template tropes."),
		axis("ux_clarity", report.Scores.UXClarity, "How clearly the UI communicates actions, states and priorities."),
		axis("product_alignment", report.Scores.ProductAlignment, "Fit between the UI treatment and the product's actual job."),
	}

	sort.SliceStable(axes, func(i, j int) bool {
		return axes[i].Name < axes[j].Name
	})
	return axes
}

func axis(name string, score int, explanation string) types.ScoreAxis {
	return types.ScoreAxis{Name: name, Score: score, Explanation: explanation}
}

func mergeWeights(custom map[string]int) map[string]int {
	out := make(map[string]int, len(defaultWeights))
	for k, v := range defaultWeights {
		out[k] = v
	}
	for k, v := range custom {
		if v > 0 {
			out[k] = v
		}
	}
	return out
}

func recommendThreshold(total int, report types.CritiqueReport) int {
	if hasBlockingIssues(report) {
		return 80
	}
	if total >= 85 {
		return 80
	}
	return 75
}

func hasBlockingIssues(report types.CritiqueReport) bool {
	for _, item := range report.Problems {
		if strings.EqualFold(item.Severity, "high") {
			return true
		}
	}
	return false
}
