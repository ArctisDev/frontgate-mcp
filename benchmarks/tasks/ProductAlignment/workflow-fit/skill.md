# Skill: Matching UI to User Workflows

## Category
ProductAlignment

## Description
Real users do not interact with software through abstract CRUD operations. They have specific workflows: reviewing and approving requests, triaging incoming issues, batching similar updates, comparing options before deciding. This skill teaches AI agents to design screens around how users actually accomplish goals, rather than defaulting to generic entity-list-plus-detail-page patterns that force users to reconstruct their workflow from atomic operations.

Workflow fit requires understanding the sequence and frequency of actions. The most common action should be the most accessible. Related actions should be proximate. Transitions between workflow steps should feel natural, not like navigating between disconnected pages. Agents must learn to infer likely workflows from the product type and data model, then structure the UI to minimize friction for those flows.

A common anti-pattern is building a "list everything" interface that requires users to search, filter, select, and then act, when the workflow could be streamlined with batch actions, inline editing, or status-based grouping. Another anti-pattern is separating related actions across different pages when they are typically performed together.

## Expected Outcome
- Primary workflows require minimal navigation and fewer clicks than generic CRUD patterns
- Related actions are grouped and accessible from the same context
- Screen layouts prioritize the most common workflow step, not equal space for all operations
- Batch operations are available when the workflow naturally involves acting on multiple items

## Heuristic Triggers
- `CountNewTopLevelComponents`: Can reveal over-engineered page structures that fragment workflows
- `CountCardMentions`: Detects over-reliance on card layouts when workflows might benefit from tables or lists

## Scoring Axes Most Affected
- ux_clarity: Workflow-aligned UIs make next steps obvious
- product_alignment: Reflects understanding of how the product is actually used
- density_control: Efficient workflows often require denser information display than generic layouts
