# Skill: Input Variants

## Category
Component

## Description
Form inputs are fundamental to user interaction, and inconsistency across input fields degrades perceived quality immediately. When developers write raw `<input>`, `<select>`, or `<textarea>` elements without using the project's established input components, the resulting forms have mismatched heights, border styles, focus rings, and label placements. A well-designed input system provides consistent sizing, validation states, helper text slots, and accessibility attributes out of the box, eliminating the need to reimplement these details per form.

This skill teaches the agent to recognize when raw HTML input primitives are used instead of the project's `Input`, `Select`, `Textarea`, or `Checkbox` components. The agent should learn to scan the repository for these shared components and compose forms by assembling them, applying consistent `size`, `variant`, and `state` props. When custom input behavior is needed, the agent should wrap the existing primitive rather than replacing it entirely. This approach ensures that focus management, error styling, and label associations remain uniform across every form in the product.

## Expected Outcome
- Raw `<input>`, `<select>`, and `<textarea>` tags are flagged when project-specific input components exist
- Input components are used with consistent size and variant props across forms
- Error and helper text patterns use the component's built-in slots rather than custom markup
- Label associations and aria attributes follow the shared input component's conventions

## Heuristic Triggers
- `CountRawPrimitiveAdditions`: Detects raw `<input>`, `<select>`, `<textarea>` elements added without using project components
- `CountExistingComponentReuse`: Measures whether Input, Select, and Textarea components are composed into new forms
- `ExtractClassNames`: Identifies custom input classes that diverge from the design system's input tokens
- `FindArbitrarySpacing`: Detects ad-hoc margin/padding on inputs that should use the component's built-in spacing

## Scoring Axes Most Affected
- component_reuse: Core measure of whether shared input primitives are leveraged
- token_adherence: Ad-hoc input styles bypass design tokens for borders, backgrounds, and focus indicators
- spacing_consistency: Raw inputs often carry arbitrary spacing that breaks form rhythm
- ux_clarity: Consistent inputs reduce cognitive overhead and improve form completion rates
