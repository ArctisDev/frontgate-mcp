# Skill: Avoiding Prop Drilling

## Category
CleanCode

## Description
Prop drilling occurs when data is passed through multiple layers of components that do not actually consume it, creating tightly coupled component trees and making refactoring painful. This skill teaches the AI agent to recognize when props are threaded through intermediate components unnecessarily and to apply alternatives such as React Context, component composition via children/render props, or restructuring the component hierarchy to colocate state with the consumers that need it.

By avoiding prop drilling, the resulting codebase becomes easier to maintain and reason about. Each component receives only the data it directly uses, reducing cognitive overhead for developers and minimizing the surface area for bugs when prop signatures change. The agent should learn to identify drilling patterns early and suggest the most appropriate abstraction for the situation at hand.

## Expected Outcome
- Detect props that pass through components without being consumed by them
- Suggest Context providers or composition patterns when prop chains exceed two levels
- Reduce the total number of prop declarations across the component tree
- Keep component interfaces lean and focused on their direct responsibilities

## Heuristic Triggers
- `CountNewTopLevelComponents`: Detects creation of wrapper components whose sole purpose is forwarding props
- `ExtractClassNames`: Surfaces tightly coupled class hierarchies that mirror prop chains
- `CountAddedWrapperDivs`: Flags unnecessary structural wrappers introduced to pass props through

## Scoring Axes Most Affected
- component_reuse: Components become more reusable when they do not depend on specific parent prop shapes
- template_likeness: Clean component boundaries produce templates that feel intentional rather than tangled
