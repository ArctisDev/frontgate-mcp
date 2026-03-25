# Skill: Balancing Visual Weight Across the Page

## Category
Hierarchy

## Description
Visual weight is the perceived heaviness of an element based on its size, color, contrast, and density. A page where all the heavy elements cluster in one corner feels unbalanced and draws the eye unevenly. Agents often concentrate bold headings, colored badges, and large cards in the upper-left or center of the layout, leaving the rest of the page visually flat and uninviting.

This skill teaches distributing visual weight intentionally across the page. Bold and colorful elements should be spread to create a balanced composition — if a large card sits on the left, a cluster of smaller but equally vivid elements can counterbalance it on the right. Use lighter elements (muted text, subtle backgrounds, thin borders) to fill transitional space between heavy zones. The goal is a layout where the eye moves naturally across the full page rather than fixating on one dense region.

## Expected Outcome
- Bold, large, and colorful elements distributed across the layout rather than clustered
- Visual counterbalance between heavy and light regions of the page
- Transitional zones use lighter elements to bridge heavy areas
- Eye movement naturally traverses the full page width

## Heuristic Triggers
- `FindOversizedSidebarClasses`: Detects sidebars that dominate the layout's visual weight
- `CountNewTopLevelComponents`: Flags top-level component additions that may unbalance the layout

## Scoring Axes Most Affected
- layout_balance: Even weight distribution creates a visually stable composition
- template_likeness: Balanced layouts more closely resemble polished production designs
