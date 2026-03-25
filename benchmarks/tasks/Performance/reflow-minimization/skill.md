# Skill: Minimizing Layout Reflows

## Category
Performance

## Description
Layout reflows (also called layout thrashing) occur when the browser must recalculate element positions and dimensions, which is one of the most expensive operations in the rendering pipeline. This skill teaches AI agents to minimize reflows by batching DOM reads and writes, avoiding layout-triggering CSS properties in animations, and understanding which properties trigger layout, paint, or composite operations. Properties like `width`, `height`, `top`, `left`, `margin`, and `padding` trigger layout, while `transform` and `opacity` only trigger compositing, which is GPU-accelerated and much cheaper.

The read-write-read pattern is a common cause of forced synchronous layout. Reading a DOM property (like `offsetHeight`) after writing to it (like changing `style.height`) forces the browser to synchronously compute layout to return an accurate value. Agents must learn to batch all reads together, then all writes together, to avoid this pattern. In React, this often manifests when measuring elements in `useEffect` after state changes that affect layout.

CSS containment (`contain: layout style paint size`) tells the browser that an element's internal changes do not affect outside elements, allowing it to skip layout recalculation for ancestors. This is particularly useful for components with frequently changing content like lists, carousels, or animated sections.

## Expected Outcome
- Animations use `transform` and `opacity` instead of layout-triggering properties
- DOM reads and writes are batched to avoid forced synchronous layout
- CSS `contain` property is used to isolate frequently changing components
- Layout-triggering properties are avoided in frequently executed code paths

## Heuristic Triggers
- `FindArbitrarySpacing`: Arbitrary spacing values may indicate layout-based positioning that triggers reflow
- `FindHardcodedPaletteClasses`: May reveal inline styles that bypass CSS optimization

## Scoring Axes Most Affected
- spacing_consistency: Layout-triggering spacing changes can cause visual jitter
- visual_noise: Reflows cause visual disruptions that increase perceived noise
