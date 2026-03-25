# Skill: Minimizing JavaScript Bundle Size

## Category
Performance

## Description
JavaScript bundle size directly impacts load time, parse time, and interactivity delay. Every kilobyte of unnecessary JavaScript costs users on slow connections and mobile devices. This skill teaches AI agents to minimize bundle impact by choosing lightweight alternatives, using tree-shaking friendly import patterns, and avoiding large dependencies for simple tasks. Common violations include importing entire utility libraries like lodash when only one function is needed, or using moment.js when date-fns or native Intl APIs suffice.

Tree-shaking only works when imports are specific. `import { debounce } from 'lodash-es'` allows the bundler to exclude unused functions, while `import _ from 'lodash'` pulls in the entire library. Similarly, `import { format } from 'date-fns/format'` is tree-shakeable while `import { format } from 'date-fns'` may not be depending on configuration. Agents must learn to write imports that give the bundler maximum opportunity to eliminate dead code.

Beyond import patterns, agents should evaluate whether a library is needed at all. Many common tasks like string manipulation, simple animations, and basic data transformations can be accomplished with native browser APIs or small inline utilities rather than pulling in large dependency chains.

## Expected Outcome
- Imports use specific paths that enable tree-shaking (e.g., `lodash-es/debounce` not `lodash`)
- Large libraries are replaced with lightweight alternatives or native APIs where possible
- Bundle analysis tools identify and eliminate unused dependencies
- Shared code is extracted into common chunks rather than duplicated across routes

## Heuristic Triggers
- `CountGenericSaaSLanguage`: May indirectly reveal template-based code that includes unnecessary default dependencies
- `CountExistingComponentReuse`: Component reuse reduces bundle size by avoiding duplicate implementations

## Scoring Axes Most Affected
- component_reuse: Reusing existing components avoids adding duplicate code to the bundle
- token_adherence: Design tokens in shared utilities reduce per-component bundle overhead
