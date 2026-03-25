# Skill: Consistent Icon System Usage

## Category
Design

## Description
Icons should come from a single, consistent icon system. Agents that mix icon libraries—pulling a Lucide icon here, a Heroicon there, and maybe a custom SVG for a third element—create visual inconsistency. Different icon libraries use different stroke weights, corner radii, grid sizes, and optical adjustments. Mixing them in the same interface is the typographic equivalent of using five different fonts on one page. The result looks patched together rather than designed.

This skill teaches the agent to identify and exclusively use the project's designated icon library. The agent should inspect the project's dependencies (package.json), existing component imports, and icon usage patterns to determine which library is canonical—whether it's Lucide, Heroicons, Radix Icons, or a custom set. Once identified, all new icon usage should come from that library. The agent should also maintain consistent icon sizing (using the project's standard size props or classes), consistent placement relative to text labels, and consistent coloring through semantic tokens rather than hardcoded values. If a needed icon doesn't exist in the project's library, the agent should choose the closest available match rather than introducing a new library.

## Expected Outcome
- All icons in new code come from the project's designated icon library—no mixed sources
- Icon sizes follow the project's standard sizing conventions (e.g., h-4 w-4 for inline, h-5 w-5 for buttons)
- Icons use semantic color tokens (text-muted-foreground, text-primary) rather than hardcoded colors
- The agent does not introduce a new icon library dependency without explicit instruction

## Heuristic Triggers
- `CountRawPrimitiveAdditions`: Flags when new SVG markup or icon imports from previously unused libraries are added
- `ExtractClassNames`: Surfaces icon-related classes and imports to verify they come from the canonical library
- `CountExistingComponentReuse`: Evaluates whether the agent reuses existing icon-wrapping components rather than creating new ones

## Scoring Axes Most Affected
- component_reuse: Using existing icon components from the project's library is a core form of reuse
- template_likeness: Consistent iconography is a hallmark of polished, production-grade interfaces
- visual_noise: Mixed icon styles create subtle but pervasive visual inconsistency
