# Skill: Using Semantic HTML Elements

## Category
Accessibility

## Description
Semantic HTML provides built-in accessibility features that generic elements like div and span cannot offer. A button element is focusable, clickable by keyboard, and announced as a button by screen readers. A nav element is recognized as a navigation landmark. A section with a heading creates a document outline that assistive technologies use to help users jump between content areas. This skill teaches the AI agent to prefer semantic elements over generic ones whenever the content or role matches an available HTML element.

The agent should learn to recognize when a div with an onClick handler should be a button, when a collection of links should be wrapped in a nav, when content sections should use header, main, footer, and aside, and when lists of items should use ul or ol instead of stacked divs. Semantic HTML is the foundation of accessibility; ARIA attributes and JavaScript enhancements should only be layered on top when semantics alone are not enough.

## Expected Outcome
- Replace div-based interactive elements with native button, a, or input elements
- Use nav, main, aside, header, and footer to define page landmarks
- Wrap content sections in section or article elements with appropriate headings
- Use ordered and unordered lists for list content rather than repeated div structures

## Heuristic Triggers
- `CountAddedWrapperDivs`: High wrapper div counts suggest opportunities to replace generic containers with semantic elements
- `ExtractClassNames`: Class names like nav-wrapper or footer-inner hint at semantic elements that should replace them
- `HasClearHeadingStructure`: Heading hierarchy directly depends on semantic sectioning elements being used correctly

## Scoring Axes Most Affected
- ux_clarity: Semantic HTML makes the page structure self-documenting for both developers and assistive technologies
- template_likeness: Templates built with semantic HTML feel purposeful and align with web standards
