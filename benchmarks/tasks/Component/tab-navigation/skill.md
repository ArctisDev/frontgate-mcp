# Skill: Tab Navigation

## Category
Component

## Description
Tabs are a primary mechanism for organizing related content into switchable panels. When tab implementations vary across the product, users encounter different active-state indicators, inconsistent underline or pill styles, and misaligned panel content. A shared Tab component ensures that active and inactive states use the same color tokens, transition timing, and focus styles. It also handles keyboard navigation (arrow keys to switch tabs, Enter to activate) and ARIA roles (`tablist`, `tab`, `tabpanel`) in a single place, preventing accessibility regressions that occur when tabs are built from raw divs.

This skill teaches the agent to use the project's Tab component for all content switching scenarios. The agent should learn to compose tab triggers and panels using the shared component's API, applying consistent `selected`, `disabled`, and `icon` props. When building views with multiple tab-like sections (settings pages, detail panels, dashboards), the agent should reuse the same Tab primitive rather than writing custom CSS for active underline transitions or click handlers. This guarantees that tab navigation feels identical whether the user is in a settings panel, a data dashboard, or a content editor.

## Expected Outcome
- Tab implementations use the project's shared Tab component rather than custom markup
- Active and inactive tab states follow the design system's color and typography tokens
- Keyboard navigation and ARIA roles are handled by the shared component, not reimplemented
- Tab panel content spacing is consistent across different tab instances

## Heuristic Triggers
- `CountRawPrimitiveAdditions`: Detects custom tab markup built from divs or spans instead of the Tab component
- `ExtractClassNames`: Identifies ad-hoc tab classes (e.g., `.tab-active`, `.tab-custom`) that bypass the design system
- `FindArbitrarySpacing`: Detects custom padding on tab triggers that breaks alignment with the panel content
- `CountExistingComponentReuse`: Measures whether the Tab component is used when tabs are needed

## Scoring Axes Most Affected
- component_reuse: Using the shared Tab component prevents duplicate tab implementations
- token_adherence: Custom tab styles override tokens for active indicators, hover states, and focus rings
- ux_clarity: Consistent tab behavior lets users switch content confidently across the product
- visual_noise: Multiple tab styles increase cognitive load and reduce perceived system coherence
