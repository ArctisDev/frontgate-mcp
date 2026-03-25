# Skill: Sidebar Efficiency

## Category
Layout

## Description
This skill teaches the AI agent to constrain sidebar width so it does not dominate the layout and starve the main content area of space. Sidebars are auxiliary navigation or context panels, and when they expand beyond their intended proportion, the entire page layout suffers. The agent should learn to cap sidebar widths at 300px for standard desktop layouts and use flexible units (like `min()` or `clamp()`) to ensure sidebars gracefully collapse or convert to drawers on smaller screens.

A common mistake is letting the sidebar grow to accommodate long navigation labels or deeply nested menu structures. Instead, the agent should use truncation, collapsible sections, or icon-only modes to keep the sidebar compact. When a sidebar must display rich content like filters or property panels, it should use internal scrolling rather than expanding its own width. The sidebar's width should be a token or CSS custom property that is referenced consistently, not hardcoded in multiple places.

The agent should also recognize when a sidebar is unnecessary entirely. Not every page benefits from a persistent sidebar. On pages with simple navigation or content-focused layouts, removing the sidebar entirely in favor of a top navigation or breadcrumb pattern improves density and reduces visual noise. The decision to include a sidebar should be deliberate and justified by the page's information architecture.

## Expected Outcome
- Sidebar widths are constrained to 300px or less on desktop layouts
- Sidebars use responsive collapse behavior on screens narrower than 1024px
- Sidebar width is defined as a token or custom property, not hardcoded inline
- Internal scrolling is used for overflow content rather than expanding sidebar width

## Heuristic Triggers
- `FindOversizedSidebarClasses`: Detects sidebar elements with widths exceeding 300px or using excessive flex-basis values
- `FindArbitrarySpacing`: Flags hardcoded width values in sidebar definitions that bypass design tokens
- `ExtractClassNames`: Identifies class names suggesting wide or unbounded sidebar configurations

## Scoring Axes Most Affected
- layout_balance: Properly sized sidebars maintain the content-to-chrome ratio that users expect
- density_control: Narrow sidebars preserve screen real estate for primary content
- template_likeness: Standard sidebar widths match established UI conventions and feel professional
