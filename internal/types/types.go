package types

type ToolResult struct {
	Text           string      `json:"text"`
	StructuredData interface{} `json:"structured_data,omitempty"`
	IsError        bool        `json:"is_error,omitempty"`
}

type AnalyzeUIContextInput struct {
	ProjectPath string `json:"project_path"`
	Task        string `json:"task,omitempty"`
}

type RepoContext struct {
	ProjectPath            string            `json:"project_path"`
	DetectedFramework      string            `json:"detected_framework"`
	UsesNextAppRouter      bool              `json:"uses_next_app_router"`
	UsesTailwind           bool              `json:"uses_tailwind"`
	UsesShadcn             bool              `json:"uses_shadcn"`
	UsesRadix              bool              `json:"uses_radix"`
	UsesDaisyUI            bool              `json:"uses_daisyui"`
	RelevantFiles          []string          `json:"relevant_files"`
	RelevantDirectories    []string          `json:"relevant_directories"`
	DesignTokens           map[string]string `json:"design_tokens"`
	SpacingPatterns        SpacingPatterns   `json:"spacing_patterns"`
	BaseComponents         []string          `json:"base_components"`
	LayoutPatterns         []string          `json:"layout_patterns"`
	VisualRisks            []string          `json:"visual_risks"`
	DetectedDependencies   []string          `json:"detected_dependencies"`
	DetectedScripts        map[string]string `json:"detected_scripts"`
	TargetFileHints        []string          `json:"target_file_hints"`
	VisualSystemSummary    string            `json:"visual_system_summary"`
	PatternInventory       PatternInventory  `json:"pattern_inventory"`
	PrimitiveCatalog       PrimitiveCatalog  `json:"primitive_catalog"`
	RouteInventory         []RouteInfo       `json:"route_inventory"`
	ShellPatterns          []string          `json:"shell_patterns"`
	ArbitrarySpacingValues []string          `json:"arbitrary_spacing_values"`
}

type SpacingPatterns struct {
	MostCommonPadding []MetricCount `json:"most_common_padding"`
	MostCommonGap     []MetricCount `json:"most_common_gap"`
	MostCommonMargin  []MetricCount `json:"most_common_margin"`
	HasArbitrary      bool          `json:"has_arbitrary"`
}

type MetricCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type PatternInventory struct {
	Cards    []string `json:"cards"`
	Forms    []string `json:"forms"`
	Modals   []string `json:"modals"`
	Tables   []string `json:"tables"`
	Tabs     []string `json:"tabs"`
	Sidebars []string `json:"sidebars"`
}

type PrimitiveCatalog struct {
	Buttons    []string `json:"buttons"`
	Inputs     []string `json:"inputs"`
	Dialogs    []string `json:"dialogs"`
	Cards      []string `json:"cards"`
	Tables     []string `json:"tables"`
	Tabs       []string `json:"tabs"`
	Sidebars   []string `json:"sidebars"`
	Forms      []string `json:"forms"`
	Navigation []string `json:"navigation"`
	Feedback   []string `json:"feedback"`
}

type RouteInfo struct {
	Route       string `json:"route"`
	PageFile    string `json:"page_file,omitempty"`
	LayoutFile  string `json:"layout_file,omitempty"`
	LoadingFile string `json:"loading_file,omitempty"`
	ErrorFile   string `json:"error_file,omitempty"`
	UsesClient  bool   `json:"uses_client"`
	HasForm     bool   `json:"has_form"`
	HasTable    bool   `json:"has_table"`
	HasSidebar  bool   `json:"has_sidebar"`
}

type BuildUISpecInput struct {
	Task          string      `json:"task"`
	Context       RepoContext `json:"context"`
	ReferenceFile string      `json:"reference_file,omitempty"`
}

type UISpec struct {
	Objective           string   `json:"objective"`
	Scope               []string `json:"scope"`
	ProbableFiles       []string `json:"probable_files"`
	ComponentsToReuse   []string `json:"components_to_reuse"`
	DesignConstraints   []string `json:"design_constraints"`
	AntiPatternsToAvoid []string `json:"anti_patterns_to_avoid"`
	AcceptanceCriteria  []string `json:"acceptance_criteria"`
	ExecutionNotes      []string `json:"execution_notes"`
	StructuredPrompt    string   `json:"structured_prompt"`
}

type CritiqueUIOutputInput struct {
	Task          string         `json:"task"`
	Context       RepoContext    `json:"context"`
	Diff          string         `json:"diff"`
	RenderRequest *RenderRequest `json:"render_request,omitempty"`
}

type RenderRequest struct {
	URL            string `json:"url"`
	WaitFor        string `json:"wait_for,omitempty"`
	ScreenshotPath string `json:"screenshot_path,omitempty"`
	ViewportWidth  int    `json:"viewport_width,omitempty"`
	ViewportHeight int    `json:"viewport_height,omitempty"`
	Command        string `json:"command,omitempty"`
	WorkingDir     string `json:"working_dir,omitempty"`
}

type PlaywrightReport struct {
	URL                 string            `json:"url"`
	ScreenshotPath      string            `json:"screenshot_path,omitempty"`
	TotalElements       int               `json:"total_elements"`
	MaxDepth            int               `json:"max_depth"`
	RepeatedWrappers    int               `json:"repeated_wrappers"`
	OverflowingElements int               `json:"overflowing_elements"`
	UniqueFontSizes     []string          `json:"unique_font_sizes"`
	UniqueGaps          []string          `json:"unique_gaps"`
	UniquePaddings      []string          `json:"unique_paddings"`
	OversizedSidebarPX  int               `json:"oversized_sidebar_px"`
	Notes               []string          `json:"notes"`
	RawSignals          map[string]string `json:"raw_signals,omitempty"`
}

type CritiqueIssue struct {
	Category string `json:"category"`
	Severity string `json:"severity"`
	Detail   string `json:"detail"`
	Action   string `json:"action"`
}

type CritiqueScores struct {
	SpacingConsistency   int `json:"spacing_consistency"`
	TypographicHierarchy int `json:"typographic_hierarchy"`
	LayoutBalance        int `json:"layout_balance"`
	ComponentReuse       int `json:"component_reuse"`
	TokenAdherence       int `json:"token_adherence"`
	DensityControl       int `json:"density_control"`
	VisualNoise          int `json:"visual_noise"`
	TemplateLikeness     int `json:"template_likeness"`
	UXClarity            int `json:"ux_clarity"`
	ProductAlignment     int `json:"product_alignment"`
}

type CritiqueReport struct {
	Summary           string            `json:"summary"`
	Problems          []CritiqueIssue   `json:"problems"`
	Scores            CritiqueScores    `json:"scores"`
	CorrectiveActions []string          `json:"corrective_actions"`
	PlaywrightReport  *PlaywrightReport `json:"playwright_report,omitempty"`
}

type ScoreUIQualityInput struct {
	CritiqueReport CritiqueReport `json:"critique_report"`
	Weights        map[string]int `json:"weights,omitempty"`
}

type ScoreAxis struct {
	Name        string `json:"name"`
	Score       int    `json:"score"`
	Explanation string `json:"explanation"`
}

type QualityScore struct {
	TotalScore           int         `json:"total_score"`
	Axes                 []ScoreAxis `json:"axes"`
	RecommendedThreshold int         `json:"recommended_threshold"`
}

type GateUIChangeInput struct {
	Score     QualityScore   `json:"score"`
	Critique  CritiqueReport `json:"critique"`
	Threshold int            `json:"threshold,omitempty"`
}

type GateResult struct {
	Approved        bool     `json:"approved"`
	Threshold       int      `json:"threshold"`
	Justification   string   `json:"justification"`
	RequiredChanges []string `json:"required_changes"`
}
