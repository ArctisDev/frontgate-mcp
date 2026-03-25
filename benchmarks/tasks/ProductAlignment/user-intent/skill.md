# Skill: Designing for User Intent

## Category
ProductAlignment

## Description
Every screen should make the user's most likely next action obvious and accessible. This means the primary action should be visually prominent, the path to secondary actions should be clear but not competing, and the interface should reduce the total steps required for common tasks. This skill teaches AI agents to design with intent-awareness, prioritizing what users actually want to accomplish over what the data model technically supports.

Intent-aware design requires understanding action frequency and importance. The "Create New" button might be the most important action on a list page. The "Save" action might need to be accessible without scrolling on a form. Destructive actions should require confirmation but not be hidden so deeply that users cannot find them. Agents must learn to identify the primary intent for each screen from the product context and ensure the UI reflects that hierarchy.

A common failure is treating all actions equally: giving every button the same visual weight, placing all operations in a dropdown menu, or requiring the same number of clicks for both frequent and rare operations. Intent-aware design creates a clear visual hierarchy that guides users toward their goal.

## Expected Outcome
- The primary action on each screen is visually distinct and immediately accessible
- Frequently used actions require fewer clicks than infrequent ones
- The interface layout guides the eye toward the most important content and actions first
- Secondary and destructive actions are accessible but do not compete with primary actions

## Heuristic Triggers
- `CountCardMentions`: Can reveal when card layouts obscure primary actions behind uniform containers
- `CountAddedWrapperDivs`: Excessive nesting can bury primary actions under multiple interaction layers

## Scoring Axes Most Affected
- ux_clarity: Intent-driven design makes next steps immediately obvious
- product_alignment: Shows understanding of what users actually want to do with the product
- density_control: Intent hierarchy informs how much space different elements deserve
