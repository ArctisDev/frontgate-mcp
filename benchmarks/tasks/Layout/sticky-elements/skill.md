# Skill: Sticky Elements

## Category
Layout

## Description
This skill teaches the AI agent to use `position: sticky` purposefully and sparingly for elements that benefit from remaining visible during scroll. Sticky headers (navigation bars that stick to the top as the user scrolls down) and sticky table headers (column labels that remain visible while scrolling through data rows) are legitimate uses that improve usability. Sticky sidebars (table of contents that follow the user through a long article) can also be valuable on content-heavy pages. However, the agent should not make elements sticky simply because they might be useful to always see. Each sticky element consumes screen real estate while the user scrolls, reducing the visible content area.

The agent must ensure that sticky elements have a defined ancestor context. `position: sticky` only works within a scrollable ancestor, and if the element is a direct child of a full-height body with no overflow container, it will behave as `position: relative`. The agent should verify that the sticky element's parent has the appropriate height and overflow properties to enable the sticky behavior. Additionally, the agent should set a `top` value (e.g., `top: 0` for sticky headers) that accounts for any other sticky elements above it to prevent overlap.

The agent should avoid stacking multiple sticky elements, which compounds the loss of vertical content space. If a page has a sticky header, a sticky sub-navigation, and a sticky announcement banner, the effective viewport for content shrinks dramatically. In such cases, the agent should consolidate or make some elements scroll away. Sticky elements should also have a `z-index` that places them above regular content but below modals and dropdowns, maintaining a clear stacking order.

## Expected Outcome
- `position: sticky` is used only for elements with a clear usability benefit during scroll
- Sticky elements have correct `top` values and are within a proper scrollable ancestor context
- Multiple sticky elements are not stacked, preserving content viewport space
- `z-index` values for sticky elements follow a consistent stacking hierarchy

## Heuristic Triggers
- `FindArbitrarySpacing`: Detects hardcoded `top` or `z-index` values in sticky element definitions that bypass design tokens
- `FindOversizedSidebarClasses`: Flags sticky sidebars that consume excessive width and viewport space simultaneously
- `CountNewTopLevelComponents`: Identifies new components that add sticky behavior without clear justification

## Scoring Axes Most Affected
- density_control: Proper sticky usage preserves content viewport space rather than shrinking it with persistent chrome
- ux_clarity: Purposeful sticky elements (like table headers) improve readability and navigation
- layout_balance: Correctly stacked sticky elements maintain z-order and prevent visual overlap issues
