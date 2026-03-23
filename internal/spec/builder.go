package spec

import (
	"fmt"
	"strings"

	"github.com/ArctisDev/frontgate/internal/types"
)

func Build(in types.BuildUISpecInput) types.UISpec {
	context := in.Context
	scope := []string{
		"Improve the target screen without redesigning the whole product shell.",
		"Keep the existing design system and reuse current primitives before creating anything new.",
		"Focus on hierarchy, spacing rhythm, density control and product fit.",
	}

	constraints := []string{
		"Stay inside the existing Next.js App Router + Tailwind + shadcn-style stack.",
		"Prefer existing tokens, spacing scale and primitives over arbitrary values.",
		"Do not add decorative widgets, vanity metrics or empty card shells.",
		"Do not widen sidebars or create boxed layouts that reduce content efficiency.",
		"Improve clarity through typography, grouping and layout balance instead of visual clutter.",
	}

	antiPatterns := []string{
		"Do not invent new dashboard widgets unless the task explicitly requires new information architecture.",
		"Do not add arbitrary padding, gap or font-size values when existing scales work.",
		"Do not stack multiple nested cards for a single interaction.",
		"Do not create premium-looking chrome that weakens task completion or density.",
		"Do not replace existing reusable primitives with ad hoc markup.",
	}

	files := chooseFiles(context, in.ReferenceFile)
	reuse := chooseReuse(context)
	acceptance := []string{
		"The target page feels less template-like and more aligned with the product context.",
		"Spacing uses the established scale consistently, with fewer arbitrary values and wrapper layers.",
		"Visual hierarchy is clearer: title, supporting text, actions and content blocks have distinct emphasis.",
		"Existing primitives are reused for buttons, cards, inputs, dialogs, tables and tabs wherever applicable.",
		"No decorative UI was added without supporting a real product action or state.",
	}

	notes := []string{
		fmt.Sprintf("Detected stack: %s", context.VisualSystemSummary),
		fmt.Sprintf("Most relevant files: %s", strings.Join(files, ", ")),
		fmt.Sprintf("Reusable primitives available: %s", strings.Join(reuse, ", ")),
	}

	objective := deriveObjective(in.Task)
	prompt := buildPrompt(objective, scope, files, reuse, constraints, antiPatterns, acceptance, notes)

	return types.UISpec{
		Objective:           objective,
		Scope:               scope,
		ProbableFiles:       files,
		ComponentsToReuse:   reuse,
		DesignConstraints:   constraints,
		AntiPatternsToAvoid: antiPatterns,
		AcceptanceCriteria:  acceptance,
		ExecutionNotes:      notes,
		StructuredPrompt:    prompt,
	}
}

func deriveObjective(task string) string {
	task = strings.TrimSpace(task)
	if task == "" {
		return "Tighten the target UI using the product's existing visual system, with stronger hierarchy and less generic SaaS styling."
	}
	return fmt.Sprintf("Execute this task with stricter UI discipline: %s", task)
}

func chooseFiles(ctx types.RepoContext, referenceFile string) []string {
	var files []string
	if referenceFile != "" {
		files = append(files, referenceFile)
	}
	files = append(files, ctx.TargetFileHints...)
	files = append(files, ctx.RelevantFiles...)
	if len(files) > 10 {
		files = files[:10]
	}
	return dedupe(files)
}

func chooseReuse(ctx types.RepoContext) []string {
	reuse := append([]string{}, ctx.BaseComponents...)
	reuse = append(reuse, ctx.PatternInventory.Cards...)
	reuse = append(reuse, ctx.PatternInventory.Forms...)
	reuse = append(reuse, ctx.PatternInventory.Tables...)
	if len(reuse) > 12 {
		reuse = reuse[:12]
	}
	return dedupe(reuse)
}

func buildPrompt(objective string, scope, files, reuse, constraints, antiPatterns, acceptance, notes []string) string {
	var b strings.Builder
	b.WriteString("Objective:\n")
	b.WriteString("- " + objective + "\n\n")

	b.WriteString("Scope:\n")
	for _, item := range scope {
		b.WriteString("- " + item + "\n")
	}
	b.WriteString("\nProbable Files:\n")
	for _, item := range files {
		b.WriteString("- " + item + "\n")
	}
	b.WriteString("\nReuse These Components First:\n")
	for _, item := range reuse {
		b.WriteString("- " + item + "\n")
	}
	b.WriteString("\nDesign Constraints:\n")
	for _, item := range constraints {
		b.WriteString("- " + item + "\n")
	}
	b.WriteString("\nAvoid:\n")
	for _, item := range antiPatterns {
		b.WriteString("- " + item + "\n")
	}
	b.WriteString("\nAcceptance Criteria:\n")
	for _, item := range acceptance {
		b.WriteString("- " + item + "\n")
	}
	b.WriteString("\nRepo Notes:\n")
	for _, item := range notes {
		b.WriteString("- " + item + "\n")
	}
	b.WriteString("\nImplementation instruction:\n")
	b.WriteString("- Make the smallest coherent set of changes that improves hierarchy, spacing and product fit.\n")
	b.WriteString("- Reuse existing primitives and tokens, then run critique_ui_output and score_ui_quality before considering the task done.\n")
	return b.String()
}

func dedupe(items []string) []string {
	set := map[string]struct{}{}
	var out []string
	for _, item := range items {
		if item == "" {
			continue
		}
		if _, ok := set[item]; ok {
			continue
		}
		set[item] = struct{}{}
		out = append(out, item)
	}
	return out
}
