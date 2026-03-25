# Skill: Progressive Feature Discovery

## Category
ProductAlignment

## Description
Features hidden behind menus, tooltips, or undocumented interactions are features that most users will never find. Conversely, presenting all features at once overwhelms users and creates visual clutter. This skill teaches AI agents to implement progressive disclosure that introduces features contextually when users are most likely to need them, building a path from basic to advanced usage without requiring external documentation.

Progressive discovery operates on the principle that features should surface when the user's context makes them relevant. A formatting toolbar appears when text is selected. Import options appear when viewing an empty list. Advanced filters appear after basic search returns too many results. Agents must learn to identify feature relationships and design contextual introduction points rather than dumping all capabilities into a single menu or settings page.

The balance is between discoverability and simplicity. Every feature needs a home, but not every feature needs to be visible at all times. Primary features should be immediately visible. Secondary features should be one interaction away. Advanced features can be two interactions away, but should be hinted at through contextual cues.

## Expected Outcome
- Primary features are immediately visible on the relevant screen
- Secondary features are accessible through one clear interaction (button, menu, expand)
- Advanced features are hinted at contextually when the user's situation suggests they might be useful
- Feature introduction avoids tutorial overlays in favor of inline guidance

## Heuristic Triggers
- `CountNewTopLevelComponents`: Can reveal feature bloat where too many top-level sections compete for attention
- `FindOversizedSidebarClasses`: May indicate sidebar navigation that has grown to accommodate too many features without progressive disclosure

## Scoring Axes Most Affected
- ux_clarity: Users can find features when they need them without being overwhelmed
- product_alignment: Feature organization reflects the product's usage patterns
- visual_noise: Progressive disclosure reduces visible complexity
