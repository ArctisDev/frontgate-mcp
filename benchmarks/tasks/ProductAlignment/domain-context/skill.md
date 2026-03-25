# Skill: Reflecting Product Domain in UI

## Category
ProductAlignment

## Description
A product's domain should permeate every aspect of its interface, from the data it displays to the metaphors it uses for navigation and organization. A healthcare application should feel clinical and precise. A creative tool should feel expressive and flexible. A logistics platform should emphasize efficiency and status clarity. This skill teaches AI agents to recognize domain signals and ensure the UI communicates the product's purpose through visual and informational design choices.

Domain context goes beyond copy into structural decisions. What data gets prominence on a detail page? What relationships between entities are surfaced? How are statuses represented? A CRM should foreground contact relationships and communication history. An inventory system should emphasize stock levels and reorder thresholds. Agents must learn to identify the domain from product descriptions and codebase context, then make UI decisions that align with domain conventions.

The danger of ignoring domain context is producing interfaces that feel generic and interchangeable. When every product uses the same card layout, the same sidebar structure, and the same dashboard pattern regardless of domain, users lose the contextual cues that help them navigate and understand the product quickly.

## Expected Outcome
- UI structure prioritizes the data and relationships most important to the product's domain
- Status indicators and data representations use domain-appropriate metaphors
- Information hierarchy reflects what domain experts would expect to see first
- Component choices reinforce the product's category (e.g., tables for data-heavy domains, cards for visual domains)

## Heuristic Triggers
- `CountGenericSaaSLanguage`: Flags copy that ignores domain-specific vocabulary
- `ExtractClassNames`: Can reveal generic layout patterns that don't adapt to domain needs
- `CountExistingComponentReuse`: Detects when generic components are used without domain customization

## Scoring Axes Most Affected
- product_alignment: Measures whether the interface feels like it belongs to its specific domain
- template_likeness: Domain-specific UIs should not look like generic templates
- ux_clarity: Domain-appropriate patterns help users orient faster
