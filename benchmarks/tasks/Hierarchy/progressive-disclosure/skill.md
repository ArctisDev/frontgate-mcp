# Skill: Progressive Disclosure of Information

## Category
Hierarchy

## Description
Presenting all information at once overwhelms users and buries the content they actually need. Agents often generate every field, setting, and detail visible on the initial render because completeness feels thorough. But a settings page with forty visible fields or a dashboard with twelve charts forces the user to process everything simultaneously, even when they only need a fraction of it at any given time.

This skill teaches revealing complexity only when the user requests it. Use collapsible sections for secondary settings, tabs for related-but-distinct content areas, and expandable rows for detail views. The initial state should show the most commonly needed information and provide clear affordances for accessing deeper detail. Progressive disclosure respects the user's cognitive bandwidth and keeps the interface approachable for new users while remaining powerful for advanced ones.

## Expected Outcome
- Secondary content hidden behind expandables, tabs, or accordions
- Initial page state shows only the most commonly needed information
- Clear affordances (chevrons, "Show more" links) indicate hidden content
- Advanced detail accessible without leaving the current page context

## Heuristic Triggers
- `CountNewTopLevelComponents`: Excessive top-level components may indicate missing progressive disclosure
- `CountCardMentions`: Card-heavy layouts without collapse options lack disclosure patterns

## Scoring Axes Most Affected
- density_control: Progressive disclosure reduces initial content density
- ux_clarity: Showing only relevant information reduces cognitive overload
