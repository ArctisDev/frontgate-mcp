# Skill: Card Usage Discipline

## Category
Layout

## Description
This skill teaches the AI agent to reserve card components for situations where a distinct visual grouping of related content is genuinely needed, rather than reflexively wrapping every section in a card. Cards create visual boundaries through borders, shadows, and rounded corners, and overusing them turns a clean layout into a grid of competing boxes. The agent should ask: does this content need to be visually separated from its surroundings, or would a simple section with spacing suffice? Cards are appropriate for items in a list or grid where each item is an independent entity (e.g., product cards, user profiles, dashboard widgets), but they add unnecessary visual weight to inline content like a settings form or a help article.

When the agent does use cards, it should ensure the card styling is consistent: uniform border-radius, shadow depth, and padding. The agent should reuse a single card component or utility class rather than creating bespoke card styles for each use case. If the existing codebase has a `Card` component, the agent should use it. If a new card variant is needed, it should extend the existing card rather than creating a parallel implementation.

The agent should also be cautious about nesting cards inside cards, which creates confusing visual hierarchies. If a card contains a list of items, those items typically do not need their own card treatment unless they are being presented as independently interactive elements. Flat content within a card is almost always preferable to nested card structures.

## Expected Outcome
- Cards are used only where visual grouping is needed, not as a default wrapper
- A single card component or utility class is reused across the application
- Card nesting is avoided or limited to one level of depth
- Card styling (radius, shadow, padding) is consistent and token-driven

## Heuristic Triggers
- `CountCardMentions`: Detects excessive card usage in contexts where simpler section layouts would suffice
- `CountNewTopLevelComponents`: Flags new components that introduce card wrappers without clear grouping needs
- `FindHardcodedPaletteClasses`: Identifies card elements with hardcoded shadow or border values instead of tokens

## Scoring Axes Most Affected
- visual_noise: Reducing unnecessary cards lowers the number of competing visual boundaries
- density_control: Flat layouts without excess cards use screen space more efficiently
- component_reuse: Reusing a standard card component ensures consistency and reduces code duplication
