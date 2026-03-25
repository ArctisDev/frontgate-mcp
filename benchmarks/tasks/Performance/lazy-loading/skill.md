# Skill: Lazy Loading and On-Demand Content

## Category
Performance

## Description
Loading all page content upfront wastes bandwidth and blocks rendering for content the user may never see. This skill teaches AI agents to implement lazy loading strategies that defer non-critical content until it is needed. The primary techniques include React.lazy with Suspense for component-level code splitting, dynamic imports for heavy dependencies, and Intersection Observer for loading content as it enters the viewport. Each technique serves a different scenario and agents must learn to choose the appropriate one.

Component-level lazy loading with `React.lazy()` and `<Suspense>` is ideal for heavy UI sections like charts, editors, or modals that are not immediately visible. Dynamic `import()` statements allow deferring entire feature modules until a user action triggers them. Intersection Observer-based lazy loading works well for images, lists, and content sections below the fold. Agents must also understand the pitfalls: lazy loading above-the-fold content creates visible delays, and aggressive lazy loading can cause layout shifts when content appears.

The goal is not to lazy load everything but to identify the critical rendering path and defer everything else. The initial bundle should contain only what is needed for the first meaningful paint, with additional content loaded progressively as the user interacts or scrolls.

## Expected Outcome
- Heavy components (charts, editors, modals) are lazy loaded with React.lazy and Suspense
- Below-the-fold content uses Intersection Observer for deferred loading
- Dynamic imports split feature modules that are not needed for initial render
- Suspense fallbacks provide appropriate loading states without layout shifts

## Heuristic Triggers
- `CountNewTopLevelComponents`: Can reveal eagerly loaded sections that should be deferred
- `FindArbitrarySpacing`: Layout shifts from lazy-loaded content may appear as spacing inconsistencies

## Scoring Axes Most Affected
- density_control: Lazy loading affects how much content is initially rendered
- visual_noise: Proper lazy loading reduces initial visual complexity
