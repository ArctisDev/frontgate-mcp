# Skill: Shadow System Consistency

## Category
Design

## Description
Shadows communicate elevation and interactivity in a UI. A well-designed shadow system defines a small set of shadow levels—typically shadow-sm for subtle depth on cards, shadow-md for dropdowns and popovers, shadow-lg for modals, and shadow-xl for top-level overlays. When agents introduce arbitrary shadow values like `shadow-[0_3px_10px_rgba(0,0,0,0.12)]` or mix dozens of different blur and spread combinations, the elevation hierarchy becomes ambiguous and the interface loses its sense of spatial layering.

This skill teaches the agent to use the project's shadow tokens or Tailwind's built-in shadow scale rather than composing custom box-shadow declarations. The agent should map use cases to the correct elevation tier: form inputs get no shadow or shadow-sm, cards get shadow-sm or shadow-md, floating elements get shadow-lg. If the project has defined custom shadow tokens in its design system, the agent should discover and use those. Arbitrary shadow values should be flagged and replaced with the nearest standard tier.

## Expected Outcome
- The agent uses standard shadow utilities (shadow-sm, shadow-md, shadow-lg) from the project's design system
- No arbitrary shadow values (shadow-[...]) appear in new components
- Shadow usage maps to a clear elevation hierarchy: interactive surfaces have appropriate depth cues
- The agent distinguishes between hover/focus states and static shadows using transition utilities

## Heuristic Triggers
- `FindArbitrarySpacing`: Catches bracket-notation shadow utilities that deviate from the design system's shadow scale
- `ExtractClassNames`: Surfaces shadow-related classes for audit against allowed elevation tokens
- `CountRawPrimitiveAdditions`: Flags introduction of raw shadow values outside the project's token set

## Scoring Axes Most Affected
- token_adherence: Measures whether shadow usage follows the project's elevation token system
- visual_noise: Inconsistent shadows create visual clutter and break the depth hierarchy
- template_likeness: A cohesive shadow system makes components feel polished and production-ready
