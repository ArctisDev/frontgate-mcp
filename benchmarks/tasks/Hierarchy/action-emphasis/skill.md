# Skill: Clear Primary and Secondary Action Differentiation

## Category
Hierarchy

## Description
When a page presents multiple buttons or actions, the user needs to instantly recognize which action is primary, which is secondary, and which is tertiary. Agents frequently render all buttons with the same solid variant, or worse, use color alone to differentiate without establishing a clear visual hierarchy. This creates decision paralysis — the user cannot tell which action the interface is recommending.

This skill teaches a consistent action variant system: primary actions use a filled or solid button style, secondary actions use an outline or ghost style, and destructive or tertiary actions use text-only or muted styles. The visual weight difference between variants should be significant enough that the primary action is unmistakable. Apply this consistently across the entire application so users learn the pattern and can rely on it.

## Expected Outcome
- Primary actions rendered with filled/solid button style
- Secondary actions rendered with outline or ghost button style
- Clear visual weight difference between action variants
- Consistent variant usage across all pages and components

## Heuristic Triggers
- `CountCardMentions`: Card-heavy layouts often bury action differentiation
- `CountGenericSaaSLanguage`: Generic button labels reduce action clarity

## Scoring Axes Most Affected
- ux_clarity: Distinct action styles eliminate ambiguity about what to do next
- template_likeness: Consistent button variants match production design system patterns
