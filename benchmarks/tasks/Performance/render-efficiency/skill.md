# Skill: Efficient React Rendering

## Category
Performance

## Description
React's rendering model can cause unnecessary re-renders that degrade performance, especially in large component trees. This skill teaches AI agents to write rendering-efficient React code by understanding when components re-render and how to prevent wasteful renders. Key techniques include `React.memo` for component-level memoization, `useMemo` for expensive computations, `useCallback` for stable function references, and stable key props for list items. Each technique has specific use cases and overusing them can be as harmful as underusing them.

The most common source of unnecessary re-renders is creating new object or function references on every render. When a parent renders, all children re-render unless they are memoized and their props are referentially stable. `useCallback` keeps function references stable, `useMemo` keeps computed values stable, and `React.memo` prevents re-renders when props have not changed. Agents must learn to identify the component tree structure and apply memoization strategically at component boundaries where prop changes are infrequent.

Stable keys in lists are critical for React's reconciliation algorithm. Using array indices as keys forces React to re-render all list items when the list changes, while stable unique IDs allow React to update only the changed items. This is especially important for large lists where reconciliation overhead is significant.

## Expected Outcome
- Components with expensive renders use `React.memo` with appropriate comparison functions
- Expensive computations are wrapped in `useMemo` to avoid recalculation on every render
- Callbacks passed to child components are wrapped in `useCallback` for referential stability
- List items use stable unique keys, not array indices

## Heuristic Triggers
- `CountRawPrimitiveAdditions`: Raw elements in lists may lack proper key props
- `CountNewTopLevelComponents`: Excessive component nesting can increase re-render surface area

## Scoring Axes Most Affected
- layout_balance: Efficient rendering maintains smooth layout updates
- density_control: Dense UIs with many components benefit most from render optimization
