# Skill: Tooltip Usage

## Category
Component

## Description
Tooltips provide supplementary information on hover or focus, but they are frequently overused as a crutch for unclear UI. When icons without labels rely on tooltips to explain their function, users on touch devices (where hover does not exist) are left guessing. When tooltips repeat the visible label text verbatim, they add noise without value. A well-applied tooltip strategy reserves tooltips for clarifying truncated content, explaining icon-only actions, and providing keyboard shortcuts or additional context that does not fit inline. The tooltip component itself should handle positioning, delay timing, and accessibility (`aria-describedby` or `role="tooltip"`).

This skill teaches the agent to evaluate whether a tooltip adds genuine value or whether the information should be visible inline instead. The agent should learn to use the project's Tooltip component for consistent positioning, delay, and styling. Tooltips should not be used as the sole means of conveying critical information. When an icon button's purpose is ambiguous, the agent should prefer visible text labels or use the Tooltip component with proper `aria` attributes. This ensures that tooltips enhance the experience for pointer users without excluding touch and screen reader users.

## Expected Outcome
- Tooltips are used to clarify icon-only actions, not to repeat visible text
- Critical information is displayed inline rather than hidden behind a tooltip
- Tooltip positioning, delay, and styling use the shared component's defaults
- ARIA attributes are properly set for screen reader accessibility

## Heuristic Triggers
- `CountGenericSaaSLanguage`: Detects tooltip content that uses vague or placeholder language (e.g., "Click here for more info")
- `CountRawPrimitiveAdditions`: Flags custom tooltip markup built from title attributes or ad-hoc hover divs
- `CountExistingComponentReuse`: Measures whether the Tooltip component is used consistently
- `CountCardMentions`: Identifies cards where tooltips hide important information that should be visible on the card surface

## Scoring Axes Most Affected
- ux_clarity: Purposeful tooltips clarify; overused tooltips obscure information from users
- component_reuse: Using the shared Tooltip component ensures consistent positioning and behavior
- visual_noise: Excessive tooltip triggers (dotted underlines, info icons) clutter the interface
- density_control: Replacing inline labels with tooltips can artificially reduce apparent density while harming usability
