# Skill: Semantic Color Tokens

## Category
Design

## Description
In a well-maintained design system, colors are abstracted behind semantic CSS variables (e.g., `--primary`, `--secondary`, `--surface`, `--danger`) rather than referenced directly through Tailwind's palette utilities like `bg-blue-500`, `text-red-600`, or `border-gray-300`. When an AI agent hardcodes palette classes, it couples the UI to a specific color value, making future theme changes, dark mode support, and brand updates fragile and error-prone. Semantic tokens decouple intent from implementation.

This skill teaches the agent to inspect the project's existing CSS variables and Tailwind config to identify which semantic tokens are available, then apply those tokens consistently. The agent should map contextual intent (primary action, muted text, surface background) to the correct token rather than picking colors that "look right." When no appropriate token exists, the agent should flag the gap instead of inventing a new hardcoded value.

## Expected Outcome
- The agent replaces hardcoded palette classes (e.g., `bg-blue-500`, `text-gray-700`) with semantic token utilities (e.g., `bg-primary`, `text-muted-foreground`)
- New components reference existing CSS custom properties from the project's design system rather than Tailwind's default palette
- The agent identifies missing semantic tokens and suggests additions instead of falling back to raw palette values
- Theme consistency is maintained across light and dark modes without manual intervention

## Heuristic Triggers
- `FindHardcodedPaletteClasses`: Detects usage of Tailwind's raw palette utilities (bg-blue-500, text-red-600, etc.) instead of semantic token classes
- `ExtractClassNames`: Surfaces all color-related classes so the agent can audit whether semantic tokens are being used

## Scoring Axes Most Affected
- token_adherence: Directly measures whether the agent uses project tokens over hardcoded palette values
- visual_noise: Hardcoded colors can create inconsistency that increases perceived clutter across the UI
