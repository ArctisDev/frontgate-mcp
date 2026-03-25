# Skill: Organized Import Statements

## Category
CleanCode

## Description
Import statements are often the first thing a developer reads when opening a file, and their organization directly affects readability. This skill teaches the AI agent to group imports into clear tiers: external packages first, then internal project modules, then local relative imports. Within each group, alphabetical ordering helps locate specific imports quickly. Blank lines or comments separating groups create a visual hierarchy that mirrors the dependency graph.

The agent should learn to detect mixed import groups, redundant imports, and unused imports that add noise. When a file has a long list of imports, it is often a signal that the component has too many dependencies and might benefit from decomposition. Organized imports also make merge conflicts less likely because additions and removals are predictable.

## Expected Outcome
- Separate imports into external, internal, and local groups with blank line dividers
- Alphabetize imports within each group for quick scanning
- Remove unused imports to reduce noise and potential build warnings
- Ensure default and named imports follow the conventions used in the rest of the codebase

## Heuristic Triggers
- `FindArbitrarySpacing`: Inconsistent spacing between import groups often correlates with broader spacing issues
- `CountNewTopLevelComponents`: New components with sprawling imports signal potential organizational problems
- `CountExistingComponentReuse`: Low reuse scores alongside many imports suggest components are being duplicated instead of imported

## Scoring Axes Most Affected
- template_likeness: Clean imports make generated templates feel professional and ready to use
- visual_noise: Removing unused imports and maintaining order reduces cognitive overhead at the top of every file
