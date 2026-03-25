# Skill: Descriptive Naming Conventions

## Category
CleanCode

## Description
Names are the primary form of documentation in any codebase. Poorly named variables, components, and functions force every reader to mentally decode what the code actually does. This skill teaches the AI agent to use meaningful, self-describing names that communicate intent without requiring comments or additional context. Variable names should describe what data they hold, component names should describe what they render or represent, and function names should describe what action they perform.

The agent should avoid abbreviations except for widely accepted conventions, use consistent casing patterns, and ensure that names do not mislead. For example, a boolean variable should read like a question or a statement of state, a handler function should begin with a verb, and a component should be named after the thing it displays. Good naming reduces the need for inline documentation and makes code reviews faster and more effective.

## Expected Outcome
- Use descriptive variable names that convey purpose without abbreviation
- Name components after the UI element or feature they represent
- Prefix handler functions with action verbs like handle, on, or toggle
- Maintain consistent casing: PascalCase for components, camelCase for variables and functions

## Heuristic Triggers
- `CountRawPrimitiveAdditions`: Raw primitives with generic names like data, item, or val indicate weak naming
- `CountGenericSaaSLanguage`: Generic language in strings often mirrors generic naming in code
- `ExtractClassNames`: Class names that are abbreviated or opaque signal naming convention violations

## Scoring Axes Most Affected
- ux_clarity: Clear naming translates to clear UI copy and consistent user-facing terminology
- token_adherence: Well-named design tokens and classes reinforce the project's token vocabulary
