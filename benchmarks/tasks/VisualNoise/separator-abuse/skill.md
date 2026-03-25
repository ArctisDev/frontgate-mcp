# Skill: Replacing Separators with Spacing

## Category
VisualNoise

## Description
Horizontal rules and border-based dividers are the most common way agents separate content sections. But `<hr>` elements and `border-b` utilities add visual weight proportional to a full-width line — they interrupt the reading flow and create hard visual stops. When every section is demarcated by a horizontal line, the page looks like a form from the 1990s rather than a modern interface.

This skill teaches that whitespace is the superior separator. Generous margin or padding between sections communicates grouping and hierarchy without drawing a line through the layout. Background color shifts between adjacent sections also create clear boundaries without literal lines. Reserve horizontal rules for genuinely distinct content domains — for example, separating the main content from a footer — not for every minor content transition.

## Expected Outcome
- Horizontal rules and border-based dividers removed from within content areas
- Spacing tokens used consistently to separate sections
- Background alternation used for major section transitions
- Clean, line-free content areas with clear visual grouping

## Heuristic Triggers
- `FindArbitrarySpacing`: Detects ad-hoc spacing values that should use consistent tokens
- `CountRawPrimitiveAdditions`: Flags raw separator elements added outside the component system

## Scoring Axes Most Affected
- visual_noise: Removing lines eliminates horizontal visual interruptions
- spacing_consistency: Proper spacing tokens replace ad-hoc divider placement
