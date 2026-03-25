package spec

import (
	"fmt"
	"strings"

	"github.com/ArctisDev/frontgate/internal/references"
	"github.com/ArctisDev/frontgate/internal/types"
)

func Build(in types.BuildUISpecInput) types.UISpec {
	context := in.Context
	objective := deriveObjective(in.Task)

	files := chooseFiles(context, in.ReferenceFile)
	reuse := chooseReuse(context)

	artDirection := buildArtDirection(in.Task, context)
	colorStrategy := buildColorStrategy(context)
	typographyStrategy := buildTypographyStrategy()
	layoutApproach := buildLayoutApproach(in.Task)
	visualElements := buildVisualElements()
	animationStrategy := buildAnimationStrategy()
	feelStatement := buildFeelStatement()
	antiPatterns := buildAntiPatterns()
	acceptance := buildAcceptanceCriteria()

	scope := []string{
		"Transform this screen into a visually distinctive, art-directed experience.",
		"Every design decision must serve identity, not just function.",
		"The result should feel like it was designed by an art director, not assembled by a template engine.",
	}

	notes := []string{
		fmt.Sprintf("Detected stack: %s", context.VisualSystemSummary),
		fmt.Sprintf("Most relevant files: %s", strings.Join(files, ", ")),
		fmt.Sprintf("Reusable primitives available: %s", strings.Join(reuse, ", ")),
	}

	artBrief := references.GetArtBrief("")
	if artBrief != "" {
		notes = append(notes, "VISUAL REFERENCES: The user has provided design reference images. See the art brief below.")
		artDirection = append(artDirection, "\n"+artBrief)
	}

	prompt := buildPrompt(
		objective, scope, files, reuse,
		artDirection, colorStrategy, typographyStrategy,
		layoutApproach, visualElements, animationStrategy,
		feelStatement, antiPatterns, acceptance, notes,
	)

	return types.UISpec{
		Objective:           objective,
		Scope:               scope,
		ProbableFiles:       files,
		ComponentsToReuse:   reuse,
		DesignConstraints:   append(artDirection, append(colorStrategy, typographyStrategy...)...),
		AntiPatternsToAvoid: antiPatterns,
		AcceptanceCriteria:  acceptance,
		ExecutionNotes:      notes,
		StructuredPrompt:    prompt,
	}
}

func deriveObjective(task string) string {
	task = strings.TrimSpace(task)
	if task == "" {
		return "Create a visually distinctive, art-directed interface with strong identity, premium feel, and memorable visual presence."
	}
	return fmt.Sprintf("Execute with art direction and visual identity as top priorities: %s", task)
}

func buildArtDirection(task string, ctx types.RepoContext) []string {
	direction := []string{
		"ART DIRECTION: This is not a template — it is a product with identity.",
		"Think editorial design, campaign landing pages, fashion sites, premium portfolios — not SaaS boilerplate.",
		"Every section must have its own visual personality. No two sections should feel like copies of each other.",
		"Composition must mix order with visual tension: structured grids broken with intention, asymmetric layouts, and bold visual anchors.",
		"Use large breathing space contrasting with dense content blocks to create rhythm.",
		"Images, 3D renders, shapes, or graphic objects should serve as the primary visual anchor of compositions — not decoration.",
		"Accept influences from: brutalism (soft), poster design, fashion landing pages, concept interfaces, campaign sites.",
	}

	if strings.Contains(strings.ToLower(task), "landing") || strings.Contains(strings.ToLower(task), "hero") {
		direction = append(direction,
			"HERO SECTION: Must dominate the viewport. Large headline with real weight. Subheadline with purpose. CTA that feels intentional, not generic.",
			"Avoid: terminal mockups, fake dashboard screenshots, stock illustration hero images.",
			"Prefer: bold typography as the hero itself, or a single strong visual element (3D object, abstract shape, product render) with minimal chrome.",
		)
	}

	if strings.Contains(strings.ToLower(task), "dashboard") || strings.Contains(strings.ToLower(task), "app") {
		direction = append(direction,
			"APP/UTILITY SCREENS: Even utility screens deserve design identity. Use the product's color accent strategically on interactive elements.",
			"Don't let data density kill visual hierarchy. Use whitespace as a design tool, not leftover space.",
		)
	}

	return direction
}

func buildColorStrategy(ctx types.RepoContext) []string {
	strategy := []string{
		"COLOR STRATEGY:",
		"- Base: strong neutrals — black, off-white, light gray, graphite, blue-gray, or deep tones.",
		"- Choose ONE dominant accent color per screen/section. Not multiple competing colors.",
		"- Approved accents: acid green, intense orange, deep red, electric blue, lilac glow, soft purple, ice blue.",
		"- The accent must create contrast and visual signature — it should never feel like decoration.",
		"- Use color sparingly and with purpose. Let the neutral base do 80% of the work.",
	}

	if len(ctx.DesignTokens) > 0 {
		strategy = append(strategy,
			fmt.Sprintf("- Existing project tokens detected (%d variables). Use them as the foundation, but extend with the accent strategy above.", len(ctx.DesignTokens)),
		)
	} else {
		strategy = append(strategy,
			"- No design tokens detected in the project. Define CSS custom properties for your color system before applying colors.",
		)
	}

	return strategy
}

func buildTypographyStrategy() []string {
	return []string{
		"TYPOGRAPHY STRATEGY:",
		"- Typography is the protagonist of this design. Headlines must dominate with presence and excellent contrast.",
		"- Use modern sans-serif: grotesk, neo-grotesk, geometric, or display fonts with personality.",
		"- Recommended font families: Inter (versatile), Space Grotesk (geometric character), Instrument Sans (editorial), Sora (geometric warmth), Clash Display (bold display).",
		"- For display/hero: use a font weight of 600-800 at large sizes (clamp(2.5rem, 5vw, 5rem)).",
		"- Secondary text should be more contained and functional — smaller, lighter, less prominent.",
		"- Avoid: default system fonts, overly neutral/corporate fonts, fonts that look like every other SaaS.",
		"- Font loading: use next/font with Google Fonts subset to avoid FOUT/FOIT. Never rely on external CDN links for critical fonts.",
	}
}

func buildLayoutApproach(task string) []string {
	approach := []string{
		"LAYOUT APPROACH:",
		"- Avoid the predictable pattern: hero + 3 feature cards + logos + CTA. This is the definition of generic.",
		"- Create compositions with rhythm: alternate between full-width sections and contained editorial blocks.",
		"- Use asymmetric grids: 40/60, 30/70 splits instead of always 50/50 or centered.",
		"- Break the grid intentionally: elements can overflow their containers, overlap sections, or extend beyond expected boundaries.",
		"- Use sections framed within a central canvas when it reinforces the aesthetic (editorial feel).",
		"- Large rounded corners (rounded-2xl, rounded-3xl) when they serve the composition, not everywhere.",
		"- Think of the page as a digital magazine or campaign poster, not a utility form.",
	}

	if strings.Contains(strings.ToLower(task), "landing") {
		approach = append(approach,
			"- Suggested section flow (NOT a template):",
			"  1. Hero: massive headline + single visual anchor + subtle motion",
			"  2. Statement: a bold claim or proof point, not a feature list",
			"  3. Feature showcase: ONE feature told as a story (image + text), not a grid of cards",
			"  4. Social proof: woven into the design, not a separate logos bar",
			"  5. How it works: visual process, not 3 steps with icons",
			"  6. CTA: emotionally driven, not just 'Get Started'",
			"- Each section should feel like a different page of a magazine, not a repeating template.",
		)
	}

	return approach
}

func buildVisualElements() []string {
	return []string{
		"VISUAL ELEMENTS:",
		"- Use impactful imagery: 3D renders, abstract objects, strong product screenshots, graphic shapes.",
		"- Glow, blur, gradients, grain textures, glass effects — only when they add sophistication, not just because they look cool.",
		"- Shapes and graphic elements must reinforce identity and direction. They should never feel like random decoration.",
		"- Mesh gradients on backgrounds (subtle, not rainbow). Noise/grain texture overlays for depth.",
		"- Glassmorphism on elevated panels: backdrop-blur, semi-transparent backgrounds with subtle borders.",
		"- Geometric shapes as section dividers or visual anchors (not just horizontal rules).",
		"- Consider: a strong visual element per section that makes it memorable (a large number, a shape, a color block, an image crop).",
	}
}

func buildAnimationStrategy() []string {
	return []string{
		"ANIMATION & MOTION:",
		"- Motion must serve a purpose: guide attention, provide feedback, create depth, or establish rhythm.",
		"- Scroll-triggered reveals: elements entering the viewport with staggered timing (opacity + translate).",
		"- Hover states: meaningful transitions on interactive elements (scale, color shift, shadow depth).",
		"- Micro-interactions: button press feedback, loading states, toggle animations.",
		"- Parallax (subtle): background elements moving at different speed than foreground for depth.",
		"- Use CSS animations and Framer Motion. Avoid: auto-playing carousels, infinite loops, animations that block interaction.",
		"- Timing: use ease-out for entrances, ease-in for exits. Duration 200-400ms for UI, 600-1000ms for decorative.",
		"- Respect prefers-reduced-motion: provide a fallback that disables non-essential animations.",
	}
}

func buildFeelStatement() []string {
	return []string{
		"THE FEELING:",
		"- Premium. High-contrast. Art-directed. Visually strong.",
		"- Clean, but not boring. Experimental, but with control. Modern, but not a caricature of futurism.",
		"- The interface must have its own identity, immediate impact, and structural clarity.",
		"- It should look like it was designed by a creative director for a product launch, not assembled from a component library.",
	}
}

func buildAntiPatterns() []string {
	return []string{
		"ABSOLUTELY AVOID:",
		"- Generic SaaS patterns: hero with gradient text + 3 feature cards + logos bar + pricing table + CTA.",
		"- Dashboard mockups or terminal windows in hero sections.",
		"- Multiple competing colors. Rainbow gradients. Neon on dark without purpose.",
		"- Tiny text, cramped spacing, or walls of unbroken content.",
		"- Decorative elements that don't serve the composition (random dots, lines, circles).",
		"- Stock illustrations or undraw-style graphics.",
		"- Rounded corners on everything. Use them with intention.",
		"- Shadow on every element. Reserve elevation cues for true depth needs.",
		"- Icons as decoration next to every heading. Let typography speak.",
		"- Templates that could belong to any SaaS. The result must feel specific to THIS product.",
		"- External CDN font links that might 404. Always use next/font or self-hosted fonts.",
	}
}

func buildAcceptanceCriteria() []string {
	return []string{
		"ACCEPTANCE CRITERIA:",
		"- The page has a clear visual identity that distinguishes it from generic SaaS templates.",
		"- Typography is expressive, well-weighted, and creates a strong scan path.",
		"- Color is used strategically: one accent creates the signature, neutrals do the heavy lifting.",
		"- Layout has rhythm and variation: no two consecutive sections feel identical.",
		"- Visual elements (images, shapes, textures) serve the composition, not just fill space.",
		"- Animations are purposeful and enhance the experience without blocking interaction.",
		"- Font loading is handled correctly (no FOUT, no 404s, no external CDN dependency for critical fonts).",
		"- The result feels premium, intentional, and memorable — something a user would screenshot and share.",
	}
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

func buildPrompt(
	objective string, scope, files, reuse []string,
	artDirection, colorStrategy, typographyStrategy []string,
	layoutApproach, visualElements, animationStrategy []string,
	feelStatement, antiPatterns, acceptance, notes []string,
) string {
	var b strings.Builder

	b.WriteString("=== DESIGN BRIEF ===\n\n")
	b.WriteString("Objective:\n- " + objective + "\n\n")

	b.WriteString("Scope:\n")
	for _, item := range scope {
		b.WriteString("- " + item + "\n")
	}

	b.WriteString("\nTarget Files:\n")
	for _, item := range files {
		b.WriteString("- " + item + "\n")
	}

	b.WriteString("\nReuse These Components:\n")
	for _, item := range reuse {
		b.WriteString("- " + item + "\n")
	}

	b.WriteString("\n")
	for _, item := range artDirection {
		b.WriteString(item + "\n")
	}

	b.WriteString("\n")
	for _, item := range colorStrategy {
		b.WriteString(item + "\n")
	}

	b.WriteString("\n")
	for _, item := range typographyStrategy {
		b.WriteString(item + "\n")
	}

	b.WriteString("\n")
	for _, item := range layoutApproach {
		b.WriteString(item + "\n")
	}

	b.WriteString("\n")
	for _, item := range visualElements {
		b.WriteString(item + "\n")
	}

	b.WriteString("\n")
	for _, item := range animationStrategy {
		b.WriteString(item + "\n")
	}

	b.WriteString("\n")
	for _, item := range feelStatement {
		b.WriteString(item + "\n")
	}

	b.WriteString("\n")
	for _, item := range antiPatterns {
		b.WriteString(item + "\n")
	}

	b.WriteString("\n")
	for _, item := range acceptance {
		b.WriteString(item + "\n")
	}

	b.WriteString("\nRepo Notes:\n")
	for _, item := range notes {
		b.WriteString("- " + item + "\n")
	}

	b.WriteString("\nImplementation instruction:\n")
	b.WriteString("- This is a DESIGN task, not a refactoring task. You are creating visual identity.\n")
	b.WriteString("- Every CSS decision must be intentional. No defaults without thought.\n")
	b.WriteString("- Use next/font for all typography. Never use external CDN links for fonts.\n")
	b.WriteString("- After implementation, run critique_ui_output and score_ui_quality.\n")
	b.WriteString("- If the score is below 75, iterate. The design must pass the quality gate.\n")

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
