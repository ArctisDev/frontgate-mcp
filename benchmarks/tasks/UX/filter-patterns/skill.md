# Skill: Filter Implementation Patterns

## Category
UX

## Description
Filters help users narrow down large content sets to find what they need, but poorly implemented filters create more frustration than value. The best filter patterns use filter chips — compact, removable tags that show the user exactly which filters are active at any time. Each chip should display the filter name and selected value (e.g., "Status: Active") and offer a clear way to remove it. A "Clear all" action should be available when multiple filters are active. The result count should update in real-time as filters change, giving users immediate feedback on how their selections affect the content. Without result counts, users are guessing whether their filters are working.

The AI agent should learn that filters should never require a separate "Apply" button when the results can update in real-time. Delayed application forces users to guess what will happen before they see results. Filters should persist across navigation within a session — removing all filters when a user navigates away and returns is frustrating. Avoid hiding active filters behind a collapsed panel; the current filter state should always be visible. Default filters should be meaningful (e.g., showing only active items by default) rather than showing everything. Avoid offering filters that return zero results — if a filter combination yields nothing, indicate this before the user applies it.

## Expected Outcome
- Active filters are displayed as removable chips showing name and selected value
- Result count updates in real-time as filters are applied or removed
- A "Clear all" action is available when multiple filters are active
- Filters persist across navigation within the same session

## Heuristic Triggers
- `CountGenericSaaSLanguage`: Detects vague filter labels like "Filter by" without specifics
- `CountNewTopLevelComponents`: Detects ad-hoc filter implementations instead of shared filter components
- `FindArbitrarySpacing`: Flags inconsistent spacing between filter chips and result areas

## Scoring Axes Most Affected
- ux_clarity: Users must see what filters are active and how they affect results
- component_reuse: Filter chips and controls should be shared across all filtered views
