# Skill: Proper ARIA Labeling

## Category
Accessibility

## Description
ARIA attributes provide essential context to assistive technologies, enabling screen reader users to understand and interact with dynamic interfaces. This skill teaches the AI agent to apply aria-label when visible text is insufficient, aria-labelledby when another element should serve as the label, and aria-describedby when supplementary information enhances understanding. Without these attributes, interactive elements become opaque to users relying on assistive technology.

The agent should learn to audit every interactive element: buttons with only icons need aria-label, form inputs need associated labels via htmlFor or aria-labelledby, and complex widgets need descriptive text that clarifies their purpose. ARIA labeling is not a replacement for semantic HTML but a complement that fills gaps where native elements cannot convey enough meaning on their own.

## Expected Outcome
- Add aria-label to icon-only buttons and controls that lack visible text
- Use aria-labelledby to associate visible headings or text with their corresponding interactive regions
- Apply aria-describedby to connect helper text, validation messages, or instructions to form fields
- Avoid redundant ARIA on elements that already have accessible names via native semantics

## Heuristic Triggers
- `CountCardMentions`: Card components often contain icon buttons that need explicit aria-label values
- `CountGenericSaaSLanguage`: Generic placeholder text often signals missing or vague ARIA descriptions
- `ExtractClassNames`: Utility classes applied to interactive elements may obscure their semantic role, requiring ARIA

## Scoring Axes Most Affected
- ux_clarity: Proper ARIA labeling ensures every interactive element communicates its purpose to all users
- component_reuse: Well-labeled components are more reusable because their accessibility contracts are explicit
