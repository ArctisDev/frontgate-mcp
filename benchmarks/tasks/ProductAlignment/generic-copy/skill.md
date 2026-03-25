# Skill: Replacing Generic SaaS Copy

## Category
ProductAlignment

## Description
Generic SaaS language is one of the most pervasive problems in modern web interfaces. Words like "insights", "analytics", "overview", "dashboard", and "create new" appear everywhere without conveying any specific meaning about what the product actually does or what the user will find. This skill teaches AI agents to recognize when copy defaults to these placeholder patterns and replace them with language that reflects the specific product domain and user mental model.

The core principle is that every label, heading, and call-to-action should communicate something meaningful about the product's value proposition. Instead of "View Insights," a project management tool might say "Check Sprint Progress." Instead of "Analytics," a fitness app might show "Your Weekly Training Load." Agents must learn to read the surrounding context, understand the product's purpose, and generate copy that a user would actually use in conversation about the product.

This extends beyond individual labels to the overall voice and tone of the interface. Generic copy creates a sense that the product could be anything, which means it feels like nothing in particular. Product-specific copy builds trust and communicates competence by showing that the builders understand the domain deeply.

## Expected Outcome
- Headings and section titles use domain-specific language rather than generic SaaS terms
- Call-to-action buttons describe the specific outcome of the action, not abstract verbs
- Navigation labels reflect the product's actual mental model, not default CRUD terminology
- Empty states, tooltips, and helper text reinforce the product's unique vocabulary

## Heuristic Triggers
- `CountGenericSaaSLanguage`: Detects occurrences of overused SaaS terms like "insights", "analytics", "overview" that lack product context
- `FindArbitrarySpacing`: Can reveal layout hacks used to compensate for poorly-fitting generic labels

## Scoring Axes Most Affected
- product_alignment: Directly measures whether copy reflects the specific product rather than generic SaaS patterns
- ux_clarity: Users understand what actions do when labels are specific rather than abstract
