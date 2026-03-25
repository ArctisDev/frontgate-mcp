# Skill: Success Feedback Patterns

## Category
UX

## Description
Success feedback confirms that a user's action completed as intended and reduces uncertainty. Without confirmation, users may repeat actions, reload pages, or doubt whether their input was received. The right feedback pattern depends on the action's significance and context. Toast notifications work well for transient confirmations like saving a setting or copying a link. Inline confirmations suit form submissions where the user stays on the same page. State changes — like a button shifting from "Add to Cart" to "Added" — provide immediate visual feedback without interrupting flow. For significant actions like deleting data or completing a purchase, a more prominent confirmation with detail is appropriate.

The AI agent should learn to match feedback intensity to action significance. Minor actions like toggling a setting need subtle inline feedback. Major actions like publishing content or completing a workflow warrant more visible confirmations. Success toasts should auto-dismiss after a few seconds for non-critical actions but persist for significant ones. Avoid showing success messages for every micro-interaction — over-notification creates noise and trains users to ignore feedback. Never show success feedback when the action has not actually completed on the server.

## Expected Outcome
- Success feedback matches the significance of the action performed
- Toast notifications auto-dismiss for minor actions and persist for significant ones
- Button state changes provide immediate inline confirmation without page disruption
- No premature success messages shown before server confirmation

## Heuristic Triggers
- `CountGenericSaaSLanguage`: Detects generic success text like "Success!" or "Done" without context
- `CountCardMentions`: Flags unnecessary card wrappers around simple success toasts
- `CountNewTopLevelComponents`: Detects ad-hoc success components instead of reusing shared toast systems

## Scoring Axes Most Affected
- ux_clarity: Users need explicit confirmation that their action succeeded
- component_reuse: Shared toast and notification systems ensure consistent feedback patterns
