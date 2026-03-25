# Skill: Announcing Errors to Assistive Technology

## Category
Accessibility

## Description
When form validation fails or an operation encounters an error, the message must be communicated to assistive technology users immediately and clearly. This skill teaches the AI agent to use aria-live regions with assertive politeness for critical errors, role="alert" for inline validation messages, and to associate error text with the corresponding form field using aria-describedby. Without these measures, screen reader users may submit a form and receive no indication that something went wrong.

The agent should learn to announce errors at the moment they occur rather than relying on users to discover them by tabbing through fields, to move focus to the first invalid field on form submission, and to clear error announcements when the error is resolved. Error handling is a critical part of the user experience, and accessible error communication ensures that all users can recover from mistakes.

## Expected Outcome
- Use role="alert" or aria-live="assertive" for inline form validation messages
- Associate error text with form fields using aria-describedby so screen readers announce the error on focus
- Move focus to the first invalid field when a form submission fails validation
- Clear aria-live error regions when the user corrects the input to avoid stale announcements

## Heuristic Triggers
- `CountGenericSaaSLanguage`: Generic error messages like "Something went wrong" reduce the helpfulness of error announcements
- `CountAddedWrapperDivs`: Error messages wrapped in non-semantic divs may not be announced unless ARIA attributes are added
- `ExtractClassNames`: Error styling classes applied without accompanying ARIA attributes create visual-only feedback

## Scoring Axes Most Affected
- ux_clarity: Accessible error announcements ensure every user understands what went wrong and how to fix it
- component_reuse: Form components with built-in error announcement patterns are more reliable and easier to reuse
