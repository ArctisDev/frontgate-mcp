# Skill: Accurate Product State Representation

## Category
ProductAlignment

## Description
Every product has multiple data states that users encounter: loading, empty, error, populated, partial, and various intermediate states. Generic implementations treat these as afterthoughts, showing spinners for loading, "No items found" for empty, and red text for errors. This skill teaches AI agents to design state representations that feel intentional and product-aware, providing useful context and appropriate guidance for each state.

Loading states should communicate what is being loaded and, where possible, show skeleton screens that preview the expected layout. Empty states should explain why the view is empty and suggest the most likely first action. Error states should distinguish between user errors, network issues, and system failures, offering appropriate recovery paths. Populated states should handle edge cases like single items, large datasets, and partial data gracefully.

The key insight is that states are not exceptions to be handled but integral parts of the user experience that deserve the same design attention as the primary data view. Users spend significant time in empty and loading states, especially when first using a product.

## Expected Outcome
- Loading states use skeleton screens that match the expected content layout
- Empty states include contextual descriptions and clear calls-to-action specific to the view
- Error states provide actionable recovery information rather than generic error messages
- Populated states handle edge cases (single item, very large datasets, partial data) gracefully

## Heuristic Triggers
- `CountRawPrimitiveAdditions`: Can reveal when state handling uses basic HTML elements instead of purpose-built components
- `FindArbitrarySpacing`: May indicate layout issues when switching between different data states

## Scoring Axes Most Affected
- product_alignment: State handling shows understanding of the product's data model and user needs
- ux_clarity: Users understand what happened and what to do next in any state
- visual_noise: Well-designed states reduce noise; poorly designed states add confusion
