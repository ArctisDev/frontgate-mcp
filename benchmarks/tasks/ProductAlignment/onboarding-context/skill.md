# Skill: Contextual Onboarding

## Category
ProductAlignment

## Description
Traditional onboarding relies on sequential tutorial overlays that users dismiss without reading and forget immediately. Contextual onboarding instead integrates guidance into the product experience, teaching users as they encounter new features and workflows. This skill teaches AI agents to design onboarding that feels like part of the product rather than an interruption, using empty states, inline hints, and progressive feature introduction to guide users through their first experience.

Contextual onboarding recognizes that users learn best when they have a reason to learn. An empty dashboard that explains what will appear there once data exists is more useful than a popup tour explaining every feature. A form field with an example value teaches format expectations better than a tooltip that appears on hover. Agents must learn to identify the moments where guidance is most valuable and design contextual hints that enhance rather than obstruct the experience.

The key metric is whether users successfully complete their first meaningful task. Onboarding should reduce the time and friction to that first success, not exhaustively document every feature before the user has any context for understanding them.

## Expected Outcome
- Empty states serve as onboarding moments, explaining what will appear and suggesting first actions
- Inline hints guide users through their first workflow without requiring a separate tutorial mode
- Progressive disclosure introduces complexity only after users have mastered basics
- Guidance is dismissible and does not reappear once the user has demonstrated competence

## Heuristic Triggers
- `CountAddedWrapperDivs`: Excessive wrapping can indicate tutorial overlay structures that should be replaced with inline guidance
- `FindArbitrarySpacing`: Layout hacks may appear when overlay-based onboarding disrupts normal page flow

## Scoring Axes Most Affected
- ux_clarity: Contextual guidance is clearer than abstract tutorials
- product_alignment: Onboarding should use the product's own vocabulary and workflows
- visual_noise: Inline guidance adds less noise than overlay-based tutorials
