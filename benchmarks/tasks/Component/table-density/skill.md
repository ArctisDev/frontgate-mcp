# Skill: Table Density

## Category
Component

## Description
Tables are critical for data-heavy interfaces, and their layout quality directly impacts how much information users can scan at a glance. Poorly structured tables waste vertical space with oversized row heights, use inconsistent column widths, and apply arbitrary padding that disrupts the data grid. A well-calibrated table component balances readability with density, ensuring that headers are visually distinct, rows have appropriate hover states, and columns size proportionally to their content. Overly padded tables force unnecessary scrolling, while cramped tables reduce legibility.

This skill teaches the agent to evaluate table implementations for optimal data density. The agent should learn to detect excessive padding and margin on table cells, oversized row heights that reduce visible data, and columns that do not respect content-proportional sizing. When building new table views, the agent should reuse the project's existing `Table` or `DataTable` component with its density presets (compact, default, comfortable) rather than writing raw `<table>` markup with ad-hoc styles. This ensures that data presentation remains efficient and consistent across dashboards and list views.

## Expected Outcome
- Table rows use project-defined density presets rather than arbitrary height values
- Column widths are proportional to content, avoiding fixed widths that truncate or waste space
- Cell padding follows the design system's spacing scale instead of magic numbers
- Header styling and sticky behavior use the shared Table component's built-in API

## Heuristic Triggers
- `FindArbitrarySpacing`: Detects custom padding/margin on `td`, `th` elements that override the table component's defaults
- `CountRawPrimitiveAdditions`: Flags raw `<table>` elements used when a Table component is available
- `FindOversizedSidebarClasses`: Identifies table containers with excessive width allocation
- `CountAddedWrapperDivs`: Detects unnecessary wrapper divs around tables that disrupt layout flow

## Scoring Axes Most Affected
- density_control: Direct measure of whether table content is efficiently packed
- spacing_consistency: Arbitrary cell padding breaks the visual rhythm of data grids
- layout_balance: Proper column sizing ensures tables fit their container without overflow or dead space
- component_reuse: Using the project's Table component ensures density presets are applied uniformly
