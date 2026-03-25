# Skill: Onboarding Flow Design

## Category
UX

## Description
Onboarding introduces users to a product's core value and guides them toward their first meaningful interaction. Effective onboarding uses progressive disclosure — showing only what is needed at each step rather than overwhelming users with every feature upfront. This means breaking the experience into logical stages: account setup, first action, first result, and then deeper exploration. Contextual tooltips and inline hints that appear when a user encounters a feature for the first time are more effective than upfront tutorial modals that users click through without reading. The goal is to reduce time-to-value, not to document every feature.

The AI agent should learn that onboarding should never feel like a wall of information. Avoid multi-step wizards with more than three or four steps for initial setup. Skip patterns and the ability to return later are essential — users who already know the product should not be forced through a guided tour. Progress indicators help set expectations for multi-step flows. Empty states within the product should reinforce onboarding by reminding users of their next action. Avoid forcing users to watch videos, read tooltips, or complete checklists before they can access the product.

## Expected Outcome
- Onboarding flows have no more than 3-4 steps for initial setup
- Users can skip or defer onboarding without losing access to the product
- Contextual help appears inline at the moment of need, not in advance
- Progress indicators are present for multi-step setup flows

## Heuristic Triggers
- `CountNewTopLevelComponents`: Detects excessive onboarding-specific components that should reuse existing UI
- `CountGenericSaaSLanguage`: Flags vague instructional text like "Get started by exploring"
- `CountCardMentions`: Detects unnecessary card wrappers that add visual weight to simple step content

## Scoring Axes Most Affected
- density_control: Onboarding should be lean and avoid information overload
- template_likeness: Onboarding should feel integrated, not bolted on as a separate experience
