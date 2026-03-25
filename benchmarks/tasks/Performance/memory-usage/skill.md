# Skill: Preventing Memory Leaks in React

## Category
Performance

## Description
Memory leaks in React applications accumulate over time, causing increasing memory usage, degraded performance, and eventually browser tab crashes. This skill teaches AI agents to properly clean up side effects in React components using `useEffect` cleanup functions. Every subscription, event listener, timer, observer, and asynchronous operation that is started during a component's lifecycle must be stopped or cancelled when the component unmounts or when the effect re-runs.

The most common leak patterns include: `addEventListener` without `removeEventListener`, `setInterval` without `clearInterval`, `setTimeout` without `clearTimeout`, `new IntersectionObserver()` without `observer.disconnect()`, and fetch requests without `AbortController.abort()`. Agents must learn to always return a cleanup function from `useEffect` when setting up any resource that needs disposal. The cleanup function runs both on unmount and before the effect re-executes, making it the correct place for all teardown logic.

In addition to explicit leaks, agents should be aware of closure-based memory issues where event handlers or callbacks capture large objects in their closure scope, preventing garbage collection even after the listener is removed.

## Expected Outcome
- Every `useEffect` that creates subscriptions, listeners, or timers returns a cleanup function
- Event listeners added with `addEventListener` have corresponding `removeEventListener` in cleanup
- Intervals and timeouts are cleared in the effect cleanup function
- Fetch requests use `AbortController` to cancel on unmount or dependency change

## Heuristic Triggers
- `CountRawPrimitiveAdditions`: Raw DOM manipulation may create listeners that bypass React cleanup
- `CountNewTopLevelComponents`: Many top-level components with effects increase leak surface area

## Scoring Axes Most Affected
- density_control: Dense UIs with many components and effects are at higher risk for leaks
- layout_balance: Memory pressure can cause visible performance degradation in layouts
