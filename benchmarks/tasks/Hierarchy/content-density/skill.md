# Skill: Optimal Information Density

## Category
Hierarchy

## Description
Information density refers to how much content is visible per unit of screen space. Agents tend to err in one of two directions: too sparse, with excessive whitespace that forces endless scrolling and hides context, or too cramped, with small fonts and minimal padding that overwhelms the user. Neither extreme serves the user well — a sparse dashboard wastes space while a cramped table is unreadable.

This skill teaches finding the right density for the content type and user context. Data-heavy views like tables and dashboards can be denser because users expect to scan rows and compare values. Marketing pages and onboarding flows should be sparser to guide the user through a narrative. Within a single page, maintain consistent density — do not mix a spacious hero section with a cramped data table. Use padding, font size, and line height tokens to enforce a consistent density profile across the interface.

## Expected Outcome
- Content density appropriate to the information type and user context
- Consistent density across all sections of a single page
- Data views denser than narrative or marketing content
- No sections that are conspicuously sparse or cramped relative to neighbors

## Heuristic Triggers
- `FindArbitrarySpacing`: Arbitrary spacing values often indicate density inconsistency
- `CountRawPrimitiveAdditions`: Raw size and spacing primitives bypass density tokens

## Scoring Axes Most Affected
- density_control: Proper density keeps content readable without wasting space
- template_likeness: Consistent density matches the feel of polished production interfaces
