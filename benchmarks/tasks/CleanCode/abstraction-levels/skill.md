# Skill: Consistent Abstraction Levels

## Category
CleanCode

## Description
Every function and component should operate at a single level of abstraction. When a component mixes high-level layout concerns with low-level DOM manipulation or data fetching details, it becomes difficult to read and hard to test. This skill teaches the AI agent to keep components focused on one concern and to extract lower-level details into helper functions or child components. A well-abstracted component reads like a narrative: high-level intent at the top, details tucked away in named helpers below.

The agent should learn to detect when a component's body contains both JSX layout structure and imperative logic, when data transformation happens inline alongside rendering, or when event handlers contain more than a few lines of business logic. Each of these signals a mixing of abstraction levels that should be addressed by extraction.

## Expected Outcome
- Keep rendering logic separate from data transformation and business logic
- Extract detailed or imperative code into named helper functions
- Use child components to isolate low-level UI concerns from high-level layout
- Ensure each function reads at one consistent level of detail

## Heuristic Triggers
- `CountNewTopLevelComponents`: New components created during refactoring signal abstraction extraction is happening
- `FindArbitrarySpacing**: Inconsistent spacing within a file often indicates mixed concerns fighting for visual space
- `CountRawPrimitiveAdditions`: Inline primitive values scattered through JSX suggest data details that should be abstracted

## Scoring Axes Most Affected
- ux_clarity: Consistent abstraction makes the UI's intent readable to both developers and the agent
- density_control: Properly abstracted code naturally maintains a comfortable density of information
