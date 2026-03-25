# Skill: Keyboard Focus Management

## Category
Accessibility

## Description
Focus management determines how keyboard users navigate through an interface. This skill teaches the AI agent to implement logical tab order, visible focus indicators, and focus trapping within modals or dialogs. When focus is not managed intentionally, keyboard users can lose their place, tab into hidden content, or become trapped in overlays with no way to return to the main page. Each of these failures makes the application unusable without a mouse.

The agent should learn to use tabIndex strategically, to move focus programmatically when opening and closing modals, and to ensure that focus styles are visible and consistent. The natural DOM order should generally dictate tab order, but when visual order diverges from DOM order, the agent must reconcile the two so that keyboard navigation matches what sighted users experience. Focus management is one of the most commonly neglected areas of accessibility and one of the most impactful to get right.

## Expected Outcome
- Ensure all interactive elements are reachable via Tab key in a logical reading order
- Implement visible focus indicators using outline or ring utilities that meet contrast requirements
- Trap focus within modals and restore focus to the triggering element on close
- Avoid positive tabIndex values that disrupt the natural tab sequence

## Heuristic Triggers
- `FindOversizedSidebarClasses`: Large sidebars may contain focusable elements that trap or skip focus unexpectedly
- `HasClearHeadingStructure`: Heading hierarchy correlates with logical focus flow through page sections
- `CountAddedWrapperDivs`: Extra wrapper divs can break the tab order by inserting non-interactive containers between focusable elements

## Scoring Axes Most Affected
- layout_balance: Proper focus flow requires balanced layouts where interactive elements follow a predictable spatial order
- visual_noise: Visible focus indicators must stand out against the surrounding design without adding clutter
