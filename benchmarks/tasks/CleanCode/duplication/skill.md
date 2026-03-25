# Skill: Eliminating Code Duplication

## Category
CleanCode

## Description
The DRY principle states that every piece of knowledge should have a single, unambiguous representation in a system. When patterns are duplicated across components, every change needs to be applied in multiple places, increasing the risk of inconsistency and bugs. This skill teaches the AI agent to recognize repeated code patterns and to extract them into shared utilities, custom hooks, or reusable components. Duplication is not just about identical code blocks; it includes structurally similar patterns that differ only in variable names or minor details.

The agent should learn to distinguish between accidental duplication, which should be refactored, and intentional duplication, where two similar patterns might diverge in the future. Premature abstraction can be as harmful as no abstraction at all. The goal is to consolidate genuine repetition while leaving room for independent evolution where appropriate.

## Expected Outcome
- Detect identical or near-identical code blocks across multiple files
- Extract repeated logic into shared utility functions or custom hooks
- Consolidate duplicate UI patterns into reusable components with flexible props
- Avoid over-abstraction by assessing whether duplicated code is likely to diverge

## Heuristic Triggers
- `CountCardMentions`: Repeated card markup across files indicates a shared Card component is needed
- `CountExistingComponentReuse`: High creation of similar components with low reuse signals duplication
- `FindHardcodedPaletteClasses`: The same palette classes appearing everywhere suggests a shared style utility is missing

## Scoring Axes Most Affected
- component_reuse: DRY refactoring directly increases reuse metrics by creating shared building blocks
- template_likeness: Templates built from shared components feel more cohesive and professionally structured
