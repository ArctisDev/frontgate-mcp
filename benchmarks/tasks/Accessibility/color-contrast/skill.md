# Skill: WCAG Color Contrast Requirements

## Category
Accessibility

## Description
Color contrast is one of the most objective and testable accessibility criteria. WCAG 2.1 requires a minimum contrast ratio of 4.5:1 for normal text and 3:1 for large text (18px bold or 24px regular). This skill teaches the AI agent to select foreground and background color combinations that meet these thresholds, to avoid placing text over images without sufficient overlay, and to check that UI components and graphical objects maintain at least a 3:1 ratio against adjacent colors.

The agent should learn to recognize when hardcoded color values from a palette may fail contrast checks against certain backgrounds, when opacity or gradients reduce effective contrast, and when color alone is used to convey meaning without a secondary indicator. Tools like contrast checkers can automate verification, but the agent should develop an intuition for which combinations are likely to pass and which need adjustment.

## Expected Outcome
- Ensure all body text meets a 4.5:1 contrast ratio against its background
- Ensure large text meets a 3:1 contrast ratio against its background
- Verify that UI components like borders and icons maintain 3:1 contrast against adjacent colors
- Never rely solely on color to convey information; pair it with text, icons, or patterns

## Heuristic Triggers
- `FindHardcodedPaletteClasses`: Hardcoded color classes may not have been tested against all background contexts
- `CountGenericSaaSLanguage`: Generic call-to-action text with low contrast ratios reduces conversion and accessibility
- `ExtractClassNames`: Utility classes that apply color without contrast verification can introduce violations

## Scoring Axes Most Affected
- token_adherence: Design tokens with baked-in contrast checks ensure accessibility is maintained as tokens are applied
- visual_noise: High contrast pairs can feel harsh if not balanced, so the agent must achieve contrast without visual clutter
