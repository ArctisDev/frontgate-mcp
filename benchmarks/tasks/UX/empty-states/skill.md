# Skill: Empty State Design

## Category
UX

## Description
Empty states are the screens users see when there is no data to display — an empty inbox, a blank dashboard, a project with no tasks yet. These moments are critical touchpoints where users decide whether to engage further or abandon the product. A well-designed empty state does more than fill space; it educates the user about what belongs in this area, why it is currently empty, and provides a clear, low-friction action to populate it. Generic illustrations with vague copy like "Nothing here yet" fail to guide users toward meaningful interaction.

The AI agent should learn that effective empty states combine contextual explanation with a strong call to action. The messaging should reference the specific entity the user is viewing (e.g., "This project has no tasks" rather than "No items found"). When appropriate, include a primary action button that initiates the most common workflow for that screen. Avoid adding decorative illustrations that consume space without conveying useful information. Empty states should never feel like dead ends — they are invitations to begin.

## Expected Outcome
- Empty views contain a heading that references the specific entity (project, folder, workspace)
- A descriptive line explains what type of content belongs in this view
- A primary action button is present to create or add the first item
- No decorative-only illustrations or SVGs that add no informational value

## Heuristic Triggers
- `CountGenericSaaSLanguage`: Detects generic placeholder text like "Nothing here yet" or "No items found"
- `CountCardMentions`: Flags unnecessary card wrappers around simple empty state content
- `CountAddedWrapperDivs`: Detects over-nested DOM structure in what should be a simple centered layout

## Scoring Axes Most Affected
- ux_clarity: Users need to understand why the view is empty and what to do next
- density_control: Empty states should be concise and not padded with unnecessary visual elements
