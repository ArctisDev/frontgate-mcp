# Skill: Modal Placement

## Category
Layout

## Description
This skill teaches the AI agent to implement modals and dialogs with proper positioning, overlay management, and focus handling. A modal should be centered within the viewport using fixed or sticky positioning with a semi-transparent backdrop that blocks interaction with the underlying page. The agent should use native `<dialog>` elements where possible, as they provide built-in focus trapping, backdrop styling, and escape key handling. When a custom modal implementation is necessary, the agent must ensure it replicates these accessibility behaviors rather than creating a simple absolutely-positioned div that lacks keyboard and screen reader support.

The backdrop (overlay) should be a separate element with a neutral dark or blurred background, not a border or shadow on the modal itself. The agent should ensure the backdrop covers the entire viewport and prevents scroll on the body while the modal is open. Modal content should have a reasonable `max-width` (typically 480-640px for standard dialogs, up to 960px for complex forms) and should not exceed the viewport height. If a modal's content is too tall, it should scroll internally rather than pushing the modal off-screen.

The agent should avoid positioning modals relative to trigger elements (tooltip-like behavior) unless specifically implementing a popover. True modals are viewport-centered and demand the user's full attention. Positioning a modal near a button or link creates ambiguity about whether it is a modal, a dropdown, or a tooltip. The agent should also ensure that only one modal is open at a time and that closing the modal returns focus to the element that triggered it.

## Expected Outcome
- Modals are viewport-centered with fixed or sticky positioning
- A semi-transparent backdrop covers the full viewport and blocks interaction with background content
- Modal content scrolls internally if it exceeds viewport height, rather than overflowing
- Focus is trapped within the modal while open and returned to the trigger element on close

## Heuristic Triggers
- `FindArbitrarySpacing`: Detects hardcoded positioning values (top, left, transform) in modal implementations that bypass layout tokens
- `CountAddedWrapperDivs`: Flags extra wrapper divs around modal content that could be consolidated
- `CountRawPrimitiveAdditions`: Identifies new raw HTML elements for modal overlays that should use existing dialog primitives

## Scoring Axes Most Affected
- ux_clarity: Properly centered modals with backdrops clearly communicate their priority and interrupt the user's flow intentionally
- layout_balance: Viewport-centered modals maintain visual equilibrium regardless of underlying page scroll position
- component_reuse: Using a standard modal component ensures consistent behavior and reduces bespoke implementations
