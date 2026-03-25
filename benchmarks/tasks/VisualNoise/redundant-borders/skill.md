# Skill: Eliminating Redundant Borders

## Category
VisualNoise

## Description
Borders are a common default for separating content, but when applied to every card, list item, sidebar section, and footer block, they create a grid of lines that fragments the visual field. Agents tend to overuse `border` utilities because they are a safe, explicit way to delineate areas. The result is a UI that looks like a wireframe rather than a finished product.

This skill teaches using background color shifts, consistent spacing, and subtle elevation changes to create visual separation instead of drawing lines everywhere. A slight background tint on alternating sections or adequate whitespace between groups communicates structure without the harshness of borders. Reserve visible borders for interactive elements like inputs and tables where they serve a functional purpose.

## Expected Outcome
- Minimal use of visible border lines across the layout
- Visual separation achieved through background contrast and spacing
- Borders reserved for form inputs, tables, and interactive surfaces
- Reduced overall line density in the rendered UI

## Heuristic Triggers
- `FindArbitrarySpacing`: Can reveal spacing compensating for missing borders that should be handled by proper grouping
- `FindHardcodedPaletteClasses`: Detects hardcoded border color classes used decoratively

## Scoring Axes Most Affected
- visual_noise: Fewer lines means a calmer, less cluttered visual field
- spacing_consistency: Proper spacing replaces borders as the primary grouping mechanism
