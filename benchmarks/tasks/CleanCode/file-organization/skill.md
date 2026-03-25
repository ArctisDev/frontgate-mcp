# Skill: Logical File Organization

## Category
CleanCode

## Description
A disorganized file structure makes it difficult for developers to locate code, understand module boundaries, and onboard new team members. This skill teaches the AI agent to group related files together, apply consistent naming conventions for directories, and maintain a clear hierarchy that mirrors the application's feature set. Whether the project follows feature-based, domain-based, or layer-based organization, the agent should respect the existing convention and suggest placements that align with it.

The agent should learn to detect when new files are created at the wrong level of the directory tree, when feature folders lack index files for clean exports, and when shared utilities are scattered rather than centralized. Good file organization is the foundation for scalable codebases and directly impacts developer productivity and code discoverability.

## Expected Outcome
- Place new files within the correct feature or domain directory
- Maintain consistent naming conventions across all directories and files
- Ensure shared utilities and components live in clearly designated shared directories
- Use index files to provide clean public interfaces for feature modules

## Heuristic Triggers
- `CountNewTopLevelComponents`: Flags components created at the root level that belong in feature-specific subdirectories
- `ExtractClassNames`: Reveals class name patterns that suggest components belong together in a shared directory
- `FindHardcodedPaletteClasses`: Palette classes appearing in multiple directories signal that a shared style module is needed

## Scoring Axes Most Affected
- template_likeness: Well-organized files produce templates that are easy to navigate and adapt
- component_reuse: Proper file grouping makes it obvious when existing components can be reused rather than re-created
