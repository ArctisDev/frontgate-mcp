# Skill: Grid System Fundamentals

## Category
Layout

## Description
This skill teaches the AI agent how to construct responsive layouts using CSS Grid and Flexbox in a principled way. Grid and Flexbox are the foundational layout primitives in modern CSS, and misusing them leads to fragile, hard-to-maintain structures. The agent should learn to use Grid for two-dimensional layouts (rows and columns together) and Flexbox for one-dimensional flows (a row or a column), choosing the right tool for each context rather than nesting Flexbox containers where a Grid would be cleaner.

A consistent grid system requires defining a predictable set of column counts and gap values. The agent should avoid inventing ad-hoc grid configurations per section and instead reuse a small set of grid patterns (e.g., 12-column grid, 2-column sidebar layout, 3-column card grid). Gaps between grid items must use spacing tokens or CSS custom properties rather than hardcoded pixel values to maintain visual rhythm across the interface. When the agent introduces a new grid layout, it should verify that column counts and gap sizes align with existing patterns in the codebase.

Overriding grid behavior with margin hacks or negative margins is a sign of layout fragility. The agent should prefer `gap`, `align-items`, and `justify-content` for spacing and alignment within grid and flex containers. If a layout requires complex margin overrides to achieve the desired result, the grid definition itself is likely wrong and should be reconsidered.

## Expected Outcome
- Layouts use CSS Grid for two-dimensional structures and Flexbox for single-axis flows
- Grid column counts come from a small reusable set rather than ad-hoc values
- Spacing between grid items uses `gap` with design tokens, not hardcoded pixel margins
- Grid and flex alignment properties are used instead of margin hacks for positioning

## Heuristic Triggers
- `FindArbitrarySpacing`: Detects hardcoded pixel values in gap or margin that bypass spacing tokens
- `CountNewTopLevelComponents`: Flags when new top-level components introduce novel grid patterns without reusing existing ones
- `ExtractClassNames`: Identifies class names suggesting one-off grid configurations that break consistency

## Scoring Axes Most Affected
- layout_balance: Proper grid systems create consistent visual rhythm and proportional column distributions
- spacing_consistency: Token-based gaps ensure uniform spacing across all grid-based layouts
- component_reuse: Reusing a small set of grid patterns reduces code duplication and improves maintainability
