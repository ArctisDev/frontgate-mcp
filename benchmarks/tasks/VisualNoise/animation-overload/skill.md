# Skill: Reducing Unnecessary Animations

## Category
VisualNoise

## Description
Animations serve three legitimate purposes: providing feedback on user actions, smoothing transitions between states, and directing attention to important changes. Agents frequently add fade-ins, slide-ups, bounces, and pulse effects to make the UI feel dynamic and alive. But animation overload creates a jittery, distracting experience where movement competes with content for the user's focus.

This skill teaches that every animation must earn its place. Buttons should have a brief transition on hover for feedback. Modals should animate in to preserve spatial context. Content that loads dynamically can fade in to avoid jarring pop-ins. Beyond these cases, do not animate section entrances, decorative elements, or static content. A calm interface with selective motion is more professional and usable than one where everything slides, fades, and bounces.

## Expected Outcome
- Animations limited to hover feedback, state transitions, and modal/drawer entrances
- No animation on static content sections or decorative elements
- Animation durations under 300ms for micro-interactions
- Consistent easing curves across all animated elements

## Heuristic Triggers
- `CountNewTopLevelComponents`: Can reveal animated wrapper components added purely for motion effects
- `CountAddedWrapperDivs`: Animation wrapper divs are a common pattern for applying entrance animations

## Scoring Axes Most Affected
- visual_noise: Fewer animations means a calmer, less distracting interface
- ux_clarity: Purposeful motion improves comprehension while random motion degrades it
