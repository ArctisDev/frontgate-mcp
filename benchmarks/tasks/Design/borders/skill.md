# Skill: Border Usage and Visual Separation

## Category
Design

## Description
Borders are one of the most overused visual separation techniques in UI development. Agents frequently wrap every card, panel, and section in a solid border, creating a grid-like cage effect that fragments the layout and increases visual noise. Modern design systems prefer spacing, background color differentiation, and subtle shadows to delineate regions, reserving borders for specific contexts like form inputs, table cells, and dividers where a literal boundary line is semantically appropriate.

This skill teaches the agent to evaluate whether a border is truly necessary before adding one. When visual separation is needed, the agent should first consider increasing the gap between elements, applying a slightly different background tone (e.g., bg-muted vs bg-card), or adding a subtle shadow. If a border is the right tool, the agent should use the project's border token (border-border, border-input) rather than a hardcoded color. The agent should also avoid double-border situations where both a parent and child element have borders, creating a visually heavy frame.

## Expected Outcome
- The agent defaults to spacing, background contrast, and shadows for visual separation before reaching for borders
- When borders are used, they reference semantic border tokens (border-border, border-input) rather than hardcoded colors
- No unnecessary double-border patterns appear (e.g., a bordered card containing bordered sections)
- The agent reduces border usage in new code compared to naive implementations

## Heuristic Triggers
- `CountCardMentions`: High card counts combined with border classes suggest over-bordering
- `ExtractClassNames`: Surfaces border-related classes to audit for overuse and hardcoded colors
- `CountAddedWrapperDivs`: Extra wrapper divs often carry unnecessary borders that add visual weight

## Scoring Axes Most Affected
- visual_noise: Overuse of borders is a primary source of UI clutter
- density_control: Borders consume visual space and can make layouts feel cramped
- layout_balance: Clean separation through spacing and background creates more balanced layouts than border-heavy designs
