# Skill: Managing Side Effects Properly

## Category
CleanCode

## Description
Side effects such as data fetching, subscriptions, timers, and DOM manipulation must be managed carefully in React to avoid memory leaks, stale closures, and inconsistent renders. This skill teaches the AI agent to use useEffect with proper dependency arrays, to always return cleanup functions when subscriptions or timers are involved, and to avoid performing side effects during the render phase. The agent should also learn to move side effects into event handlers or custom hooks when they are triggered by user actions rather than by mount or state changes.

Improper side-effect management is a leading cause of subtle bugs in React applications. Effects without cleanup cause growing memory usage over time, effects with missing dependencies produce stale data, and render-time effects cause unnecessary re-execution. The agent should treat every useEffect as a potential source of bugs and verify its correctness.

## Expected Outcome
- Always provide a dependency array to useEffect that accurately reflects what the effect reads
- Return a cleanup function from effects that set up subscriptions, timers, or listeners
- Avoid performing side effects directly during render; place them inside useEffect or event handlers
- Group related effects and extract complex effect logic into custom hooks

## Heuristic Triggers
- `CountNewTopLevelComponents`: Extracting effect-heavy logic into dedicated components or hooks improves clarity
- `FindArbitrarySpacing`: Unusual spacing around useEffect blocks often signals hastily added effects
- `CountGenericSaaSLanguage`: Vague loading or error messages may indicate effects that lack proper state management

## Scoring Axes Most Affected
- ux_clarity: Properly managed effects ensure loading states, error states, and data freshness are communicated clearly
- template_likeness: Templates with well-managed side effects are more robust and easier to adapt
