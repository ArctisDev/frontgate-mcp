# Skill: Respecting Reduced Motion Preferences

## Category
Accessibility

## Description
Some users experience motion sickness, vestibular disorders, or cognitive sensitivities that make animations disorienting or physically uncomfortable. The prefers-reduced-motion media query allows applications to detect when a user has requested minimal motion at the operating system level. This skill teaches the AI agent to respect this preference by reducing or removing non-essential animations, transitions, and parallax effects for users who have opted in to reduced motion.

The agent should learn to wrap animation-heavy styles in a media query that either disables them entirely or provides a simpler alternative such as a fade instead of a slide. Essential motion, such as loading spinners that communicate progress, should be simplified rather than removed. The goal is to preserve the information that motion conveys while eliminating the motion itself for users who cannot tolerate it.

## Expected Outcome
- Wrap non-essential CSS animations and transitions in a prefers-reduced-motion media query
- Provide static alternatives for animated elements such as carousels and scroll-triggered effects
- Simplify essential motion like loading indicators to minimal opacity changes or static states
- Test that the application remains fully functional and informative with reduced motion enabled

## Heuristic Triggers
- `FindArbitrarySpacing`: Arbitrary spacing values may correspond to animation offsets that need reduced-motion alternatives
- `CountNewTopLevelComponents`: Animated hero sections or splash components are common sources of excessive motion
- `FindHardcodedPaletteClasses`: Color transition classes applied via hardcoded utilities may need reduced-motion overrides

## Scoring Axes Most Affected
- visual_noise: Respecting reduced motion preferences directly reduces visual noise for sensitive users
- density_control: Static layouts tend to be denser and more focused, improving readability for all users
