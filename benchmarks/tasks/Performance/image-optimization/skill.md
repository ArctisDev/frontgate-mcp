# Skill: Image Optimization for Web

## Category
Performance

## Description
Images are often the largest assets on a web page and the primary contributor to slow load times. This skill teaches AI agents to implement proper image optimization using modern tools and formats. Agents must learn to use Next.js's built-in `next/image` component (or equivalent framework image components) which handles automatic resizing, format negotiation, and lazy loading. Unoptimized images loaded through raw `<img>` tags bypass all these optimizations and directly impact Core Web Vitals scores, particularly Largest Contentful Paint (LCP).

Beyond component selection, agents must understand responsive image sizing using `srcset` and `sizes` attributes, serving appropriately dimensioned images for different viewports rather than sending a single large image to all devices. Modern formats like WebP and AVIF offer 25-50% size reductions over JPEG/PNG with equivalent visual quality. Agents should also learn to distinguish between above-the-fold images (which should be eagerly loaded with `priority`) and below-the-fold images (which should be lazy loaded by default).

The combination of proper sizing, modern formats, and appropriate loading strategy can reduce page weight by megabytes on image-heavy pages and dramatically improve perceived performance.

## Expected Outcome
- All images use framework-optimized components (`next/image`) instead of raw `<img>` tags
- Responsive `srcset` and `sizes` attributes serve appropriate image dimensions per viewport
- Modern formats (WebP, AVIF) are used with appropriate fallbacks
- Above-the-fold images use `priority` loading; below-the-fold images are lazy loaded

## Heuristic Triggers
- `CountRawPrimitiveAdditions`: Detects raw `<img>` elements that bypass image optimization
- `FindArbitrarySpacing`: Can reveal layout shifts caused by missing image dimensions

## Scoring Axes Most Affected
- visual_noise: Optimized images load faster, reducing perceived visual disruption
- density_control: Proper image sizing prevents layout shifts that distort density
