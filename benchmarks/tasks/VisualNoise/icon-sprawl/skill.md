# Skill: Limiting Icon Density

## Category
VisualNoise

## Description
Icons are powerful affordances for actions and status, but when every heading, list item, sidebar link, and stat card gets an icon, the interface becomes an icon gallery. Agents frequently prepend icons to labels because icon libraries make it easy and because icons add visual interest. The problem is that excessive iconography creates a pattern-matching burden — users cannot quickly distinguish meaningful action icons from decorative ones.

This skill teaches using icons purposefully: for interactive buttons where the icon clarifies the action, for status indicators where the icon communicates state, and for navigation landmarks where the icon aids recognition. Do not add icons to every text label, section heading, or informational paragraph. When in doubt, omit the icon — a clear label alone is often more effective.

## Expected Outcome
- Icons present only on actions, status indicators, and key navigation items
- Section headings and informational text free of decorative icons
- Icon count proportional to the number of interactive elements
- Visual field less cluttered by small graphic elements

## Heuristic Triggers
- `CountNewTopLevelComponents`: Can reveal icon-heavy component additions
- `CountRawPrimitiveAdditions`: Flags raw icon elements added outside the component system

## Scoring Axes Most Affected
- visual_noise: Fewer icons reduce the number of small visual elements competing for attention
- density_control: Removing decorative icons frees space for content
