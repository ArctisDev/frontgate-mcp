# Skill: Breadcrumb Design

## Category
Component

## Description
Breadcrumbs provide contextual wayfinding in applications with deep navigation hierarchies, such as multi-level settings, nested resource views, or content management systems. When breadcrumbs are absent or inconsistently implemented, users lose spatial awareness and struggle to navigate back to parent contexts. A well-designed breadcrumb component handles separator styling, truncation for long paths, responsive collapsing on mobile, and proper semantic markup (`nav` with `aria-label="Breadcrumb"` and `ol` with `li` elements). Ad-hoc breadcrumb implementations often skip these accessibility requirements and use inconsistent separators or link styling.

This skill teaches the agent to add breadcrumbs to views that have a clear parent-child hierarchy, using the project's Breadcrumb component rather than raw markup. The agent should learn to populate breadcrumbs from the current route or resource context, ensuring each level is a clickable link except the final (current) item. When paths are too long for mobile screens, the agent should use the component's built-in truncation or collapsing behavior. This improves navigation efficiency and reduces the number of users who feel "lost" in deeply nested views.

## Expected Outcome
- Breadcrumbs are present on views with two or more levels of navigation depth
- Each breadcrumb level is a link except the current page, which is plain text
- Breadcrumb separators and styling use the shared component's design tokens
- Semantic markup (nav, ol, li) and aria labels are provided by the component

## Heuristic Triggers
- `HasClearHeadingStructure`: Checks whether pages with breadcrumbs also have proper heading hierarchy for accessibility
- `CountExistingComponentReuse`: Measures whether the Breadcrumb component is used instead of custom link lists
- `CountRawPrimitiveAdditions`: Detects raw `<nav>` or `<ol>` elements used for breadcrumb-like navigation
- `ExtractClassNames`: Identifies custom breadcrumb classes that diverge from the design system's navigation tokens

## Scoring Axes Most Affected
- ux_clarity: Breadcrumbs dramatically improve spatial orientation in deep navigation hierarchies
- component_reuse: Using the shared Breadcrumb component ensures consistent separator and truncation behavior
- token_adherence: Ad-hoc breadcrumb styles bypass tokens for link color, separator icon, and typography
- template_likeness: Products with breadcrumbs feel more navigable and professionally structured
