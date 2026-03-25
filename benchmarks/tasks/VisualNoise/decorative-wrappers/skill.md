# Skill: Removing Decorative Wrapper Divs

## Category
VisualNoise

## Description
Wrapper divs that exist purely to add a margin, padding, or a background color fragment the DOM and introduce unnecessary visual layers. When an agent generates markup, it often adds intermediate containers that serve no structural purpose — they exist only because the agent defaulted to wrapping content in extra elements to achieve spacing or color. This inflates the HTML tree and makes the layout harder to reason about.

This skill teaches writing clean, semantic HTML without presentation-only wrappers. Use utility classes or inline styles directly on the meaningful container rather than adding an extra div. Every element in the DOM should serve a structural, semantic, or genuinely interactive purpose. If removing a div changes nothing visually or functionally, it should not exist.

## Expected Outcome
- Flatter DOM structure with fewer nesting levels
- No divs added solely for background color or margin
- Cleaner class lists on meaningful containers
- Lower CountAddedWrapperDivs score in output

## Heuristic Triggers
- `CountAddedWrapperDivs`: Counts wrapper divs that add no structural value
- `ExtractClassNames`: Surfaces classes that indicate pure-wrapper patterns

## Scoring Axes Most Affected
- visual_noise: Fewer layers means fewer visual boundaries competing for attention
- template_likeness: Flatter markup resembles production templates more closely
