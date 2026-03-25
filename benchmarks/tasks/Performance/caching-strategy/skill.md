# Skill: Effective Caching Strategy

## Category
Performance

## Description
Caching eliminates redundant computation and network requests by storing results for reuse. This skill teaches AI agents to implement caching at multiple layers: HTTP caching headers for browser-level caching, service workers for offline and repeat-visit caching, and application-level caching with SWR/React Query for in-session data reuse. Each layer serves a different purpose and agents must understand when and how to configure each one appropriately.

HTTP caching headers (`Cache-Control`, `ETag`, `Last-Modified`) control how long the browser stores and reuses responses. Static assets should use long-lived immutable caching with content-hash filenames. API responses should use appropriate `max-age` values or `stale-while-revalidate` to serve cached data while refreshing in the background. Agents must learn to distinguish between resources that change frequently (short cache or no-cache) and resources that are versioned by filename (immutable, long-lived cache).

Application-level caching with SWR or React Query provides in-memory caching with configurable revalidation strategies. The `staleTime` option determines how long cached data is considered fresh, while `cacheTime` determines how long unused data stays in memory. Proper configuration prevents unnecessary refetches while ensuring data freshness matches user expectations.

## Expected Outcome
- Static assets use immutable caching with content-hash filenames and long `max-age`
- API responses use `stale-while-revalidate` or appropriate cache durations
- SWR/React Query `staleTime` is configured to match data freshness requirements
- Service workers cache critical resources for offline capability where appropriate

## Heuristic Triggers
- `CountExistingComponentReuse`: Reused components benefit from shared cache layers
- `CountGenericSaaSLanguage`: Template-based code may use default cache settings without optimization

## Scoring Axes Most Affected
- token_adherence: Consistent patterns enable centralized cache configuration
- component_reuse: Shared components with shared caching reduce per-component overhead
