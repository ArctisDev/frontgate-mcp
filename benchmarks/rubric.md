# Human Evaluation Rubric

Use this rubric for blind comparison between `codex_baseline` and `codex_frontgate`.

Score each axis from `1` to `5`.

## 1. Product Alignment

- `1`: feels generic, could belong to any SaaS template
- `3`: somewhat aligned but still contains generic decisions
- `5`: clearly fits the product's actual workflow and tone

## 2. Hierarchy Clarity

- `1`: weak title/body/action hierarchy, hard to scan
- `3`: hierarchy exists but still feels flat in places
- `5`: clear information priority and strong scan path

## 3. Spacing Consistency

- `1`: arbitrary rhythm, visible inconsistency
- `3`: mostly acceptable but drifts in panels or sections
- `5`: consistent and disciplined spacing scale

## 4. Visual Noise

- `1`: too many cards, wrappers or decorative blocks
- `3`: moderate noise but still acceptable
- `5`: restrained, purposeful visual structure

## 5. Density Control

- `1`: too sparse or too cramped
- `3`: workable but not well balanced
- `5`: efficient and calm use of space

## 6. Primitive Reuse

- `1`: many ad hoc UI structures instead of repo primitives
- `3`: mixed reuse
- `5`: strong reuse of existing primitives

## 7. UX Clarity

- `1`: user intent and main action are unclear
- `3`: usable but not especially clear
- `5`: obvious next step, clear states and grouping

## 8. Overall Preference

- `1`: clearly worse
- `3`: similar / no strong preference
- `5`: clearly better

## Blind Evaluation Procedure

1. Hide which arm produced which result
2. Compare screenshots, diff summary and final UI
3. Score both arms independently
4. Record preferred variant
5. Add short free-text notes about why

## Required Notes

Each evaluator should answer:

- What felt generic?
- What felt product-specific?
- What wasted space?
- What improved task clarity?
- Which version would you ship and why?
