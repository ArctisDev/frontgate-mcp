# Skill: Minimum Touch Target Sizes

## Category
Accessibility

## Description
Interactive elements that are too small to tap reliably cause frustration for users on touch devices and are particularly problematic for users with motor impairments. WCAG 2.5.8 requires that touch targets have a minimum area of 24 by 24 CSS pixels, with a best practice recommendation of 44 by 44 pixels. This skill teaches the AI agent to ensure that buttons, links, form controls, and any other interactive elements meet these minimum dimensions, including their padding and hit area.

The agent should learn to detect when small icons or text links are presented as tappable elements without sufficient surrounding padding, when interactive elements in dense toolbars are crowded too close together, and when CSS reduces an element's effective tap target through overflow hidden or negative margins. Ensuring adequate touch target size benefits all mobile users, not just those with accessibility needs.

## Expected Outcome
- Ensure all interactive elements have a minimum tap target of 44x44 CSS pixels
- Add padding or min-width and min-height to small icon buttons that would otherwise fall below the threshold
- Maintain at least 8px of spacing between adjacent interactive elements to prevent mis-taps
- Avoid shrinking hit areas with CSS properties like overflow hidden or negative margins on interactive containers

## Heuristic Triggers
- `FindArbitrarySpacing`: Tight spacing values between interactive elements may indicate insufficient touch target separation
- `CountCardMentions`: Card actions and icon buttons inside cards are common sources of undersized touch targets
- `FindOversizedSidebarClasses`: Oversized sidebars may contain dense lists of links or buttons with small tap targets

## Scoring Axes Most Affected
- density_control: Enforcing minimum touch targets requires balancing density with tappable area
- layout_balance: Adequate spacing between interactive elements contributes to a balanced and comfortable layout
