# Skill: Purposeful Gradient Usage

## Category
VisualNoise

## Description
Gradients add depth and visual richness when used well — a subtle gradient on a hero banner or a soft fade on a card overlay can elevate the design. But agents tend to apply gradients to backgrounds, buttons, text, and section dividers indiscriminately because gradients look more polished than flat colors. The result is a UI where every surface has a color shift, creating a shimmering, unfocused visual experience.

This skill teaches using gradients sparingly and with clear intent. A gradient should serve one of three purposes: creating depth on an elevated surface, guiding the eye along a content hierarchy, or adding brand personality to a hero or feature section. Do not apply gradients to buttons, small cards, sidebar panels, or form elements. For these, flat colors from the token palette are the correct choice.

## Expected Outcome
- Gradients limited to hero sections, feature highlights, or overlay effects
- Buttons, inputs, and small surfaces use flat token colors
- No more than one or two gradient instances per viewport
- Gradient angles and stops follow a consistent system

## Heuristic Triggers
- `FindHardcodedPaletteClasses`: Detects arbitrary gradient classes outside the design system
- `CountRawPrimitiveAdditions`: Flags raw gradient CSS added without token backing

## Scoring Axes Most Affected
- visual_noise: Fewer gradients mean fewer competing color transitions
- token_adherence: Gradient tokens ensure consistency and prevent arbitrary color stops
