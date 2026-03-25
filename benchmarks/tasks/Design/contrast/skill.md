# Skill: Color Contrast and Visual Emphasis

## Category
Design

## Description
Color contrast governs both accessibility and visual hierarchy. Agents that use low-contrast text on backgrounds (e.g., light gray text on a white surface, or a muted color on a similar-toned card) create interfaces that are difficult to read, especially for users with visual impairments. WCAG 2.1 requires a minimum 4.5:1 contrast ratio for normal text and 3:1 for large text. Beyond compliance, contrast is the primary mechanism for establishing what matters most on a page—high-contrast elements draw the eye, while low-contrast elements recede.

This skill teaches the agent to use contrast purposefully. Primary actions (CTAs, submit buttons) should have the highest contrast against their background. Secondary actions should be visually subordinate through reduced contrast or outline styles. Body text must always meet WCAG AA standards. The agent should use the project's semantic color tokens (foreground, muted-foreground, accent-foreground) which are pre-calibrated for contrast, rather than picking arbitrary colors that may fail accessibility checks. When the agent needs to emphasize content, it should increase contrast rather than adding borders, badges, or other visual noise.

## Expected Outcome
- Text elements meet WCAG AA contrast ratios against their backgrounds (4.5:1 for body text, 3:1 for large text)
- Primary interactive elements have the highest contrast, secondary elements are visually subordinate
- The agent uses semantic foreground tokens (text-foreground, text-muted-foreground) rather than hardcoded gray values
- Visual emphasis is achieved through contrast, not through adding extra decorative elements

## Heuristic Triggers
- `FindHardcodedPaletteClasses`: Detects hardcoded color classes that may not meet contrast requirements against the current background
- `ExtractClassNames`: Surfaces color and background class combinations to audit for contrast compliance
- `CountRawPrimitiveAdditions`: Flags new color values introduced outside the semantic token system that may have untested contrast ratios

## Scoring Axes Most Affected
- ux_clarity: Proper contrast ensures content is readable and interactive elements are identifiable
- token_adherence: Using semantic foreground tokens ensures contrast is maintained by the design system
- visual_noise: Achieving emphasis through contrast reduces the need for additional visual decorators
