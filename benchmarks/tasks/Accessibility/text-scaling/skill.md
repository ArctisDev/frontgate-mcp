# Skill: Supporting Text Scaling

## Category
Accessibility

## Description
Users with low vision often increase their browser's default font size or use operating system scaling to make text readable. Applications that use fixed pixel units for font sizes and container dimensions break at these higher scales, causing text to overflow, overlap, or become clipped. This skill teaches the AI agent to use relative units like rem and em for font sizes, line heights, and spacing, and to design layouts that reflow gracefully at up to 200% browser zoom without horizontal scrolling.

The agent should learn to avoid hardcoding pixel values for typography, to use container queries or flexible grid systems that adapt to available space, and to test layouts at various zoom levels. Text scaling support is not just an accessibility requirement but a responsiveness concern that benefits users on high-DPI displays, projectors, and small screens alike.

## Expected Outcome
- Use rem or em units for all font sizes, line heights, and text-related spacing
- Ensure layouts reflow without horizontal scrolling at 200% browser zoom
- Avoid fixed-width containers that clip text when the base font size increases
- Test with browser zoom at 150% and 200% to verify no content is lost or overlapped

## Heuristic Triggers
- `FindArbitrarySpacing`: Arbitrary pixel spacing values may not scale proportionally with font size changes
- `FindOversizedSidebarClasses`: Oversized sidebars with fixed widths will clip content at higher zoom levels
- `CountRawPrimitiveAdditions`: Raw pixel values added to inline styles bypass responsive scaling

## Scoring Axes Most Affected
- layout_balance: Text scaling requires layouts that maintain balance at multiple viewport and font sizes
- density_control: Relative units naturally enforce proportional density rather than fixed pixel crowding
