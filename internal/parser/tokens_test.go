package parser

import "testing"

func TestExtractCSSVariables(t *testing.T) {
	body := `
	:root {
	  --background: 0 0% 100%;
	  --radius: 0.75rem;
	}
	`
	got := ExtractCSSVariables(body)
	if got["background"] != "0 0% 100%" {
		t.Fatalf("expected background token, got %#v", got)
	}
	if got["radius"] != "0.75rem" {
		t.Fatalf("expected radius token, got %#v", got)
	}
}

func TestExtractSpacingClasses(t *testing.T) {
	body := `<div className="px-4 py-6 gap-3 p-[18px]">`
	got := ExtractSpacingClasses(body)
	expected := map[string]bool{
		"px-4":     true,
		"py-6":     true,
		"gap-3":    true,
		"p-[18px]": true,
	}
	for _, item := range got {
		delete(expected, item)
	}
	if len(expected) != 0 {
		t.Fatalf("missing spacing classes: %#v", expected)
	}
}
