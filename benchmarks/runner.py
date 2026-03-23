#!/usr/bin/env python3
import argparse
import datetime as dt
import json
from pathlib import Path
import shutil
import sys


ROOT = Path(__file__).resolve().parent
TASKS_DIR = ROOT / "tasks"
RUNS_DIR = ROOT / "runs"
TEMPLATES_DIR = ROOT / "templates"


def list_tasks() -> int:
    for path in sorted(TASKS_DIR.glob("*.md")):
        print(path.stem)
    return 0


def create_run(task_id: str, arm: str, project_path: str) -> int:
    timestamp = dt.datetime.now().strftime("%Y%m%d-%H%M%S")
    run_id = f"{timestamp}-{task_id}-{arm}"
    run_dir = RUNS_DIR / run_id
    run_dir.mkdir(parents=True, exist_ok=False)

    result_template = json.loads((TEMPLATES_DIR / "result.template.json").read_text())
    result_template["run_id"] = run_id
    result_template["task_id"] = task_id
    result_template["arm"] = arm
    result_template["project_path"] = project_path
    result_template["started_at"] = dt.datetime.now().isoformat()

    (run_dir / "result.json").write_text(json.dumps(result_template, indent=2))
    shutil.copyfile(TEMPLATES_DIR / "run_notes.template.md", run_dir / "notes.md")

    meta = {
        "run_id": run_id,
        "task_id": task_id,
        "arm": arm,
        "project_path": project_path,
        "task_file": str((TASKS_DIR / f"{task_id}.md").resolve()),
    }
    (run_dir / "meta.json").write_text(json.dumps(meta, indent=2))

    print(run_dir)
    return 0


def summarize() -> int:
    rows = []
    for result_file in sorted(RUNS_DIR.glob("*/result.json")):
        try:
            data = json.loads(result_file.read_text())
        except Exception:
            continue
        rows.append(
            {
                "run_id": data.get("run_id", ""),
                "task_id": data.get("task_id", ""),
                "arm": data.get("arm", ""),
                "status": data.get("status", ""),
                "iterations": data.get("metrics", {}).get("iterations_total", 0),
                "score": data.get("metrics", {}).get("frontgate_total_score", None),
                "preference": data.get("human_evaluation", {}).get("overall_preference", 0),
            }
        )

    print(json.dumps(rows, indent=2))
    return 0


def main(argv: list[str]) -> int:
    parser = argparse.ArgumentParser(description="Frontgate benchmark runner")
    sub = parser.add_subparsers(dest="command", required=True)

    sub.add_parser("list-tasks")

    create = sub.add_parser("create")
    create.add_argument("--task", required=True, dest="task_id")
    create.add_argument("--arm", required=True, choices=["codex_baseline", "codex_frontgate"])
    create.add_argument("--project", required=True, dest="project_path")

    sub.add_parser("summarize")

    args = parser.parse_args(argv)

    if args.command == "list-tasks":
        return list_tasks()
    if args.command == "create":
        task_file = TASKS_DIR / f"{args.task_id}.md"
        if not task_file.exists():
            print(f"task not found: {args.task_id}", file=sys.stderr)
            return 1
        return create_run(args.task_id, args.arm, args.project_path)
    if args.command == "summarize":
        return summarize()
    return 1


if __name__ == "__main__":
    raise SystemExit(main(sys.argv[1:]))
