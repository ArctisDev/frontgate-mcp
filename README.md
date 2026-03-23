# Frontgate MCP

Local-first MCP server in Go that analyzes Next.js App Router + Tailwind + shadcn/ui repositories and produces structured specs, critiques, and quality gates for AI coding agents like OpenCode and Codex.

## What It Does

Frontgate acts as a frontend quality layer between you and your AI coding agent. It:

- **Analyzes** your Next.js project's design tokens, spacing patterns, reusable components, routes, and shell layouts
- **Builds specs** from natural language tasks so Codex/OpenCode generates code that follows your existing patterns
- **Critiques** generated UI diffs for spacing issues, component overuse, generic language, and missing reuse
- **Scores** output across 10 weighted quality axes (spacing, hierarchy, layout balance, component reuse, token adherence, density, noise, template-likeness, UX clarity, product alignment)
- **Gates** changes with threshold-based approval/rejection and required corrective actions

## Requirements

- **Go 1.22+**
- **Node.js** (for Playwright MCP)
- **Playwright MCP** — required for the `critique_ui_output` tool to perform DOM/CSS render analysis. Without it, critique still works but skips the optional render analysis.

## Build

```bash
go build -o frontgate-mcp ./cmd/frontgate-mcp
```

## Test

```bash
go test ./...
```

## Smoke Test

```bash
python3 scripts/smoke_mcp.py
```

Or against a specific project:

```bash
python3 scripts/smoke_mcp.py --project /absolute/path/to/your/next-app
```

## MCP Tools

| Tool | Description |
|------|-------------|
| `analyze_ui_context` | Scans a Next.js project for framework, design tokens, spacing, primitives, routes, shell patterns, and visual risks |
| `build_ui_spec` | Converts a natural language task into a structured implementation spec for Codex with component reuse suggestions, constraints, and acceptance criteria |
| `critique_ui_output` | Analyzes a generated UI diff for arbitrary spacing, generic language, card overuse, missing reuse, and optional Playwright DOM/CSS metrics |
| `score_ui_quality` | Weighted scoring across 10 quality axes with axis-level explanations |
| `gate_ui_change` | Applies a quality threshold and blocks high-severity issues, returning required corrective actions |

## Configuration

### OpenCode

The project includes an `opencode.json` at the repo root. Add the `frontgate` and `playwright` blocks to your own `opencode.json`:

```json
{
  "mcp": {
    "frontgate": {
      "type": "local",
      "command": ["/absolute/path/to/frontgate-mcp"],
      "enabled": true,
      "timeout": 120000
    },
    "playwright": {
      "type": "local",
      "command": ["npx", "@playwright/mcp@latest", "--isolated"],
      "enabled": true,
      "timeout": 25000
    }
  }
}
```

> **Important:** The `playwright` MCP server is required for `critique_ui_output` to run its optional DOM/CSS render analysis. If you omit it, critique will still work but will skip the Playwright-based analysis step.

### Codex

Add to `~/.codex/config.toml`:

```toml
[mcp_servers.frontgate]
command = "/absolute/path/to/frontgate-mcp"
startup_timeout_sec = 15
tool_timeout_sec = 120
```

For Playwright support in Codex, configure the Playwright MCP server in your Codex setup as well.

## Project Structure

```
frontgate-mcp/
  cmd/frontgate-mcp/        # Entry point (main.go)
  internal/
    mcp/                     # MCP server, tool registration
    types/                   # Shared types
    repo/                    # Repository analyzer
    spec/                    # Task-to-spec builder
    critic/                  # Diff critique engine
    scoring/                 # Weighted quality scoring
    rules/                   # Heuristic rules
    parser/                  # package.json and token parsers
    playwright/              # Optional Playwright render runner
    output/                  # Output formatting
  scripts/
    smoke_mcp.py             # Smoke test harness
    frontgate_stdio_proxy.py # stdio proxy for debugging
    run_frontgate.sh         # Launch helper
  benchmarks/                # Benchmarks
```

## How It Works With AI Agents

1. **Analyze** — The agent calls `analyze_ui_context` with your project path to understand your codebase's patterns
2. **Spec** — The agent calls `build_ui_spec` with a task description to get a structured spec with constraints and acceptance criteria
3. **Generate** — The agent (Codex/OpenCode) implements the spec
4. **Critique** — The agent calls `critique_ui_output` with the generated diff
5. **Score** — The agent calls `score_ui_quality` to get weighted scores
6. **Gate** — The agent calls `gate_ui_change` to check if the output passes quality thresholds

## Supported Stack

Currently targets:

- Next.js App Router
- Tailwind CSS
- shadcn/ui (with Radix UI primitives)

DaisyUI detection is supported. Other frameworks may partially work but are not officially supported.

## Known Limitations

- Narrow to Next.js App Router + Tailwind + shadcn/ui
- Route understanding is structural, not semantic
- Token extraction is strongest when CSS variables exist
- Playwright render analysis depends on the target project having a usable local runtime and the `playwright` package
- Critique is heuristic-based, not a substitute for product judgment
