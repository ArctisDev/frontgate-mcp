# Skill: Container Widths

## Category
Layout

## Description
This skill teaches the AI agent to apply consistent maximum width constraints to content containers so that text remains readable and layouts do not stretch awkwardly on ultrawide monitors. Without `max-width` constraints, paragraphs can extend to unreadable line lengths (80+ characters per line), and side-by-side content can spread so far apart that the eye cannot track the visual relationship between elements. The agent should learn to use `max-width` values like 640px for prose content, 1024px for standard page content, and 1280px or wider only for dashboards and data-dense interfaces.

Container widths should be defined as design tokens or CSS custom properties and referenced consistently throughout the application. The agent should avoid hardcoding different `max-width` values in each component, which leads to inconsistent content boundaries across pages. A well-designed container system uses a small set of named sizes (e.g., `content-sm`, `content-md`, `content-lg`) that map to specific pixel or `rem` values. When a new section needs a content width, the agent should pick from the existing set rather than inventing a new one.

The agent should also ensure that containers use `margin: 0 auto` or equivalent centering so that content is balanced within the viewport. Full-width sections that break out of the container should do so intentionally and re-establish their own internal container for child content. The goal is a predictable reading experience where the user always knows where to expect content boundaries.

## Expected Outcome
- Content containers use `max-width` constraints appropriate to their content type
- Container width values are defined as tokens or custom properties, not hardcoded
- A small set of named container sizes (2-4) is reused across the application
- Containers are horizontally centered and do not drift to one side of the viewport

## Heuristic Triggers
- `FindArbitrarySpacing`: Detects hardcoded `max-width` pixel values that bypass design tokens
- `CountNewTopLevelComponents`: Flags new page-level components that introduce non-standard container widths
- `ExtractClassNames`: Identifies ad-hoc class names suggesting one-off width configurations

## Scoring Axes Most Affected
- layout_balance: Proper container widths create balanced, readable layouts across screen sizes
- typographic_hierarchy: Content-width constraints ensure line lengths remain readable at every viewport
- template_likeness: Consistent container widths match established web conventions and user expectations
