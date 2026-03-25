# Skill: Sort Control Patterns

## Category
UX

## Description
Sort controls let users reorder content to match their current task — finding the most recent item, the highest priority, or an alphabetical entry. Effective sort controls are visible, clearly labeled, and indicate the current sort direction. A column header that sorts when clicked should show an arrow indicating ascending or descending order. Dropdown sort selectors should display the current sort value, not just "Sort by" with the actual selection hidden behind a click. The default sort order should be the most useful for the context: chronological for activity feeds, alphabetical for directories, priority for task lists. Letting users sort by irrelevant fields (like sort by name on a list with no names) is confusing.

The AI agent should learn that sort state should persist within a session and be visually obvious. If a user sorts a list, navigates to an item, and returns, the sort should remain applied. Sort indicators — arrows, highlighted headers, or badge labels — must be immediately visible without hovering. Avoid offering sort options that make no sense for the data (e.g., sorting a list of images by name when filenames are UUIDs). When combined with filters, sort should apply to the filtered result set, not the full dataset. Avoid resetting sort when filters change unless the sort field is no longer valid for the filtered results.

## Expected Outcome
- Current sort field and direction are clearly indicated with visible arrows or labels
- Sort state persists across navigation within the same session
- Default sort order is contextually appropriate for the content type
- Sort controls are labeled in plain language matching the data fields

## Heuristic Triggers
- `CountGenericSaaSLanguage`: Detects vague sort labels like "Sort" without specifying the current field
- `FindHardcodedPaletteClasses`: Detects hardcoded colors for sort indicators instead of semantic tokens
- `CountRawPrimitiveAdditions`: Flags raw select elements instead of a shared sort component

## Scoring Axes Most Affected
- ux_clarity: Users must see what is sorted and in which direction
- component_reuse: Sort controls should be a shared component used consistently across the product
