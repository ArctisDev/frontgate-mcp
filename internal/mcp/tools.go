package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ArctisDev/frontgate/internal/critic"
	"github.com/ArctisDev/frontgate/internal/guidelines"
	"github.com/ArctisDev/frontgate/internal/output"
	"github.com/ArctisDev/frontgate/internal/references"
	"github.com/ArctisDev/frontgate/internal/repo"
	"github.com/ArctisDev/frontgate/internal/scoring"
	"github.com/ArctisDev/frontgate/internal/spec"
	"github.com/ArctisDev/frontgate/internal/types"
)

type Tool struct {
	Name         string
	Description  string
	InputSchema  map[string]interface{}
	Handler      func(context.Context, json.RawMessage) (types.ToolResult, error)
	ReadOnlyHint bool
}

func NewTools() []Tool {
	return []Tool{
		{
			Name:        "analyze_ui_context",
			Description: "Analyze a Next.js + Tailwind + shadcn/ui repository, detect reusable primitives, tokens, spacing patterns and front-end system risks.",
			InputSchema: analyzeUISchema(),
			Handler: func(ctx context.Context, raw json.RawMessage) (types.ToolResult, error) {
				var in types.AnalyzeUIContextInput
				if err := json.Unmarshal(raw, &in); err != nil {
					return types.ToolResult{}, fmt.Errorf("invalid analyze_ui_context input: %w", err)
				}
				result, err := repo.Analyze(ctx, in.ProjectPath, in.Task)
				if err != nil {
					return types.ToolResult{}, err
				}
				return output.Success(result), nil
			},
			ReadOnlyHint: true,
		},
		{
			Name:        "build_ui_spec",
			Description: "Turn a vague UI task into a strict implementation spec for Codex using the analyzed repository context.",
			InputSchema: buildUISchema(),
			Handler: func(ctx context.Context, raw json.RawMessage) (types.ToolResult, error) {
				var in types.BuildUISpecInput
				if err := json.Unmarshal(raw, &in); err != nil {
					return types.ToolResult{}, fmt.Errorf("invalid build_ui_spec input: %w", err)
				}
				result := spec.Build(in)
				return output.Success(result), nil
			},
			ReadOnlyHint: true,
		},
		{
			Name:        "critique_ui_output",
			Description: "Critique a generated front-end diff using repo context, deterministic heuristics and optional Playwright DOM/CSS metrics.",
			InputSchema: critiqueUISchema(),
			Handler: func(ctx context.Context, raw json.RawMessage) (types.ToolResult, error) {
				var in types.CritiqueUIOutputInput
				if err := json.Unmarshal(raw, &in); err != nil {
					return types.ToolResult{}, fmt.Errorf("invalid critique_ui_output input: %w", err)
				}
				result, err := critic.Critique(ctx, in)
				if err != nil {
					return types.ToolResult{}, err
				}
				return output.Success(result), nil
			},
			ReadOnlyHint: true,
		},
		{
			Name:        "score_ui_quality",
			Description: "Score the technical-visual quality of the UI result across spacing, hierarchy, layout, reuse, density, noise and product alignment.",
			InputSchema: scoreUISchema(),
			Handler: func(ctx context.Context, raw json.RawMessage) (types.ToolResult, error) {
				var in types.ScoreUIQualityInput
				if err := json.Unmarshal(raw, &in); err != nil {
					return types.ToolResult{}, fmt.Errorf("invalid score_ui_quality input: %w", err)
				}
				result := scoring.Score(in)
				return output.Success(result), nil
			},
			ReadOnlyHint: true,
		},
		{
			Name:        "gate_ui_change",
			Description: "Apply a quality gate to the generated UI change and return an approval decision with required corrective actions.",
			InputSchema: gateUISchema(),
			Handler: func(ctx context.Context, raw json.RawMessage) (types.ToolResult, error) {
				var in types.GateUIChangeInput
				if err := json.Unmarshal(raw, &in); err != nil {
					return types.ToolResult{}, fmt.Errorf("invalid gate_ui_change input: %w", err)
				}
				result := scoring.Gate(in)
				return output.Success(result), nil
			},
			ReadOnlyHint: true,
		},
		{
			Name:        "get_design_guidelines",
			Description: "Get design guidelines and art direction principles from the skills corpus. Use this to understand how to create distinctive, non-generic UI designs with strong visual identity, creative layouts, rich color palettes, expressive typography, and purposeful animations.",
			InputSchema: getDesignGuidelinesSchema(),
			Handler: func(ctx context.Context, raw json.RawMessage) (types.ToolResult, error) {
				var in types.GetDesignGuidelinesInput
				if err := json.Unmarshal(raw, &in); err != nil {
					return types.ToolResult{}, fmt.Errorf("invalid get_design_guidelines input: %w", err)
				}
				result, err := guidelines.GetGuidelines(in.Categories, in.Query)
				if err != nil {
					return types.ToolResult{}, err
				}
				return output.Success(result), nil
			},
			ReadOnlyHint: true,
		},
		{
			Name:        "list_visual_references",
			Description: "List all visual reference images loaded in the references/ directory. These represent the desired visual direction and aesthetic quality for generated designs.",
			InputSchema: listVisualReferencesSchema(),
			Handler: func(ctx context.Context, raw json.RawMessage) (types.ToolResult, error) {
				var in types.ListVisualReferencesInput
				if err := json.Unmarshal(raw, &in); err != nil {
					return types.ToolResult{}, fmt.Errorf("invalid list_visual_references input: %w", err)
				}
				result, err := references.GetReferences(in.Directory)
				if err != nil {
					return types.ToolResult{}, err
				}
				return output.Success(result), nil
			},
			ReadOnlyHint: true,
		},
		{
			Name:        "add_visual_reference",
			Description: "Add a description to an existing visual reference image in the references/ directory. This enriches the art direction context used by build_ui_spec.",
			InputSchema: addVisualReferenceSchema(),
			Handler: func(ctx context.Context, raw json.RawMessage) (types.ToolResult, error) {
				var in types.AddVisualReferenceInput
				if err := json.Unmarshal(raw, &in); err != nil {
					return types.ToolResult{}, fmt.Errorf("invalid add_visual_reference input: %w", err)
				}
				if in.File == "" {
					return types.ToolResult{}, fmt.Errorf("file is required")
				}
				result := references.AddDescription(in.File, in.Description, in.Style, in.Elements)
				return output.Success(result), nil
			},
			ReadOnlyHint: false,
		},
	}
}

func ListToolDefinitions(tools []Tool) []map[string]interface{} {
	items := make([]map[string]interface{}, 0, len(tools))
	for _, tool := range tools {
		items = append(items, map[string]interface{}{
			"name":        tool.Name,
			"description": tool.Description,
			"inputSchema": tool.InputSchema,
			"annotations": map[string]bool{
				"readOnlyHint":    tool.ReadOnlyHint,
				"destructiveHint": false,
				"idempotentHint":  true,
				"openWorldHint":   false,
			},
		})
	}
	return items
}

func analyzeUISchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"project_path": map[string]interface{}{"type": "string", "description": "Absolute or relative path to the Next.js project to analyze.", "minLength": 1},
		"task":         map[string]interface{}{"type": "string", "description": "Optional task description to bias the analysis toward a target UI change."},
	}, []string{"project_path"})
}

func buildUISchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"task":           map[string]interface{}{"type": "string", "minLength": 1, "description": "Original front-end task from the user."},
		"context":        repoContextSchema("Structured output from analyze_ui_context."),
		"reference_file": map[string]interface{}{"type": "string", "description": "Optional likely target file such as app/deploy/page.tsx."},
	}, []string{"task", "context"})
}

func critiqueUISchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"task":           map[string]interface{}{"type": "string", "minLength": 1, "description": "Original front-end task from the user."},
		"context":        repoContextSchema("Structured output from analyze_ui_context."),
		"diff":           map[string]interface{}{"type": "string", "minLength": 1, "description": "Unified diff or concatenated changed file contents."},
		"render_request": renderRequestSchema(),
	}, []string{"task", "context", "diff"})
}

func scoreUISchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"critique_report": critiqueReportSchema(),
		"weights": map[string]interface{}{
			"type":                 "object",
			"description":          "Optional axis weights. Keys should match the score axis names.",
			"additionalProperties": map[string]interface{}{"type": "integer", "minimum": 1, "maximum": 100},
		},
	}, []string{"critique_report"})
}

func gateUISchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"score":     qualityScoreSchema(),
		"critique":  critiqueReportSchema(),
		"threshold": map[string]interface{}{"type": "integer", "minimum": 1, "maximum": 100, "description": "Optional explicit approval threshold. If omitted, frontgate uses the recommended threshold from the score."},
	}, []string{"score", "critique"})
}

func objectSchema(properties map[string]interface{}, required []string) map[string]interface{} {
	return map[string]interface{}{
		"type":                 "object",
		"properties":           properties,
		"required":             required,
		"additionalProperties": false,
	}
}

func renderRequestSchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"url":             map[string]interface{}{"type": "string", "minLength": 1, "description": "Local URL to render, e.g. http://localhost:3000/deploy."},
		"wait_for":        map[string]interface{}{"type": "string", "description": "Optional CSS selector to wait for before collecting metrics."},
		"screenshot_path": map[string]interface{}{"type": "string", "description": "Optional output path for a full-page screenshot."},
		"viewport_width":  map[string]interface{}{"type": "integer", "minimum": 320, "maximum": 3840},
		"viewport_height": map[string]interface{}{"type": "integer", "minimum": 320, "maximum": 3840},
		"command":         map[string]interface{}{"type": "string", "description": "Optional node binary override."},
		"working_dir":     map[string]interface{}{"type": "string", "description": "Working directory used to resolve the local playwright package."},
	}, []string{"url"})
}

func repoContextSchema(description string) map[string]interface{} {
	return map[string]interface{}{
		"type":        "object",
		"description": description,
		"required": []string{
			"project_path",
			"detected_framework",
			"uses_next_app_router",
			"uses_tailwind",
			"uses_shadcn",
			"uses_radix",
			"uses_daisyui",
			"relevant_files",
			"relevant_directories",
			"design_tokens",
			"spacing_patterns",
			"base_components",
			"layout_patterns",
			"visual_risks",
			"detected_dependencies",
			"detected_scripts",
			"target_file_hints",
			"visual_system_summary",
			"pattern_inventory",
			"primitive_catalog",
			"route_inventory",
			"shell_patterns",
			"arbitrary_spacing_values",
		},
		"properties": map[string]interface{}{
			"project_path":             map[string]interface{}{"type": "string"},
			"detected_framework":       map[string]interface{}{"type": "string"},
			"uses_next_app_router":     map[string]interface{}{"type": "boolean"},
			"uses_tailwind":            map[string]interface{}{"type": "boolean"},
			"uses_shadcn":              map[string]interface{}{"type": "boolean"},
			"uses_radix":               map[string]interface{}{"type": "boolean"},
			"uses_daisyui":             map[string]interface{}{"type": "boolean"},
			"relevant_files":           stringArraySchema(),
			"relevant_directories":     stringArraySchema(),
			"design_tokens":            stringMapSchema(),
			"spacing_patterns":         spacingPatternsSchema(),
			"base_components":          stringArraySchema(),
			"layout_patterns":          stringArraySchema(),
			"visual_risks":             stringArraySchema(),
			"detected_dependencies":    stringArraySchema(),
			"detected_scripts":         stringMapSchema(),
			"target_file_hints":        stringArraySchema(),
			"visual_system_summary":    map[string]interface{}{"type": "string"},
			"pattern_inventory":        patternInventorySchema(),
			"primitive_catalog":        primitiveCatalogSchema(),
			"route_inventory":          routeInfoArraySchema(),
			"shell_patterns":           stringArraySchema(),
			"arbitrary_spacing_values": stringArraySchema(),
		},
		"additionalProperties": false,
	}
}

func spacingPatternsSchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"most_common_padding": metricCountArraySchema(),
		"most_common_gap":     metricCountArraySchema(),
		"most_common_margin":  metricCountArraySchema(),
		"has_arbitrary":       map[string]interface{}{"type": "boolean"},
	}, []string{"most_common_padding", "most_common_gap", "most_common_margin", "has_arbitrary"})
}

func patternInventorySchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"cards":    stringArraySchema(),
		"forms":    stringArraySchema(),
		"modals":   stringArraySchema(),
		"tables":   stringArraySchema(),
		"tabs":     stringArraySchema(),
		"sidebars": stringArraySchema(),
	}, []string{"cards", "forms", "modals", "tables", "tabs", "sidebars"})
}

func primitiveCatalogSchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"buttons":    stringArraySchema(),
		"inputs":     stringArraySchema(),
		"dialogs":    stringArraySchema(),
		"cards":      stringArraySchema(),
		"tables":     stringArraySchema(),
		"tabs":       stringArraySchema(),
		"sidebars":   stringArraySchema(),
		"forms":      stringArraySchema(),
		"navigation": stringArraySchema(),
		"feedback":   stringArraySchema(),
	}, []string{"buttons", "inputs", "dialogs", "cards", "tables", "tabs", "sidebars", "forms", "navigation", "feedback"})
}

func routeInfoArraySchema() map[string]interface{} {
	return map[string]interface{}{
		"type": "array",
		"items": objectSchema(map[string]interface{}{
			"route":        map[string]interface{}{"type": "string"},
			"page_file":    map[string]interface{}{"type": "string"},
			"layout_file":  map[string]interface{}{"type": "string"},
			"loading_file": map[string]interface{}{"type": "string"},
			"error_file":   map[string]interface{}{"type": "string"},
			"uses_client":  map[string]interface{}{"type": "boolean"},
			"has_form":     map[string]interface{}{"type": "boolean"},
			"has_table":    map[string]interface{}{"type": "boolean"},
			"has_sidebar":  map[string]interface{}{"type": "boolean"},
		}, []string{"route", "uses_client", "has_form", "has_table", "has_sidebar"}),
	}
}

func critiqueReportSchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"summary":            map[string]interface{}{"type": "string"},
		"problems":           critiqueIssueArraySchema(),
		"scores":             critiqueScoresSchema(),
		"corrective_actions": stringArraySchema(),
		"playwright_report":  playwrightReportSchema(),
	}, []string{"summary", "problems", "scores", "corrective_actions"})
}

func qualityScoreSchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"total_score":           map[string]interface{}{"type": "integer", "minimum": 0, "maximum": 100},
		"axes":                  scoreAxisArraySchema(),
		"recommended_threshold": map[string]interface{}{"type": "integer", "minimum": 0, "maximum": 100},
	}, []string{"total_score", "axes", "recommended_threshold"})
}

func critiqueScoresSchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"spacing_consistency":   scoreValueSchema(),
		"typographic_hierarchy": scoreValueSchema(),
		"layout_balance":        scoreValueSchema(),
		"component_reuse":       scoreValueSchema(),
		"token_adherence":       scoreValueSchema(),
		"density_control":       scoreValueSchema(),
		"visual_noise":          scoreValueSchema(),
		"template_likeness":     scoreValueSchema(),
		"ux_clarity":            scoreValueSchema(),
		"product_alignment":     scoreValueSchema(),
	}, []string{
		"spacing_consistency",
		"typographic_hierarchy",
		"layout_balance",
		"component_reuse",
		"token_adherence",
		"density_control",
		"visual_noise",
		"template_likeness",
		"ux_clarity",
		"product_alignment",
	})
}

func playwrightReportSchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"url":                  map[string]interface{}{"type": "string"},
		"screenshot_path":      map[string]interface{}{"type": "string"},
		"total_elements":       map[string]interface{}{"type": "integer", "minimum": 0},
		"max_depth":            map[string]interface{}{"type": "integer", "minimum": 0},
		"repeated_wrappers":    map[string]interface{}{"type": "integer", "minimum": 0},
		"overflowing_elements": map[string]interface{}{"type": "integer", "minimum": 0},
		"unique_font_sizes":    stringArraySchema(),
		"unique_gaps":          stringArraySchema(),
		"unique_paddings":      stringArraySchema(),
		"oversized_sidebar_px": map[string]interface{}{"type": "integer", "minimum": 0},
		"fixed_elements":       map[string]interface{}{"type": "integer", "minimum": 0},
		"absolute_elements":    map[string]interface{}{"type": "integer", "minimum": 0},
		"scroll_containers":    map[string]interface{}{"type": "integer", "minimum": 0},
		"notes":                stringArraySchema(),
	}, []string{
		"url",
		"total_elements",
		"max_depth",
		"repeated_wrappers",
		"overflowing_elements",
		"unique_font_sizes",
		"unique_gaps",
		"unique_paddings",
		"oversized_sidebar_px",
		"fixed_elements",
		"absolute_elements",
		"scroll_containers",
		"notes",
	})
}

func critiqueIssueArraySchema() map[string]interface{} {
	return map[string]interface{}{
		"type": "array",
		"items": objectSchema(map[string]interface{}{
			"category": map[string]interface{}{"type": "string"},
			"severity": map[string]interface{}{"type": "string", "enum": []string{"low", "medium", "high"}},
			"detail":   map[string]interface{}{"type": "string"},
			"action":   map[string]interface{}{"type": "string"},
		}, []string{"category", "severity", "detail", "action"}),
	}
}

func scoreAxisArraySchema() map[string]interface{} {
	return map[string]interface{}{
		"type": "array",
		"items": objectSchema(map[string]interface{}{
			"name":        map[string]interface{}{"type": "string"},
			"score":       scoreValueSchema(),
			"explanation": map[string]interface{}{"type": "string"},
		}, []string{"name", "score", "explanation"}),
	}
}

func metricCountArraySchema() map[string]interface{} {
	return map[string]interface{}{
		"type": "array",
		"items": objectSchema(map[string]interface{}{
			"name":  map[string]interface{}{"type": "string"},
			"count": map[string]interface{}{"type": "integer", "minimum": 0},
		}, []string{"name", "count"}),
	}
}

func stringArraySchema() map[string]interface{} {
	return map[string]interface{}{
		"type":  "array",
		"items": map[string]interface{}{"type": "string"},
	}
}

func stringMapSchema() map[string]interface{} {
	return map[string]interface{}{
		"type":                 "object",
		"additionalProperties": map[string]interface{}{"type": "string"},
	}
}

func getDesignGuidelinesSchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"categories": map[string]interface{}{
			"type":        "array",
			"description": "Filter by categories: Design, Layout, Component, UX, CleanCode, Accessibility, VisualNoise, Hierarchy, ProductAlignment, Performance. Omit for all.",
			"items":       map[string]interface{}{"type": "string"},
		},
		"query": map[string]interface{}{
			"type":        "string",
			"description": "Search query to filter skills by content (e.g., 'color palette', 'animation', 'typography', 'card', 'grid').",
		},
	}, []string{})
}

func listVisualReferencesSchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"directory": map[string]interface{}{
			"type":        "string",
			"description": "Optional path to references directory. Defaults to ./references/ relative to the binary.",
		},
	}, []string{})
}

func addVisualReferenceSchema() map[string]interface{} {
	return objectSchema(map[string]interface{}{
		"file": map[string]interface{}{
			"type":        "string",
			"description": "Filename of the reference image (must already exist in references/ directory).",
			"minLength":   1,
		},
		"description": map[string]interface{}{
			"type":        "string",
			"description": "Description of what makes this design visually distinctive.",
			"minLength":   1,
		},
		"style": map[string]interface{}{
			"type":        "string",
			"description": "Style tag (e.g., 'premium-dark', 'organic-bold', 'editorial-clean').",
		},
		"elements": map[string]interface{}{
			"type":        "string",
			"description": "Key visual elements (e.g., 'gradients, 3D shapes, bold typography').",
		},
	}, []string{"file", "description"})
}

func scoreValueSchema() map[string]interface{} {
	return map[string]interface{}{"type": "integer", "minimum": 0, "maximum": 100}
}
