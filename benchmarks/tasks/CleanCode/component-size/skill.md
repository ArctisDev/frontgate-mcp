# Skill: Keeping Components at a Manageable Size

## Category
CleanCode

## Description
Large components are one of the most common sources of technical debt in front-end codebases. When a single component exceeds two hundred lines, it typically conflates multiple responsibilities: state management, event handling, layout composition, and data transformation all live in one file. This skill teaches the AI agent to recognize when a component is growing beyond a comfortable size and to split it into smaller, focused sub-components that each own a single concern.

The agent should learn to extract reusable UI fragments, isolate complex logic into custom hooks or utility functions, and keep render methods readable. A well-split component tree is not only easier to test and debug but also enables developers to reuse pieces independently. The goal is never strict line-count enforcement but rather a signal that the component is doing too much and should be decomposed.

## Expected Outcome
- Identify components that exceed approximately two hundred lines of code
- Suggest extracting distinct UI sections into dedicated sub-components
- Separate business logic from presentation by introducing custom hooks or helpers
- Maintain clear component boundaries so each file has a single responsibility

## Heuristic Triggers
- `CountNewTopLevelComponents`: Monitors whether new top-level components are being introduced to absorb growing responsibilities
- `CountCardMentions`: Detects repeated card-like patterns that could be consolidated into a single reusable component
- `FindArbitrarySpacing`: Spacing inconsistencies often appear when components grow too large and style rules diverge

## Scoring Axes Most Affected
- density_control: Smaller components naturally enforce tighter density control and prevent layout sprawl
- visual_noise: Decomposed components reduce visual clutter by isolating styling concerns per section
