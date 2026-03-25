# Skill: Wrapper Depth Minimization

## Category
Layout

## Description
This skill teaches the AI agent to avoid unnecessary nesting of wrapper divs that add structural complexity without contributing to the visual or semantic output. Every additional wrapper element increases the depth of the DOM tree, makes CSS selectors harder to reason about, and introduces potential layout bugs through unintended flex or grid context inheritance. The agent should learn to recognize when a wrapper div serves no purpose and can be removed, or when its styling can be merged into an adjacent element.

A common source of wrapper bloat is component composition, where each layer of a component hierarchy adds its own container div. The agent should prefer using React fragments (`<>...</>`) or the `as` prop pattern to avoid extraneous DOM nodes. When a wrapper is needed for styling purposes, the agent should ensure it does not introduce an additional flex or grid context that disrupts the parent layout. The rule of thumb is: if removing the wrapper does not change the visual output, the wrapper should be removed.

The agent should also watch for "div soup" patterns where multiple nested divs each carry a single class for margin, padding, or alignment. These should be consolidated into a single element with combined styles. Deeply nested structures make debugging layout issues harder and increase the likelihood of specificity conflicts. A flat, semantic DOM is easier to maintain, faster to render, and more accessible to assistive technologies.

## Expected Outcome
- No wrapper div exists solely to apply a single margin or alignment class
- Component output uses fragments instead of unnecessary container divs
- DOM depth for typical page sections does not exceed 6-8 levels of nesting
- Wrapper elements that remain serve a clear structural or semantic purpose

## Heuristic Triggers
- `CountAddedWrapperDivs`: Detects new wrapper elements that do not contribute semantic meaning or essential styling
- `ExtractClassNames`: Identifies class names like `wrapper`, `container`, `inner` applied to single-child elements
- `CountNewTopLevelComponents`: Flags components that introduce excessive nesting relative to their output complexity

## Scoring Axes Most Affected
- visual_noise: Fewer wrapper elements reduce DOM clutter and make the rendered structure cleaner
- density_control: Flat structures pack content more efficiently without wasted space from nested padding
- component_reuse: Components without extraneous wrappers are easier to compose and reuse in different layouts
