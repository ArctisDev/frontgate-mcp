# Skill: Page Structure

## Category
Layout

## Description
This skill teaches the AI agent to establish clear, predictable page structures using semantic landmarks like header, main, footer, and nav. A well-structured page follows a consistent pattern: a header at the top for branding and primary navigation, a main content area that fills the available space, and a footer at the bottom for secondary links and legal information. This three-part structure is a convention that users understand intuitively, and deviating from it without strong justification creates disorientation. The agent should use semantic HTML elements (`<header>`, `<main>`, `<footer>`, `<nav>`) rather than generic `<div>` tags, which improves accessibility and makes the page outline parseable by assistive technologies.

Within the main content area, the agent should use a consistent layout hierarchy: page title, optional action bar, primary content, and optional sidebar. This hierarchy should be established through CSS Grid or Flexbox layout on the page shell, not through ad-hoc margin and padding on individual sections. The page shell (the top-level layout component) should define the structural skeleton, and page-specific content should fill the slots without overriding the shell's layout rules.

The agent should avoid duplicating page structure definitions across routes. If every page defines its own header, footer, and main layout, changes to the structure require updating every page. Instead, the agent should use a shared layout component that wraps all pages and provides the structural skeleton. Page-specific content fills the main slot. This pattern ensures structural consistency and makes global layout changes trivial.

## Expected Outcome
- Pages use semantic landmark elements (header, main, footer, nav) instead of generic divs
- A shared layout component provides the page shell, and pages fill content slots
- The page structure is consistent across routes with the header, content, and footer in the same positions
- The layout shell uses CSS Grid or Flexbox for structural positioning, not ad-hoc margins

## Heuristic Triggers
- `HasClearHeadingStructure`: Detects whether the page has a logical heading hierarchy that complements its structural landmarks
- `CountNewTopLevelComponents`: Flags when new page components duplicate structural elements that should come from a shared layout
- `CountAddedWrapperDivs`: Identifies extra wrapper divs around landmark elements that break semantic structure

## Scoring Axes Most Affected
- template_likeness: Consistent page structure matches established web conventions and feels professional
- ux_clarity: Semantic landmarks help users navigate the page and understand its information hierarchy
- layout_balance: A shared layout shell ensures structural consistency across all routes
