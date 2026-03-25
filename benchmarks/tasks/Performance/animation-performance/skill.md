# Skill: GPU-Accelerated Animations

## Category
Performance

## Description
Animations that trigger layout or paint operations run on the main thread and can cause janky, stuttering motion. This skill teaches AI agents to create smooth, 60fps animations by restricting animation targets to properties that can be handled entirely by the GPU compositor thread. The two properties that can be animated without triggering layout or paint are `transform` (translate, scale, rotate) and `opacity`. All other animated properties—including `width`, `height`, `margin`, `padding`, `top`, `left`, `background-color`, and `color`—trigger expensive operations.

Agents must learn to translate common animation patterns into compositor-friendly equivalents. Instead of animating `height` to expand a section, use `transform: scaleY()` or the modern `grid-template-rows` trick. Instead of animating `top`/`left` for movement, use `transform: translate()`. Instead of toggling `display: none`, use `opacity` combined with `pointer-events`. The `will-change` hint can inform the browser to promote elements to their own compositor layer, but overuse wastes memory.

CSS transitions and animations are preferred over JavaScript-based animations because the browser can optimize them, potentially running them off the main thread. The Web Animations API provides JavaScript control while maintaining the same optimization potential.

## Expected Outcome
- All animations target `transform` and `opacity` exclusively
- No animations trigger layout (avoiding `width`, `height`, `margin`, `padding`, `top`, `left`)
- `will-change` is used sparingly and only on elements that will animate
- CSS transitions/animations are preferred over JavaScript animation loops

## Heuristic Triggers
- `FindArbitrarySpacing`: May reveal margin/padding-based animations that trigger layout
- `FindHardcodedPaletteClasses`: Could indicate color animations that trigger paint operations

## Scoring Axes Most Affected
- visual_noise: Smooth animations reduce visual disruption; janky animations increase it
- density_control: Animated elements with `will-change` consume compositor memory
