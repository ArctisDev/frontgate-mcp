# Skill: Dropdown Patterns

## Category
Component

## Description
Dropdowns and select menus are among the most common interactive elements, yet they are frequently implemented inconsistently. Raw `<select>` elements are styled differently across browsers, making them visually unpredictable. Custom dropdown implementations often lack keyboard navigation, proper focus management, or correct ARIA attributes (`listbox`, `option`, `aria-expanded`). A shared Dropdown or Select component solves these issues by providing a single implementation for trigger styling, menu positioning, search/filter capabilities, and multi-select behavior that works consistently across browsers and assistive technologies.

This skill teaches the agent to use the project's Dropdown or Select component for all menu-based selections rather than writing raw `<select>` elements or building custom dropdown logic. The agent should learn to compose dropdown triggers and menu items using the shared component's API, applying consistent `placeholder`, `disabled`, `error`, and `searchable` props. When the dropdown needs custom rendering (e.g., icons in options, grouped items), the agent should use the component's slot or render-prop API rather than building a new dropdown from scratch. This ensures that every selection interaction in the product behaves identically.

## Expected Outcome
- Raw `<select>` elements are flagged when a Dropdown component exists in the codebase
- Dropdown triggers use consistent sizing, border, and icon placement from the design system
- Keyboard navigation (arrow keys, Enter, Escape) is handled by the shared component
- Searchable and multi-select variants use the component's built-in API

## Heuristic Triggers
- `CountRawPrimitiveAdditions`: Detects raw `<select>` elements added instead of the project's Dropdown component
- `CountExistingComponentReuse`: Measures whether the Dropdown component is composed into new forms and views
- `ExtractClassNames`: Identifies custom dropdown classes that bypass the design system's trigger and menu tokens
- `FindArbitrarySpacing`: Detects ad-hoc margin/padding on dropdown containers that breaks form alignment

## Scoring Axes Most Affected
- component_reuse: Using the shared Dropdown prevents duplicate selection UI implementations
- token_adherence: Custom dropdowns bypass tokens for trigger borders, menu shadows, and option hover colors
- ux_clarity: Consistent dropdown behavior lets users select options confidently without relearning interactions
- visual_noise: Multiple dropdown styles fragment the form interface and reduce perceived polish
