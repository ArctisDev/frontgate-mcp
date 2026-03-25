# Skill: Button Consistency

## Category
Component

## Description
Button components are among the most reused primitives in any design system. When developers bypass the project's existing `Button` component and write raw `<button>` elements with ad-hoc styling, it creates visual inconsistency across the application. Different padding, font sizes, border radii, and hover states accumulate, making the UI feel disjointed and unpolished. A shared Button primitive enforces a single source of truth for sizing variants, color schemes, loading states, and icon placements.

This skill teaches the agent to detect when raw `<button>` tags or custom-styled button elements are introduced instead of reusing the established Button component. The agent should learn to scan the repository for existing button primitives and prefer composition over reinvention. When a new button-like interaction is needed, the agent should extend the existing component's API (e.g., adding a new variant) rather than creating a parallel implementation. This reduces code duplication, improves maintainability, and ensures consistent interaction patterns throughout the product.

## Expected Outcome
- Raw `<button>` elements are flagged when a Button component already exists in the codebase
- Custom button styles are identified as candidates for Button component reuse
- Button variant props (size, intent, appearance) are used consistently across new code
- Loading and disabled states use the shared Button's built-in API rather than ad-hoc implementations

## Heuristic Triggers
- `CountRawPrimitiveAdditions`: Detects raw `<button>` tags added to templates instead of the project's Button component
- `CountExistingComponentReuse`: Measures whether existing button primitives are being composed into new views
- `ExtractClassNames`: Identifies inline or custom button classes that diverge from the design system's button tokens
- `CountNewTopLevelComponents`: Flags when entirely new button-like components are created instead of reusing existing ones

## Scoring Axes Most Affected
- component_reuse: Directly measures whether shared button primitives are leveraged or bypassed
- token_adherence: Ad-hoc button styles often skip design tokens for colors, spacing, and radii
- visual_noise: Inconsistent button appearances across the app increase cognitive load for users
- template_likeness: Polished UIs reuse button primitives; raw buttons signal prototype-quality work
