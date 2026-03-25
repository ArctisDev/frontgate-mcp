# Skill: Intentional Whitespace Usage

## Category
Design

## Description
Whitespace (negative space) is not wasted space—it is one of the most powerful tools in a designer's toolkit for creating clarity, focus, and hierarchy. Agents under time pressure tend to fill every pixel with content, reducing padding, collapsing margins, and cramming elements together to "fit more on screen." This creates a dense, overwhelming interface where nothing stands out because everything competes for attention. Generous whitespace around key content areas, between form fields, and inside cards gives the eye room to rest and makes important elements pop.

This skill teaches the agent to treat whitespace as an active design decision. When building new sections, the agent should resist the urge to minimize padding. Instead, it should apply the project's standard spacing tokens liberally—py-12 or py-16 for section padding, p-6 for card interiors, gap-4 for item groups. The agent should recognize that reducing whitespace to "save space" usually hurts UX more than it helps. Form layouts should have generous gap-y values, hero sections should breathe with large vertical padding, and sidebar content should not be flush against the edges.

## Expected Outcome
- The agent applies generous, consistent padding to sections, cards, and containers rather than minimizing space
- Form fields have adequate vertical and horizontal gaps for comfortable interaction
- Hero and header sections use large vertical padding (py-16, py-24) to create a sense of openness
- The agent does not reduce existing whitespace when modifying components

## Heuristic Triggers
- `FindArbitrarySpacing`: Detects small or inconsistent spacing values that may indicate the agent is squeezing elements together
- `CountAddedWrapperDivs`: Excessive wrappers can sometimes indicate the agent is fighting whitespace issues with structural hacks instead of proper spacing
- `ExtractClassNames`: Surfaces spacing utilities to audit whether padding and margin values are appropriately generous

## Scoring Axes Most Affected
- density_control: Whitespace directly governs how dense or airy a layout feels
- ux_clarity: Generous whitespace improves readability and reduces cognitive load
- visual_noise: Proper whitespace reduces clutter by giving elements room to breathe
