# Skill: Matching Visual Weight to Importance

## Category
Hierarchy

## Description
Every page has a primary purpose — the main action the user should take, or the most important information they need to see. When agents generate UI, they often treat all elements with equal visual weight: every button gets the same style, every section gets the same heading size, every card gets the same treatment. This flattens the hierarchy and forces the user to read everything to understand what matters most.

This skill teaches making primary actions and key information visually dominant. The main call-to-action should be the most visually prominent element on the page — larger, bolder, or more colorful than secondary actions. Key metrics or status information should occupy the most prominent position and use the largest type. Secondary and tertiary information should recede through smaller type, lighter color, or less prominent positioning. The user should be able to identify the page's purpose within two seconds of looking at it.

## Expected Outcome
- Primary action visually dominant over all secondary actions
- Key information occupies the most prominent position on the page
- Secondary content visually recedes through size, color, or position
- Page purpose is immediately clear from visual hierarchy alone

## Heuristic Triggers
- `CountGenericSaaSLanguage`: Generic labels often indicate missing prioritization of content
- `HasClearHeadingStructure`: Heading hierarchy supports content priority mapping

## Scoring Axes Most Affected
- ux_clarity: Visual priority tells users what matters without reading everything
- product_alignment: Prominent primary actions reinforce the page's core purpose
