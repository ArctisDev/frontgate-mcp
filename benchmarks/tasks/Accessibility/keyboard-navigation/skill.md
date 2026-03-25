# Skill: Full Keyboard Operability

## Category
Accessibility

## Description
Every interactive operation in a web application must be achievable without a mouse. This skill teaches the AI agent to implement keyboard shortcuts and to ensure that all behaviors triggered by click or hover events also fire on Enter, Space, or arrow key presses as appropriate. Dropdown menus should open with Enter or Space, sliders should move with arrow keys, and custom components like tabs should support both horizontal arrow navigation and tab-based focus movement.

The agent should learn to avoid trapping keyboard users in dead ends, to provide keyboard equivalents for drag-and-drop interactions, and to document keyboard shortcuts where discoverability is a concern. Keyboard operability is a legal requirement under WCAG and a practical necessity for power users, developers, and anyone who cannot use a pointing device. The agent should test every interactive component for keyboard accessibility as a matter of course.

## Expected Outcome
- Implement Enter and Space key handlers on custom interactive elements that are not native buttons or links
- Support arrow key navigation within composite widgets like menus, tab lists, and toolbars
- Ensure drag-and-drop alternatives exist for keyboard-only users
- Provide visible keyboard shortcut hints or documentation for complex interactions

## Heuristic Triggers
- `CountCardMentions`: Cards with click handlers but no keyboard equivalents indicate missing keyboard operability
- `CountAddedWrapperDivs`: Wrapper divs used as click targets often lack keyboard support and should be replaced with buttons
- `FindOversizedSidebarClasses`: Oversized interactive regions may hide focusable sub-elements that need proper keyboard routing

## Scoring Axes Most Affected
- ux_clarity: Keyboard-operable interfaces communicate interaction patterns consistently across input methods
- density_control: Proper keyboard support requires clear visual hierarchy, which naturally improves content density
