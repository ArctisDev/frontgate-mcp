# Skill: Reserving Badges for Important Indicators

## Category
VisualNoise

## Description
Badges are compact visual markers designed to surface status, counts, or categories at a glance. They are effective for unread counts, status states like "Active" or "Pending", and severity levels. However, agents often badge everything — labeling every field, tagging every section, adding count badges to every nav item — until badges lose their informational weight. When everything is badged, nothing stands out.

This skill teaches reserving badges for information that requires immediate recognition: notification counts, critical status states, and key categorical labels. Do not use badges as a general-purpose label mechanism. If a piece of information can be communicated through plain text or a subtle color change, prefer that over a pill-shaped badge container. Badges should interrupt the visual flow only when the information demands it.

## Expected Outcome
- Badges limited to notification counts, status states, and critical labels
- Informational labels rendered as plain text or subtle indicators
- Badge colors tied to semantic meaning (danger, warning, success)
- No decorative or redundant badge usage

## Heuristic Triggers
- `CountCardMentions`: Badge-heavy layouts often correlate with card-heavy nesting
- `CountGenericSaaSLanguage`: Generic label text often accompanies unnecessary badges

## Scoring Axes Most Affected
- visual_noise: Fewer badges means fewer small colored shapes competing for attention
- ux_clarity: Badges used only for important signals improve scanability
