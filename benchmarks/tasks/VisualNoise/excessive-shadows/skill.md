# Skill: Limiting Excessive Shadow Usage

## Category
VisualNoise

## Description
Box shadows create a sense of elevation and depth, making elements feel interactive or important. However, when shadows appear on every card, button, sidebar panel, and modal simultaneously, the elevation system collapses — nothing feels elevated if everything is. Agents frequently apply shadow utilities broadly because shadows add a polished look, but indiscriminate use flattens the visual hierarchy.

This skill teaches using shadows sparingly and with intent. Reserve shadow elevation for elements that genuinely float above the surface: dropdowns, modals, popovers, and primary action cards. Background sections and static content panels should sit flat or use only the subtlest shadow. A well-designed UI has two to three elevation levels at most, not a shadow on every element.

## Expected Outcome
- Shadow usage limited to two or three distinct elevation levels
- Static content sections rendered without box shadows
- Dropdowns, modals, and popovers retain elevated shadows
- Clear distinction between interactive and static surfaces

## Heuristic Triggers
- `FindHardcodedPaletteClasses`: Detects shadow utility classes applied broadly
- `CountRawPrimitiveAdditions`: Can flag raw shadow values added outside the design token system

## Scoring Axes Most Affected
- visual_noise: Fewer shadows reduce competing visual accents across the page
- token_adherence: Using elevation tokens instead of arbitrary shadow values maintains system consistency
