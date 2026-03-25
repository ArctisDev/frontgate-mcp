# Skill: Clear Task Communication

## Category
ProductAlignment

## Description
Users should never wonder what they can do on a screen or what a screen is for. Task clarity means every page communicates its purpose through a combination of headings, descriptions, visual hierarchy, and action affordances. This skill teaches AI agents to ensure that the purpose of each screen is self-evident, reducing the cognitive load required to understand what is possible and what is expected.

Task clarity starts with the page heading and supporting description. A heading should describe the object or goal, not the technical operation. "Manage Team Members" is clearer than "User Administration." Supporting text should briefly explain the scope: "Invite, remove, and configure access for your team." Below the heading, the layout should make primary tasks visually obvious through prominent action buttons and clear content organization.

Ambiguity arises when screens try to serve multiple purposes equally. A page that is simultaneously a list, a form, and a report confuses users about what they should do. Agents should learn to ensure each screen has a clear primary purpose, even if secondary tasks are also available.

## Expected Outcome
- Every screen has a clear heading that describes its purpose in user-facing terms
- Supporting descriptions explain what the user can accomplish on the screen
- The visual hierarchy makes primary tasks immediately apparent
- Complex screens break information into clearly labeled sections

## Heuristic Triggers
- `HasClearHeadingStructure`: Verifies that pages have logical, accessible heading hierarchies
- `CountGenericSaaSLanguage`: Generic headings often fail to communicate specific task purpose

## Scoring Axes Most Affected
- ux_clarity: Directly measures whether users can understand what to do on each screen
- product_alignment: Task descriptions should reflect the product's actual capabilities
- typographic_hierarchy: Heading structure is essential for task communication
