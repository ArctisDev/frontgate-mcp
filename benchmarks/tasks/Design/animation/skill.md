# Skill: Purposeful Animation and Transitions

## Category
Design

## Description
Animations and transitions should serve a functional purpose: communicating state changes, providing feedback on interactions, and guiding attention during navigation. Agents that add animations for decorative purposes—spinning logos, bouncing elements, gratuitous fade-ins on page load—create interfaces that feel sluggish and gimmicky. Every millisecond of animation has a cost: it delays the user from seeing or interacting with content. Purposeful animation is fast (100-200ms for micro-interactions, 200-300ms for transitions), subtle, and directly tied to a user action or system event.

This skill teaches the agent to use Tailwind's transition utilities (transition-colors, transition-opacity, transition-transform) for hover, focus, and active states on interactive elements. The agent should apply consistent timing (duration-150 or duration-200) and easing (ease-in-out) across similar interactions so the UI feels cohesive. Loading states should use skeleton screens or subtle spinners, not elaborate animations. Page transitions should be minimal or nonexistent in data-heavy applications. The agent should never add animation to static content (headings, paragraphs, images) that the user needs to read immediately. If the project has defined animation tokens or presets, the agent should use those exclusively.

## Expected Outcome
- Interactive elements (buttons, links, inputs) have consistent hover/focus transitions using transition-colors or transition-opacity
- Animation durations are fast and purposeful (150-300ms range), not sluggish or decorative
- No gratuitous entrance animations on static content that blocks the user from reading immediately
- Loading states use the project's established patterns (skeletons, spinners) rather than custom animations

## Heuristic Triggers
- `CountRawPrimitiveAdditions`: Flags raw animation or transition CSS properties introduced outside the project's token system
- `ExtractClassNames`: Surfaces animation and transition classes to audit for consistency and purpose
- `CountNewTopLevelComponents`: Excessive new components may indicate the agent is adding animation wrapper elements instead of applying transitions to existing elements

## Scoring Axes Most Affected
- ux_clarity: Purposeful animation clarifies state changes; decorative animation obscures content
- visual_noise: Unnecessary animations are a form of visual noise that distracts from the interface's purpose
- density_control: Heavy animations can make a UI feel sluggish and reduce perceived responsiveness
