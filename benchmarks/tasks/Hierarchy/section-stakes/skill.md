# Skill: Establishing Clear Section Boundaries

## Category
Hierarchy

## Description
A page with no clear section boundaries feels like a wall of content. Users cannot tell where one topic ends and another begins, forcing them to read linearly rather than jumping to the section they need. Agents often generate continuous streams of cards, lists, and forms without demarcating where distinct functional areas begin and end, producing layouts that feel like infinite scroll rather than structured pages.

This skill teaching establishing clear section boundaries through a combination of section headers, background color alternation, and generous vertical spacing. Each major functional area should have a visible heading, a consistent background, and enough whitespace above and below to signal a transition. Background color shifts between adjacent sections (even subtle ones like white to light gray) create powerful boundaries without lines. The user should be able to identify every section on the page by its header alone.

## Expected Outcome
- Each major functional area has a visible section header
- Background color alternates between adjacent sections
- Generous vertical spacing separates major page regions
- Users can identify and navigate to any section by header alone

## Heuristic Triggers
- `HasClearHeadingStructure`: Section boundaries should map to heading hierarchy
- `FindArbitrarySpacing`: Inconsistent section spacing suggests missing boundary patterns

## Scoring Axes Most Affected
- layout_balance: Clear sections help distribute content evenly across the page
- ux_clarity: Section boundaries let users skip directly to relevant content
