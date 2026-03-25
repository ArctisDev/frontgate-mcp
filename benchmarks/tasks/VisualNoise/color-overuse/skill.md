# Skill: Using Color Intentionally

## Category
VisualNoise

## Description
Color is one of the strongest signals in a UI, but when an agent splashes multiple hues across every section — blue headers, green badges, orange alerts, purple links, red buttons — the palette becomes noise. Each color competes for attention, and the user cannot tell which colors carry meaning versus which are decorative. Agents tend to add color variety to make layouts feel lively, but this undermines communication.

This skill teaches limiting the active palette to a primary brand color, a neutral background family, and one or two semantic colors (success, danger, warning). Every color choice should serve a purpose: indicating state, guiding action, or reinforcing brand. Decorative color that communicates nothing should be removed. The goal is a calm interface where color draws the eye only where it matters.

## Expected Outcome
- Active color count reduced to brand primary plus two semantic colors
- No decorative color applied without communicative intent
- Color used consistently for the same meaning across the UI
- Lower visual noise from competing hues

## Heuristic Triggers
- `FindHardcodedPaletteClasses`: Detects hardcoded color classes outside the token palette
- `CountGenericSaaSLanguage`: Can indicate generic styling language that accompanies decorative color choices

## Scoring Axes Most Affected
- token_adherence: Color tokens enforce a controlled palette over arbitrary values
- visual_noise: Fewer distinct colors create a calmer, more cohesive appearance
