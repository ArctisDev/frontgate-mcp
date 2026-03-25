# Skill: Reducing Card-Over-Card Patterns

## Category
VisualNoise

## Description
Card containers are useful for grouping related content, but they become visual noise when nested inside other cards or layered excessively. Every card shell adds a border, background, and shadow that competes for attention. When an agent wraps every section in a card, the UI starts to feel like a stack of boxes rather than a coherent layout.

This skill teaches that not every content group needs a card shell. Spacing alone can establish visual grouping without the overhead of borders and backgrounds. Reserve card containers for interactive or elevated surfaces like clickable items, modals, or dashboard widgets. For static content sections, use consistent padding and margin to create rhythm instead of wrapping everything in bordered boxes.

## Expected Outcome
- Fewer nested card components in the output
- Content sections grouped by spacing rather than container borders
- Cards reserved for interactive or elevated surfaces only
- Reduced CountCardMentions across the generated UI

## Heuristic Triggers
- `CountCardMentions`: Detects excessive card container usage across the layout
- `CountAddedWrapperDivs`: Identifies unnecessary wrapper elements added around content

## Scoring Axes Most Affected
- visual_noise: Reducing card shells lowers the number of competing visual boundaries
- density_control: Fewer containers means content occupies space more efficiently
