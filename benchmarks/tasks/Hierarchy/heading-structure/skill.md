# Skill: Clear Heading Structure

## Category
Hierarchy

## Description
Heading elements (h1 through h6) are the backbone of document structure, yet agents frequently use them as font-size shortcuts rather than structural markers. An agent might skip from h1 to h4, use multiple h1 elements on a single page, or apply heading tags to text that is not a section title simply because the desired font size matches. This breaks both accessibility and the user's ability to scan the page hierarchy.

This skill teaches using headings strictly for document structure: one h1 for the page title, h2 for major sections, h3 for subsections, and so on in strict descending order. Font size adjustments that are not structural should use utility classes or design tokens instead of heading tags. A well-structured heading tree enables screen readers to navigate the page and helps sighted users build a mental model of the content layout at a glance.

## Expected Outcome
- Single h1 per page representing the primary page title
- Headings descend in strict order without skipping levels
- Heading tags used only for structural section titles, not for font sizing
- Screen reader navigation produces a coherent page outline

## Heuristic Triggers
- `HasClearHeadingStructure`: Validates that headings follow a proper hierarchical order
- `ExtractClassNames`: Surfaces classes that override heading semantics for styling purposes

## Scoring Axes Most Affected
- typographic_hierarchy: Proper heading levels create a clear typographic ladder
- ux_clarity: Structured headings improve scanability and accessibility
