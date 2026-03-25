# Skill: Loading Indicator Patterns

## Category
UX

## Description
Loading indicators are essential feedback mechanisms that tell users the system is working and their input has been received. The choice of loading pattern should match the expected duration and context. Skeleton screens — placeholder shapes that mimic the layout of incoming content — are ideal for page loads and list rendering because they reduce perceived wait time by showing structural intent. Spinners suit short, uncertain-duration operations like form submissions. Progress bars are best for measurable operations like file uploads or multi-step processes. Using the wrong indicator for the context creates confusion: a spinner on a full-page load feels abrupt, while a progress bar on a 200ms API call feels excessive.

The AI agent should learn to select loading patterns based on context and duration. Skeleton screens should match the shape and count of the content they replace — a list of five items should show five skeleton rows, not a single generic block. Spinners should be sized and positioned relative to the action they represent (inline for buttons, centered for panels). Avoid showing multiple different loading patterns simultaneously on the same screen, and never leave a user staring at a blank screen without any indication of activity.

## Expected Outcome
- Skeleton screens match the shape and quantity of content they replace
- Spinners are contextual and sized appropriately for their container
- Loading states do not persist indefinitely without timeout or error fallback
- No blank screens appear during data fetching — some feedback is always present

## Heuristic Triggers
- `FindArbitrarySpacing`: Detects awkward spacing shifts between loading and loaded states
- `CountRawPrimitiveAdditions`: Flags raw spinner implementations instead of shared loading components
- `CountNewTopLevelComponents`: Detects ad-hoc loading components instead of reusing existing patterns

## Scoring Axes Most Affected
- ux_clarity: Users must understand that content is loading, not missing
- component_reuse: Shared loading components ensure consistency across the product
