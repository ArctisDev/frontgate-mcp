# Skill: Designing for Scan Patterns

## Category
Hierarchy

## Description
Users do not read web pages — they scan them. Eye-tracking research consistently shows that users follow F-patterns (scanning the top, then down the left) and Z-patterns (scanning in a zigzag) depending on content density. Agents that place critical information in the bottom-right of a dense layout or bury key actions below the fold are ignoring how humans naturally consume interfaces.

This skill teaches placing key information and actions along natural eye movement paths. The top of the page and left-aligned content receive the most attention in F-pattern layouts. For Z-pattern pages, the top-left and bottom-right corners are prime positions. Headings, primary actions, and key metrics should occupy these high-attention zones. Secondary information can fill the lower-attention areas. Structure the content so the scanning path leads the user from context to action without backtracking.

## Expected Outcome
- Key information and primary actions placed in high-attention zones (top, left, corners)
- Content structured to follow a coherent scanning path
- No critical information buried in low-attention areas
- Page layout supports both F-pattern and Z-pattern scanning depending on density

## Heuristic Triggers
- `HasClearHeadingStructure`: Headings anchor the scanning path at each level
- `FindOversizedSidebarClasses`: Oversized sidebars can disrupt the natural scanning flow

## Scoring Axes Most Affected
- ux_clarity: Content aligned with scanning patterns is found faster
- layout_balance: Scan-friendly layouts distribute attention across the page naturally
