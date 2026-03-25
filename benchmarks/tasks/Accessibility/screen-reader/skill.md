# Skill: Screen Reader Compatibility

## Category
Accessibility

## Description
Screen readers convert visual interfaces into spoken or braille output, and they rely on semantic structure, ARIA roles, and live regions to do so accurately. This skill teaches the AI agent to use semantic HTML elements that convey meaning natively, to apply ARIA roles only when native semantics are insufficient, and to implement aria-live regions for dynamic content updates such as notifications, loading states, and form validation results. Without these measures, screen reader users receive incomplete or misleading information.

The agent should learn to associate labels with form controls, to use landmarks like main, nav, and aside to help users jump between page sections, and to avoid using div or span for interactive elements that should be buttons or links. Screen reader compatibility is not an afterthought but a fundamental design constraint that shapes how markup is written from the start. Every piece of dynamic content needs a strategy for how it will be announced.

## Expected Outcome
- Use semantic HTML elements like button, nav, main, and form over generic div containers
- Add ARIA roles only when native HTML semantics are genuinely insufficient
- Implement aria-live regions for toast notifications, progress indicators, and dynamic content updates
- Ensure form controls have associated labels and that error messages are programmatically linked

## Heuristic Triggers
- `CountCardMentions`: Card grids that lack ARIA roles may not announce their list structure to screen readers
- `HasClearHeadingStructure`: A clear heading hierarchy is essential for screen reader navigation between sections
- `CountGenericSaaSLanguage`: Vague ARIA labels drawn from generic copy reduce the information screen readers can convey

## Scoring Axes Most Affected
- template_likeness: Screen reader-friendly templates are more robust and universally usable
- component_reuse: Components designed with screen readers in mind have well-defined accessibility contracts
