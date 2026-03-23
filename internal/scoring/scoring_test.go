package scoring

import (
	"testing"

	"github.com/ArctisDev/frontgate/internal/types"
)

func TestScoreWeightedAverage(t *testing.T) {
	in := types.ScoreUIQualityInput{
		CritiqueReport: types.CritiqueReport{
			Scores: types.CritiqueScores{
				SpacingConsistency:   80,
				TypographicHierarchy: 70,
				LayoutBalance:        90,
				ComponentReuse:       60,
				TokenAdherence:       85,
				DensityControl:       75,
				VisualNoise:          88,
				TemplateLikeness:     78,
				UXClarity:            82,
				ProductAlignment:     74,
			},
		},
	}
	got := Score(in)
	if got.TotalScore < 70 || got.TotalScore > 90 {
		t.Fatalf("unexpected score: %#v", got)
	}
}

func TestGateRejectsBlockingIssues(t *testing.T) {
	result := Gate(types.GateUIChangeInput{
		Score: types.QualityScore{TotalScore: 88, RecommendedThreshold: 80},
		Critique: types.CritiqueReport{
			Problems: []types.CritiqueIssue{
				{Category: "layout", Severity: "high", Detail: "overflow"},
			},
			CorrectiveActions: []string{"fix overflow"},
		},
	})
	if result.Approved {
		t.Fatalf("expected gate rejection when blocking issue exists")
	}
}
