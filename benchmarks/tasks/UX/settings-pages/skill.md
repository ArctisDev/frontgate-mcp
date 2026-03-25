# Skill: Settings Page Organization

## Category
UX

## Description
Settings pages are where users customize their experience, and poor organization here leads to frustration and support burden. Effective settings design groups related options under clear section headings — account settings, notification preferences, appearance options, and integrations should each live in their own logical section. Labels should describe what the setting does in plain language, not the technical implementation. A toggle labeled "Enable feature flags v2" means nothing to most users; "Receive email notifications for new comments" is immediately clear. Related settings should be visually grouped with consistent spacing and dividers, not scattered across the page.

The AI agent should learn that settings pages need immediate save feedback. If changes save automatically, show a subtle confirmation like a brief "Saved" indicator near the changed setting. If a manual save button is used, disable it until changes are made, then show confirmation after save. Avoid burying frequently changed settings behind multiple navigation levels. Search functionality is valuable when settings exceed a single screen. Avoid mixing informational content (like current plan details) with actionable settings in the same visual hierarchy — users should immediately distinguish between data they can change and data they cannot.

## Expected Outcome
- Settings are grouped under clear, descriptive section headings
- Labels use plain language describing what the setting does, not technical terms
- Save feedback is immediate and visible — either auto-save confirmation or manual save with feedback
- Frequently changed settings are accessible without deep navigation

## Heuristic Triggers
- `FindHardcodedPaletteClasses`: Detects inconsistent styling between settings sections
- `FindArbitrarySpacing`: Flags uneven spacing between settings groups
- `CountGenericSaaSLanguage`: Detects vague setting labels like "Advanced options" or "Miscellaneous"

## Scoring Axes Most Affected
- ux_clarity: Users must understand what each setting controls and how to change it
- layout_balance: Settings groups should have consistent visual weight and spacing
