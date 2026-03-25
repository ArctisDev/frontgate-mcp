# Skill: Responsive Breakpoints

## Category
Layout

## Description
This skill teaches the AI agent to design layouts that adapt gracefully across screen sizes using a mobile-first approach. Mobile-first means writing base styles for the smallest screen and layering on complexity with `min-width` media queries as the viewport grows. This approach produces leaner CSS because the mobile layout is the default, and desktop enhancements are additive. The agent should avoid the common anti-pattern of designing desktop-first and then overriding with `max-width` queries, which leads to bloated and contradictory stylesheets.

A disciplined breakpoint strategy uses a small, consistent set of breakpoints rather than device-specific values. The agent should work with breakpoints like 640px, 768px, 1024px, and 1280px, which correspond to common content-driven layout shifts rather than specific device models. Each breakpoint should trigger a meaningful layout change, such as moving from a single column to a multi-column grid, revealing a sidebar, or adjusting font sizes. The agent should not introduce breakpoints for minor pixel-level adjustments that could be handled by fluid units or `clamp()`.

When the agent encounters an existing codebase, it should identify and reuse the existing breakpoint definitions rather than adding new ones. Introducing a 960px breakpoint when the codebase already uses 1024px creates inconsistency and makes the layout harder to reason about. The agent should also ensure that content is accessible and readable at every breakpoint, avoiding layouts that break or hide critical content at intermediate screen sizes.

## Expected Outcome
- Base styles target mobile viewports, with `min-width` media queries enhancing for larger screens
- A small consistent set of breakpoints (3-5) is reused across the codebase
- Each breakpoint triggers a meaningful layout shift, not minor pixel tweaks
- Fluid units and `clamp()` handle continuous scaling between breakpoints

## Heuristic Triggers
- `FindArbitrarySpacing`: Detects hardcoded pixel values in media queries that should use breakpoint tokens
- `ExtractClassNames`: Identifies device-specific class names (e.g., `tablet-`, `iphone-`) suggesting non-principled breakpoints
- `CountRawPrimitiveAdditions`: Flags new raw HTML elements added to handle responsive concerns that existing components should manage

## Scoring Axes Most Affected
- layout_balance: Proper breakpoints ensure proportional layouts at every screen size
- template_likeness: Standard breakpoint behavior matches user expectations across devices
- ux_clarity: Responsive layouts maintain readability and interaction targets at all viewport sizes
