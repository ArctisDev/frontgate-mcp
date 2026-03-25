# Skill: Minimizing Network Requests

## Category
Performance

## Description
Every network request adds latency, consumes bandwidth, and requires the browser to perform DNS lookup, TCP connection, TLS negotiation, and HTTP processing. This skill teaches AI agents to minimize the number and size of network requests through batching, deduplication, and intelligent caching. Common anti-patterns include firing separate API calls for data that could be fetched in a single request, re-fetching data on every component mount, and not leveraging background refresh strategies.

Request batching combines multiple API calls into a single request, reducing round-trip overhead. GraphQL naturally supports this through query batching, while REST APIs can implement batch endpoints. Agents must learn to identify when multiple components request related data and consolidate those requests. SWR and React Query patterns provide built-in request deduplication, ensuring that multiple components requesting the same data result in a single network call with shared results.

Prefetching data for likely next actions (hovering over a link, approaching the bottom of a list) can hide network latency by loading data before the user needs it. This requires careful balance—prefetching too aggressively wastes bandwidth, while prefetching too conservatively misses optimization opportunities.

## Expected Outcome
- Related API calls are batched into single requests where the API supports it
- SWR/React Query deduplicates identical requests across components
- Prefetching loads data for likely next actions during idle time
- Request deduplication prevents redundant fetches for the same data

## Heuristic Triggers
- `CountGenericSaaSLanguage`: Generic patterns may indicate copy-paste code that duplicates API calls
- `CountExistingComponentReuse`: Reused components with shared data fetching reduce redundant requests

## Scoring Axes Most Affected
- component_reuse: Shared data fetching reduces network overhead across reused components
- token_adherence: Consistent patterns make it easier to implement centralized request management
