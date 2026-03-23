package rules

import "testing"

func TestFindArbitrarySpacing(t *testing.T) {
	diff := `+ <div className="px-4 gap-[22px] py-[13px]">`
	got := FindArbitrarySpacing(diff)
	if len(got) != 2 {
		t.Fatalf("expected 2 arbitrary spacing values, got %d (%v)", len(got), got)
	}
}

func TestCountGenericSaaSLanguage(t *testing.T) {
	text := "Overview widget with analytics summary and quick action block"
	if CountGenericSaaSLanguage(text) < 3 {
		t.Fatalf("expected generic SaaS language count >= 3")
	}
}

func TestFindHardcodedPaletteClasses(t *testing.T) {
	diff := `+ <div className="bg-blue-500 text-slate-500 border-red-400">`
	got := FindHardcodedPaletteClasses(diff)
	if len(got) != 3 {
		t.Fatalf("expected 3 hardcoded palette classes, got %v", got)
	}
}

func TestCountExistingComponentReuse(t *testing.T) {
	diff := `+ <Button /><Input /><div />`
	got := CountExistingComponentReuse(diff, []string{"button", "input", "card"})
	if got != 2 {
		t.Fatalf("expected 2 reused components, got %d", got)
	}
}

func TestHasClearHeadingStructure(t *testing.T) {
	if !HasClearHeadingStructure(`+ <h1 className="text-3xl font-semibold">Deploy</h1>`) {
		t.Fatalf("expected heading structure detection")
	}
}
