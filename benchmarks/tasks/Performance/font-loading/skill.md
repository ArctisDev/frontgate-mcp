# Skill: Efficient Font Loading

## Category
Performance

## Description
Web fonts are render-blocking resources that can significantly delay First Contentful Paint (FCP) if not handled properly. This skill teaches AI agents to implement efficient font loading strategies that balance visual quality with performance. The foundation is using `font-display: swap` (or `optional` for non-critical fonts) to prevent invisible text during font loading. Beyond that, agents should understand font preloading, subsetting, and the trade-offs between font variety and performance.

Font preloading with `<link rel="preload">` tells the browser to fetch critical font files early in the page load, before they are discovered through CSS parsing. This is essential for above-the-fold text that uses custom fonts. However, preloading too many fonts wastes bandwidth and can delay other critical resources. Agents must learn to preload only the font weights and styles used in the initial viewport.

Limiting font variants is a practical optimization. Each font weight (300, 400, 500, 600, 700) and style (normal, italic) is a separate file. Using only 2-3 weights instead of the full range can reduce font payload by 60-70%. Variable fonts offer a middle ground, providing weight flexibility in a single file, though file sizes are larger than a single static weight.

## Expected Outcome
- All font declarations include `font-display: swap` or `font-display: optional`
- Critical above-the-fold fonts are preloaded with `<link rel="preload">`
- Font variants are limited to the weights and styles actually used
- System font stacks are used as fallbacks to minimize layout shift

## Heuristic Triggers
- `FindHardcodedPaletteClasses`: May reveal custom font classes that add unnecessary font variants
- `FindArbitrarySpacing`: Font loading can cause layout shifts that appear as spacing issues

## Scoring Axes Most Affected
- typographic_hierarchy: Font loading affects how quickly text hierarchy is visible
- visual_noise: Font swap (FOUT) or invisible text (FOIT) creates visual disruption
