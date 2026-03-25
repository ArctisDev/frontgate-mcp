package guidelines

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ArctisDev/frontgate/internal/references"
)

type Skill struct {
	Category    string `json:"category"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type GuidelinesResult struct {
	TotalSkills int      `json:"total_skills"`
	Categories  []string `json:"categories"`
	Skills      []Skill  `json:"skills"`
	Summary     string   `json:"summary"`
}

func GetGuidelines(categories []string, query string) (GuidelinesResult, error) {
	skillsDir := findSkillsDir()
	if skillsDir == "" {
		return GuidelinesResult{}, nil
	}

	allSkills, err := loadAllSkills(skillsDir)
	if err != nil {
		return GuidelinesResult{}, err
	}

	filtered := filterSkills(allSkills, categories, query)
	cats := listCategories(allSkills)
	summary := buildSummary(filtered, categories, query)

	artBrief := references.GetArtBrief("")
	if artBrief != "" {
		summary = summary + "\n\n" + artBrief
	}

	return GuidelinesResult{
		TotalSkills: len(filtered),
		Categories:  cats,
		Skills:      filtered,
		Summary:     summary,
	}, nil
}

func findSkillsDir() string {
	exe, err := os.Executable()
	if err == nil {
		candidate := filepath.Join(filepath.Dir(exe), "benchmarks", "tasks")
		if info, err := os.Stat(candidate); err == nil && info.IsDir() {
			return candidate
		}
	}

	cwd, err := os.Getwd()
	if err == nil {
		candidate := filepath.Join(cwd, "benchmarks", "tasks")
		if info, err := os.Stat(candidate); err == nil && info.IsDir() {
			return candidate
		}
		candidate = filepath.Join(cwd, "..", "benchmarks", "tasks")
		if info, err := os.Stat(candidate); err == nil && info.IsDir() {
			return candidate
		}
	}

	return ""
}

func loadAllSkills(dir string) ([]Skill, error) {
	var skills []Skill

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, catEntry := range entries {
		if !catEntry.IsDir() {
			continue
		}
		category := catEntry.Name()
		catDir := filepath.Join(dir, category)

		skillDirs, err := os.ReadDir(catDir)
		if err != nil {
			continue
		}

		for _, skillDir := range skillDirs {
			if !skillDir.IsDir() {
				continue
			}
			skillFile := filepath.Join(catDir, skillDir.Name(), "skill.md")
			content, err := os.ReadFile(skillFile)
			if err != nil {
				continue
			}

			title := extractTitle(string(content))
			desc := extractDescription(string(content))

			skills = append(skills, Skill{
				Category:    category,
				Name:        skillDir.Name(),
				Title:       title,
				Description: desc,
				Content:     string(content),
			})
		}
	}

	sort.Slice(skills, func(i, j int) bool {
		if skills[i].Category != skills[j].Category {
			return skills[i].Category < skills[j].Category
		}
		return skills[i].Name < skills[j].Name
	})

	return skills, nil
}

func filterSkills(skills []Skill, categories []string, query string) []Skill {
	if len(categories) == 0 && query == "" {
		return skills
	}

	var result []Skill
	catSet := make(map[string]bool)
	for _, c := range categories {
		catSet[strings.ToLower(c)] = true
	}
	queryLower := strings.ToLower(query)

	for _, s := range skills {
		catMatch := len(catSet) == 0 || catSet[strings.ToLower(s.Category)]
		queryMatch := query == "" ||
			strings.Contains(strings.ToLower(s.Title), queryLower) ||
			strings.Contains(strings.ToLower(s.Description), queryLower) ||
			strings.Contains(strings.ToLower(s.Content), queryLower) ||
			strings.Contains(strings.ToLower(s.Name), queryLower)

		if catMatch && queryMatch {
			result = append(result, s)
		}
	}

	return result
}

func listCategories(skills []Skill) []string {
	set := make(map[string]bool)
	for _, s := range skills {
		set[s.Category] = true
	}
	var cats []string
	for c := range set {
		cats = append(cats, c)
	}
	sort.Strings(cats)
	return cats
}

func extractTitle(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# Skill: ") {
			return strings.TrimPrefix(line, "# Skill: ")
		}
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}
	return ""
}

func extractDescription(content string) string {
	lines := strings.Split(content, "\n")
	inDesc := false
	var descLines []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "## Description" {
			inDesc = true
			continue
		}
		if inDesc && strings.HasPrefix(trimmed, "## ") {
			break
		}
		if inDesc && trimmed != "" {
			descLines = append(descLines, trimmed)
		}
	}

	if len(descLines) > 2 {
		descLines = descLines[:2]
	}
	return strings.Join(descLines, " ")
}

func buildSummary(skills []Skill, categories []string, query string) string {
	if len(skills) == 0 {
		return "No matching design guidelines found."
	}

	catCount := make(map[string]int)
	for _, s := range skills {
		catCount[s.Category]++
	}

	var parts []string
	parts = append(parts, fmt.Sprintf("Found %d design guidelines across %d categories.", len(skills), len(catCount)))

	for cat, count := range catCount {
		parts = append(parts, fmt.Sprintf("- %s: %d skills", cat, count))
	}

	if query != "" {
		parts = append(parts, fmt.Sprintf("\nFiltered by query: '%s'", query))
	}

	return strings.Join(parts, "\n")
}
