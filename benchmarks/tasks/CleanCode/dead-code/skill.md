# Skill: Removing Dead Code

## Category
CleanCode

## Description
Dead code includes unused components, functions, variables, imports, and CSS rules that accumulate over time as features evolve. Every piece of dead code is a liability: it inflates bundle size, confuses developers who wonder if it is still needed, and can hide real bugs when commented-out logic is accidentally re-enabled. This skill teaches the AI agent to identify code that is no longer referenced anywhere and to remove it cleanly without affecting active functionality.

The agent should learn to trace references across files, recognize commented-out blocks that should be deleted rather than preserved, and flag entire files that appear to have no imports. Removing dead code is a form of housekeeping that keeps the codebase lean and trustworthy. It also simplifies refactoring because there are fewer paths through the code to consider.

## Expected Outcome
- Identify and remove functions, components, and variables that are never referenced
- Delete commented-out code blocks instead of leaving them as historical artifacts
- Remove CSS classes and utility styles that are no longer applied to any element
- Ensure removal is safe by confirming no dynamic references depend on the dead code

## Heuristic Triggers
- `CountExistingComponentReuse`: Components with zero reuse instances across the project are likely dead code
- `CountCardMentions`: Unused card pattern imports signal dead template code
- `ExtractClassNames`: Class names defined but never applied in any component markup indicate dead styles

## Scoring Axes Most Affected
- density_control: Dead code inflates perceived complexity; removing it tightens the density of meaningful code
- visual_noise: Fewer unused rules and files means less noise when navigating the project
