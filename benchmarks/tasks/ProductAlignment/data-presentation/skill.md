# Skill: Contextual Data Presentation

## Category
ProductAlignment

## Description
How data is formatted and presented significantly affects whether users can quickly extract meaning from an interface. Numbers without units, dates without context, percentages without baselines, and statuses without visual differentiation all create cognitive overhead. This skill teaches AI agents to present data in formats that match user expectations and domain conventions, making information immediately interpretable without requiring mental conversion.

Data presentation decisions should reflect user needs. Financial data should show appropriate decimal precision and currency symbols. Timestamps should use relative or absolute formats depending on whether recency or specificity matters. Status indicators should combine color, icon, and text to ensure accessibility and clarity. Large numbers should be abbreviated with appropriate notation. Agents must learn to identify data types from context and apply formatting conventions that match the product's domain.

Beyond formatting, data should be presented in relationships that provide context. A metric without its trend is less useful. A date without its relative position (today, yesterday, 3 days ago) requires more processing. Data should be grouped and ordered to support the user's decision-making process.

## Expected Outcome
- Numbers include appropriate units, precision, and formatting for the domain
- Dates and times use formats appropriate to context (relative for recency, absolute for scheduling)
- Status values combine visual indicators (color, icon) with text labels for accessibility
- Metrics include context like trends, comparisons, or thresholds where meaningful

## Heuristic Triggers
- `CountRawPrimitiveAdditions`: Detects raw HTML elements used for data display instead of domain-appropriate components
- `CountGenericSaaSLanguage`: Can indicate generic data labels that don't reflect domain specifics

## Scoring Axes Most Affected
- product_alignment: Data formatting reflects domain expertise and user expectations
- ux_clarity: Properly formatted data is faster to interpret
- visual_noise: Over-formatted or inconsistently formatted data adds noise
