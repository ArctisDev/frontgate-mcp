# Skill: Logical Information Grouping

## Category
Hierarchy

## Description
Related information should be visually grouped together so users can process it as a unit. Agents often scatter related fields, metrics, or content across separate sections, or interleave unrelated items in the same container. This forces users to hunt across the page to assemble a complete picture of a topic, increasing cognitive load and reducing comprehension.

This skill teaches placing related information in proximity with clear visual boundaries separating it from unrelated content. A user profile section should contain name, email, role, and status together. A metrics dashboard should group revenue, users, and growth in a single region. Use consistent spacing, subtle background shifts, or lightweight containers to define group boundaries without heavy visual separators. The user should never have to search the page to find information that belongs together.

## Expected Outcome
- Related fields and content physically adjacent in the layout
- Clear visual boundaries between distinct information groups
- No interleaving of unrelated content within the same container
- Users can process each group independently without cross-referencing

## Heuristic Triggers
- `HasClearHeadingStructure`: Section headings should map to logical content groups
- `FindArbitrarySpacing`: Inconsistent spacing between groups suggests poor logical organization

## Scoring Axes Most Affected
- layout_balance: Well-grouped content distributes evenly across the available space
- ux_clarity: Logical grouping reduces the time needed to find related information
