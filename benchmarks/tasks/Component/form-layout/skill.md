# Skill: Form Layout

## Category
Component

## Description
Form layout determines how efficiently users can complete data entry tasks. Inconsistent label placement (sometimes above, sometimes beside), uneven spacing between fields, and misaligned error messages create friction that slows users down and increases abandonment. A well-structured form follows predictable patterns: labels sit at a consistent position relative to their inputs, field groups are separated by uniform spacing, error states appear in a fixed location, and submit actions are visually anchored at a predictable spot. These conventions reduce the user's need to reorient between forms.

This skill teaches the agent to construct forms using established layout patterns rather than ad-hoc markup. The agent should learn to apply consistent vertical rhythm between fields (typically using the design system's spacing scale), position labels according to the project's convention (top-aligned for mobile-first, left-aligned for dense admin panels), and use shared form wrapper components that handle spacing and error placement. When building new forms, the agent should compose existing field and layout primitives instead of writing raw `<form>` elements with scattered inline styles, ensuring that every form in the product feels like it belongs to the same system.

## Expected Outcome
- Label placement is consistent across all form fields (top-aligned or left-aligned per project convention)
- Spacing between fields follows the design system's vertical rhythm scale
- Error messages appear in a uniform position relative to their associated input
- Form submit/cancel actions are consistently placed and styled

## Heuristic Triggers
- `FindArbitrarySpacing`: Detects inconsistent margin/padding between form fields that breaks vertical rhythm
- `CountAddedWrapperDivs`: Flags excessive wrapper divs around form groups that introduce layout fragility
- `ExtractClassNames`: Identifies ad-hoc form layout classes that bypass the project's spacing tokens
- `HasClearHeadingStructure`: Checks whether form sections have proper heading hierarchy for accessibility

## Scoring Axes Most Affected
- spacing_consistency: Uniform field spacing is the backbone of professional form layout
- ux_clarity: Predictable label and error placement reduces user confusion during data entry
- template_likeness: Well-structured forms look like they belong to a polished product, not a prototype
- layout_balance: Proper alignment of labels, inputs, and actions creates visual harmony
