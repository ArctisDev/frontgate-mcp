# Skill: Form Validation UX

## Category
UX

## Description
Form validation balances two goals: preventing invalid submissions and not frustrating users with premature error messages. Real-time validation — checking fields as the user types or on blur — helps users correct issues incrementally, but showing errors before the user has finished typing a field is hostile. The best practice is to validate on blur for the first pass (after the user has interacted with and left a field), then validate on change after the first error has appeared for that field. Submit-time validation catches anything missed and should present a summary of all errors at the top of the form, with each error linking to the relevant field.

The AI agent should learn that validation messaging must be specific and positioned near the problem. "Email is required" near the email field is clear; "Please fix the errors below" alone at the top is not. Disable the submit button while an async validation is in flight to prevent duplicate submissions, but re-enable it promptly. For multi-step forms, validate each step before allowing progression. Avoid showing red borders and error text simultaneously on every field the moment the user lands on the form — let users engage first. Error summaries at the top of forms should list each issue with anchor links to the corresponding field.

## Expected Outcome
- Validation errors appear on blur first, then on change after initial error
- Submit button is disabled during async validation to prevent duplicate submissions
- Error summary at the top of the form lists all issues with links to fields
- Error messages are specific and positioned adjacent to the relevant field

## Heuristic Triggers
- `FindArbitrarySpacing`: Detects layout shifts when validation errors appear or disappear
- `CountAddedWrapperDivs`: Flags excessive nesting around form fields to accommodate error messages
- `FindHardcodedPaletteClasses`: Detects hardcoded red error colors instead of semantic tokens

## Scoring Axes Most Affected
- ux_clarity: Users must understand exactly what is wrong and how to fix it
- token_adherence: Error states should use design tokens for colors and spacing
