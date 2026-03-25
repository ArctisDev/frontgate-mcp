# Skill: Scroll Behavior

## Category
Layout

## Description
This skill teaches the AI agent to manage scroll areas so that users encounter a single, predictable scrolling axis rather than multiple nested scroll containers fighting for control. The most common scroll anti-pattern is creating nested scroll regions where both an inner container and the outer page scroll vertically, leading to confusing behavior where the user's scroll gesture is captured by the wrong element. The agent should ensure that the page body scrolls naturally and that only deliberate UI elements (like modals, dropdown menus, or code editors) introduce independent scroll containers. When an inner scroll area is necessary, it should be clearly bounded visually and should not appear to be part of the main page scroll.

The agent should avoid using `overflow: hidden` as a layout hack to clip content that overflows due to sizing mistakes. If content overflows a container, the correct fix is usually to adjust the container's sizing, use `min-height` constraints, or allow the content to flow naturally. Using `overflow: hidden` to silently clip content risks hiding important information without any visual indication to the user. If clipping is intentional (e.g., for a carousel or image crop), the agent should provide clear affordances like navigation arrows or a "show more" button so the user knows hidden content exists.

Scroll behavior should also consider scroll anchoring and smooth scrolling for in-page navigation. When the user clicks a link that scrolls to a section, the agent should use `scroll-behavior: smooth` or programmatic `scrollTo` with smooth behavior for a polished feel. However, the agent should not override the browser's default scroll behavior for the main page scroll unless the user has opted into it. Custom scroll behavior should be limited to controlled UI components like carousels and chat panels.

## Expected Outcome
- The page has a single primary scroll axis (vertical) without competing nested scroll containers
- `overflow: hidden` is not used as a layout hack to mask sizing problems
- Independent scroll containers (modals, code editors) are clearly bounded and visually distinct
- Scroll behavior for in-page navigation uses smooth scrolling where appropriate

## Heuristic Triggers
- `FindArbitrarySpacing`: Detects hardcoded height or max-height values that force scroll on elements that should size naturally
- `CountAddedWrapperDivs`: Flags wrapper divs with overflow properties that create unnecessary scroll contexts
- `CountRawPrimitiveAdditions`: Identifies new raw elements that introduce scroll containers without proper bounding or indication

## Scoring Axes Most Affected
- ux_clarity: Single scroll axis and clear scroll boundaries prevent user confusion about scroll target
- density_control: Avoiding overflow hidden hacks ensures all content is accessible and properly sized
- visual_noise: Eliminating nested scroll containers reduces the number of competing interactive regions
