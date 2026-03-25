# Skill: Dialog Patterns

## Category
Component

## Description
Dialogs and modals are high-attention UI surfaces that interrupt the user's workflow. When implemented inconsistently, they degrade trust: some dialogs might have backdrop overlays while others do not, close button placement varies, and action button alignment shifts between left, right, and center. A shared Dialog or Modal component solves this by providing a single implementation for overlay behavior, focus trapping, keyboard dismissal, and action button layout. Bypassing this component introduces accessibility gaps and visual inconsistency that users notice immediately.

This skill teaches the agent to use the project's existing Dialog or Modal primitive for every overlay interaction, including confirmations, forms, and informational popovers. The agent should learn to compose dialog content using slots or children rather than building custom overlay logic. When a new modal pattern is needed (e.g., a multi-step wizard), the agent should extend the existing Dialog's API or nest composable panels inside it, not create a parallel modal system. This ensures consistent backdrop behavior, animation curves, and focus management across the entire application.

## Expected Outcome
- Confirmation dialogs reuse the project's Dialog component with standard action button placement
- Form dialogs use the shared Modal with consistent header, body, and footer sections
- Backdrop overlay, focus trap, and keyboard dismiss behavior come from the shared primitive
- Dialog sizing and animation follow project conventions rather than custom implementations

## Heuristic Triggers
- `CountNewTopLevelComponents`: Flags when new modal or overlay components are created instead of reusing the Dialog
- `CountRawPrimitiveAdditions`: Detects raw overlay/backdrop markup added outside the shared Dialog system
- `CountExistingComponentReuse`: Measures whether Dialog/Modal components are composed into new views
- `CountAddedWrapperDivs`: Identifies unnecessary wrapper structures around dialog content that break the component's slot API

## Scoring Axes Most Affected
- component_reuse: Reusing the Dialog primitive ensures consistent modal behavior app-wide
- ux_clarity: Uniform dialog patterns let users predict interaction outcomes without relearning
- visual_noise: Inconsistent modal styles fragment the visual hierarchy and reduce perceived polish
- token_adherence: Custom dialog styles bypass tokens for overlay color, border radius, and shadow depth
