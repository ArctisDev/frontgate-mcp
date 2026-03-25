# Skill: Consistent Terminology Across the Interface

## Category
ProductAlignment

## Description
Terminology inconsistency is a subtle but significant source of user confusion. When the same concept is called a "project" in one place, a "workspace" in another, and a "board" in a third, users cannot build reliable mental models of the product. This skill teaches AI agents to maintain strict terminology consistency by identifying the product's canonical vocabulary and ensuring every label, description, and help text adheres to it.

Consistency operates at multiple levels. At the entity level, every data type should have exactly one name used everywhere it appears. At the action level, similar operations should use parallel language. At the system level, UI conventions should be predictable: if "Delete" removes an item in one context, it should not be "Remove" in another unless there is a meaningful distinction. Agents must learn to extract existing terminology from the codebase and product documentation, then apply it uniformly.

Terminology drift often occurs when different parts of an interface are built at different times or by different contributors. AI agents are particularly susceptible to this because they may generate different terms for the same concept in different parts of a generation task.

## Expected Outcome
- Every data entity has exactly one name used consistently across all UI surfaces
- Action verbs are consistent: the same operation uses the same verb everywhere
- Pluralization and capitalization follow a single convention throughout
- Help text, tooltips, and labels use the same terminology as headings and navigation

## Heuristic Triggers
- `CountGenericSaaSLanguage`: Identifies terms that should be replaced with product-specific equivalents
- `ExtractClassNames`: Can reveal naming inconsistencies between CSS classes and visible labels

## Scoring Axes Most Affected
- product_alignment: Consistent terminology builds a coherent product vocabulary
- ux_clarity: Users can rely on consistent terms to predict what they will find
- token_adherence: Design tokens often encode terminology conventions in their naming
