# Skill: Search Experience Design

## Category
UX

## Description
Search is often the primary way users navigate large content sets, and its implementation directly impacts product usability. A good search experience starts with an accessible, prominent input field that communicates what can be searched. Debounced search — waiting 200-400ms after the user stops typing before executing — prevents excessive API calls while maintaining responsiveness. Results should appear in a clear, scannable format with the matching term highlighted. When no results are found, the empty state should suggest alternatives: check spelling, try broader terms, or browse by category. Search that returns "No results" with no guidance is a dead end.

The AI agent should learn that search needs to handle every state gracefully. Loading state with a skeleton or spinner during the query. Results state with clear highlighting and relevant metadata. Empty results state with helpful suggestions. Error state with a retry option. Recent searches and suggested queries improve discoverability before the user types. Filters should be available but not required — the search should work well without them. Avoid requiring users to press Enter to search when real-time results would be more helpful. Keyboard navigation through results (arrow keys, Enter to select) is essential for power users.

## Expected Outcome
- Search input is debounced (200-400ms delay) to prevent excessive requests
- Empty results show actionable suggestions, not just "No results found"
- All search states are handled: loading, results, empty, and error
- Search results highlight the matching term for quick scanning

## Heuristic Triggers
- `CountGenericSaaSLanguage`: Detects unhelpful empty search messages like "No results found"
- `CountRawPrimitiveAdditions`: Flags raw input elements instead of a shared search component
- `FindArbitrarySpacing`: Detects layout shifts when results appear or disappear

## Scoring Axes Most Affected
- ux_clarity: Users must understand what they searched for and why results matched
- component_reuse: Search should use a shared component for consistency across the product
