# Skill: Error Handling and Recovery

## Category
UX

## Description
Error handling is about communicating failures clearly and giving users a path to recovery. Errors occur in many contexts: form validation failures, network timeouts, permission denials, and server errors. Each context demands a different communication style. Inline form errors should appear directly adjacent to the field that failed, with specific guidance on how to fix the issue. Toast notifications work well for action-level errors like a failed save or delete. Full-page error states suit catastrophic failures like a 500 error or failed page load. The worst pattern is the silent error — when something fails but the user sees no indication and later discovers data was lost.

The AI agent should learn that effective error messages are specific, actionable, and human-readable. Never show raw error codes or stack traces to end users. Every error state should include a recovery action: a retry button, a link to documentation, or at minimum guidance on what to try next. Error toasts should not auto-dismiss for critical failures. For forms, errors should summarize what went wrong at the top of the form while also highlighting individual fields. Avoid generic messages like "Something went wrong" without offering next steps.

## Expected Outcome
- Form errors appear inline next to the relevant field with specific fix instructions
- Action errors use toast notifications with a retry option when applicable
- No raw error codes, stack traces, or technical jargon shown to end users
- Critical errors do not auto-dismiss and require explicit user acknowledgment

## Heuristic Triggers
- `CountGenericSaaSLanguage`: Detects vague error messages like "Something went wrong"
- `FindArbitrarySpacing`: Flags layout shifts caused by error message appearance
- `CountAddedWrapperDivs`: Detects over-nested error containers that complicate the DOM

## Scoring Axes Most Affected
- ux_clarity: Error messages must communicate what happened and how to recover
- visual_noise: Poorly styled errors create visual clutter and increase anxiety
