# Skill: Typographic Scale and Hierarchy

## Category
Design

## Description
A consistent typographic scale ensures that every heading, body text, label, and caption in the UI communicates its relative importance at a glance. Agents that introduce ad-hoc font sizes (e.g., `text-[22px]`, `text-[15px]`) or mix too many weight classes (font-light, font-semibold, font-black in the same context) break the visual rhythm and make the interface feel amateurish. A mature design system restricts typography to a small, predictable set of sizes and weights defined as tokens.

This skill teaches the agent to use the project's existing heading hierarchy (h1 through h6 mapped to Tailwind's text-2xl, text-xl, etc.) and body defaults. The agent should verify heading levels are sequential—no jumping from h1 to h4—and that semantic HTML tags match the visual weight. When styling inline text, the agent should prefer utility classes that already exist in the project rather than inventing new font-size/weight combinations.

## Expected Outcome
- The agent uses a limited, project-defined set of font sizes (e.g., text-sm, text-base, text-lg, text-xl, text-2xl) instead of arbitrary values
- Heading elements follow a sequential hierarchy without skipped levels
- Font weight usage is consistent: headings use one weight tier, body text another, captions another
- No ad-hoc `text-[Npx]` or `font-[N]` arbitrary value classes appear in new code

## Heuristic Triggers
- `FindArbitrarySpacing`: Catches arbitrary pixel values used in font-size utilities like `text-[17px]`
- `HasClearHeadingStructure`: Validates that heading elements (h1-h6) are used in a logical, sequential order
- `ExtractClassNames`: Surfaces typography-related classes for audit against the allowed scale

## Scoring Axes Most Affected
- typographic_hierarchy: Directly evaluates whether font sizes and weights create a clear visual hierarchy
- template_likeness: Pages with inconsistent typography feel less like a polished, production-ready template
