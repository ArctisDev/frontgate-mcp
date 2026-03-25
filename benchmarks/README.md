# Frontgate Benchmark

This benchmark exists to answer one question:

`Does Frontgate make Codex produce better front-end output in a repeatable way?`

The benchmark compares two arms:

- `codex_baseline`: Codex without Frontgate
- `codex_frontgate`: Codex using `analyze_ui_context`, `build_ui_spec`, `critique_ui_output`, `score_ui_quality`, and `gate_ui_change`

## Benchmark Goal

Validate whether Frontgate improves:

- prompt quality
- product alignment
- visual hierarchy
- spacing consistency
- primitive reuse
- iteration efficiency
- human preference

## Structure

```text
benchmarks/
  README.md
  rubric.md
  runner.py
  tasks/
  templates/
  runs/
```

## Benchmark Corpus

The corpus contains 10 realistic front-end tasks designed to trigger weak Codex behavior:

1. deploy page looks generic and over-boxed
2. billing page is noisy and weakly grouped
3. settings form is long, flat and hard to scan
4. empty state is decorative but unclear
5. dashboard cards dominate instead of supporting decisions
6. table page wastes space and collapses hierarchy
7. sidebar consumes too much width
8. onboarding page looks like a template instead of product UX
9. analytics screen has weak information priority
10. modal flow is visually inconsistent with the rest of the app

## Evaluation Arms

### Arm A: `codex_baseline`

Run Codex with the task and normal repository context only.

No Frontgate tools.

### Arm B: `codex_frontgate`

Run:

1. `analyze_ui_context`
2. `build_ui_spec`
3. Codex implementation
4. `critique_ui_output`
5. `score_ui_quality`
6. `gate_ui_change`
7. iterate until accepted or budget exhausted

## Metrics

### Quantitative

- `iterations_total`
- `files_changed`
- `lines_added`
- `lines_removed`
- `arbitrary_spacing_count`
- `primitive_reuse_count`
- `frontgate_total_score`
- `frontgate_gate_pass`
- `time_minutes`

### Human Evaluation

See [rubric.md](rubric.md).

Each run should be rated blindly on:

- product alignment
- hierarchy clarity
- spacing consistency
- visual noise
- density control
- overall preference

## Success Criteria

Frontgate is meaningfully validated if, across the corpus:

- human preference favors `codex_frontgate` on at least 7/10 tasks
- average product-alignment score improves by at least 1 point on a 1-5 scale
- arbitrary spacing usage is lower
- primitive reuse is higher
- accepted output quality improves without exploding iteration count

## Running the Benchmark

Create a run scaffold:

```bash
python3 benchmarks/runner.py create \
  --task deploy-page-generic \
  --arm codex_frontgate \
  --project /absolute/path/to/target/project
```

This prints the run directory and creates:

- run metadata
- result skeleton
- notes file

List available tasks:

```bash
python3 benchmarks/runner.py list-tasks
```

Summarize collected results:

```bash
python3 benchmarks/runner.py summarize
```

## Recommended Workflow Per Task

1. Choose one task from `benchmarks/tasks`
2. Run `codex_baseline`
3. Save screenshots, diff, and result JSON
4. Run `codex_frontgate`
5. Save screenshots, diff, critique and score
6. Apply blind human evaluation using the same rubric
7. Repeat for all 10 tasks

## Required Artifacts Per Run

- before screenshot
- after screenshot
- unified diff
- final prompt used by Codex
- result JSON
- evaluator notes

## What This Benchmark Is Trying To Prove

Not that Frontgate can detect problems.

That part is already implemented.

This benchmark is trying to prove that Frontgate changes the trajectory of generation in a useful, measurable and repeatable way.
