# Skill: Badge Systems

## Category
Component

## Description
Badges and status indicators are compact visual elements that communicate state, counts, and categories at a glance. When badge implementations vary across the product, users encounter different color semantics (does red mean "error" or "new"?), inconsistent sizing, and mismatched border radii. A shared Badge component provides a single source of truth for color mapping (success, warning, error, info, neutral), sizing presets, and optional features like dot indicators, removable close buttons, and count truncation (e.g., "99+"). Ad-hoc badge implementations often skip accessibility considerations such as sufficient color contrast and screen reader text for status meaning.

This skill teaches the agent to use the project's Badge component for all status, count, and label indicators rather than writing custom styled spans. The agent should learn to compose badges with appropriate `variant`, `size`, and `color` props, and to use the component's count formatting (truncation at 99+, abbreviation for thousands) rather than hardcoding display logic. When badges appear alongside other components (cards, table cells, nav items), the agent should position them using the Badge component's built-in placement API or composition patterns, ensuring that badge alignment and spacing remain consistent across contexts.

## Expected Outcome
- Status indicators use the project's Badge component rather than custom colored spans
- Badge color variants map to the design system's semantic palette (not arbitrary hex values)
- Count badges use the component's built-in truncation logic for large numbers
- Badge placement alongside other components follows the design system's alignment conventions

## Heuristic Triggers
- `FindHardcodedPaletteClasses`: Detects custom color classes on badge-like elements that bypass the design system's semantic palette
- `CountRawPrimitiveAdditions`: Flags custom styled `<span>` or `<div>` elements used as badges
- `CountExistingComponentReuse`: Measures whether the Badge component is used for status and count indicators
- `ExtractClassNames`: Identifies ad-hoc badge classes (e.g., `.status-chip`, `.pill-tag`) that duplicate the Badge component's functionality

## Scoring Axes Most Affected
- component_reuse: Using the shared Badge component prevents parallel status indicator implementations
- token_adherence: Ad-hoc badges bypass semantic color tokens, leading to inconsistent color meaning
- visual_noise: Multiple badge styles fragment the visual language and reduce status legibility
- template_likeness: Products with consistent badge systems feel more polished and professionally designed
