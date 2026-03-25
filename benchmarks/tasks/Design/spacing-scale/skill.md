# Skill: Spacing Scale Consistency

## Category
Design

## Description
Spacing is the backbone of visual structure. Tailwind provides a well-defined spacing scale (p-1 through p-96, gap-1 through gap-96) that maps to a 4px base unit. When agents introduce arbitrary spacing values like `gap-[17px]`, `p-[23px]`, or `mb-[11px]`, they break the mathematical relationships that keep layouts aligned and predictable. Arbitrary values make it nearly impossible for other developers—or the agent itself—to maintain spatial consistency across components.

This skill teaches the agent to always select from the existing spacing scale. The agent should first check whether the project uses Tailwind's default scale or a customized token set (e.g., spacing-xs, spacing-sm defined in tailwind.config). When a layout needs spacing that falls between two scale steps, the agent should choose the closer step rather than invent a fractional value. Consistent spacing also applies to gaps in flex and grid containers, margin/padding on nested elements, and vertical rhythm between sections.

## Expected Outcome
- All spacing utilities use standard scale values (gap-4, p-6, mb-8) rather than arbitrary values (gap-[17px], p-[23px])
- The agent demonstrates awareness of the project's spacing tokens and applies them consistently
- Nested components maintain proportional spacing relationships (e.g., inner padding is a multiple of the base unit)
- No `FindArbitrarySpacing` violations appear in newly authored code

## Heuristic Triggers
- `FindArbitrarySpacing`: Detects bracket-notation spacing utilities like `p-[13px]`, `gap-[7px]`, `m-[22px]` that deviate from the scale
- `CountRawPrimitiveAdditions`: Flags when new raw spacing values are introduced instead of using existing tokens
- `ExtractClassNames`: Surfaces all spacing-related classes to audit against the allowed scale

## Scoring Axes Most Affected
- spacing_consistency: Primary axis—measures whether spacing adheres to a uniform scale across the UI
- density_control: Consistent spacing prevents components from feeling too cramped or too sparse
- layout_balance: Uniform spacing creates symmetrical, balanced layouts
