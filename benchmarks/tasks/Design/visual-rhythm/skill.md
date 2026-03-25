# Skill: Visual Rhythm Through Spacing Patterns

## Category
Design

## Description
Visual rhythm is the predictable repetition of spacing patterns that guides the user's eye through a page. When a layout alternates between tightly packed sections and generous whitespace without a clear pattern, the user experiences visual turbulence—their eye has to constantly recalibrate. A well-rhythmmed design uses consistent vertical spacing between sections (e.g., every major section separated by py-16 or py-24), consistent horizontal padding within containers (px-4 on mobile, px-8 on desktop), and harmonious gaps within component groups.

This skill teaches the agent to maintain predictable spacing repetition across the page. The agent should identify the project's section spacing rhythm and apply it consistently to new sections. If existing sections use py-16 for vertical padding, the agent should not introduce a new section with py-12 or py-20 without a structural reason (like a sub-section within a larger block). Horizontal rhythm—consistent container widths, aligned gutters, and matching left/right padding—should also be preserved. The agent should think in terms of a baseline grid where all vertical distances are multiples of a base unit.

## Expected Outcome
- Section-level vertical spacing follows a consistent pattern (e.g., all top-level sections use the same py value)
- Horizontal padding and container widths are uniform across the layout
- New sections introduced by the agent match the existing rhythm rather than inventing new spacing values
- The page feels cohesive and easy to scan because spacing relationships are predictable

## Heuristic Triggers
- `FindArbitrarySpacing`: Catches one-off spacing values that break the established rhythm
- `CountGenericSaaSLanguage`: While not directly related, generic content can mask rhythm issues that become visible with real content
- `CountAddedWrapperDivs`: Unnecessary wrappers add invisible spacing that disrupts the intended rhythm

## Scoring Axes Most Affected
- layout_balance: Rhythmic spacing creates visually balanced pages
- spacing_consistency: Rhythm is the macro-level expression of consistent spacing
- template_likeness: Pages with strong visual rhythm feel like polished, production-grade templates
