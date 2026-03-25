package references

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Reference struct {
	File        string `json:"file"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Style       string `json:"style"`
	Elements    string `json:"elements"`
}

type ReferencesMeta struct {
	References []Reference `json:"references"`
	ArtBrief   string      `json:"art_brief"`
}

type ReferencesResult struct {
	Count     int         `json:"count"`
	Directory string      `json:"directory"`
	ArtBrief  string      `json:"art_brief"`
	Refs      []Reference `json:"references"`
}

func GetReferences(refsDir string) (ReferencesResult, error) {
	if refsDir == "" {
		refsDir = findRefsDir()
	}
	if refsDir == "" {
		return ReferencesResult{}, nil
	}

	meta, err := loadMeta(refsDir)
	if err != nil {
		meta = ReferencesMeta{}
	}

	files, err := listImageFiles(refsDir)
	if err != nil {
		return ReferencesResult{}, err
	}

	refs := mergeReferences(files, meta.References, refsDir)
	artBrief := meta.ArtBrief
	if artBrief == "" {
		artBrief = buildDefaultBrief(refs)
	}

	return ReferencesResult{
		Count:     len(refs),
		Directory: refsDir,
		ArtBrief:  artBrief,
		Refs:      refs,
	}, nil
}

func GetArtBrief(refsDir string) string {
	if refsDir == "" {
		refsDir = findRefsDir()
	}
	if refsDir == "" {
		return ""
	}

	meta, err := loadMeta(refsDir)
	if err == nil && meta.ArtBrief != "" {
		return meta.ArtBrief
	}

	files, err := listImageFiles(refsDir)
	if err != nil || len(files) == 0 {
		return ""
	}

	refs := mergeReferences(files, nil, refsDir)
	return buildDefaultBrief(refs)
}

func findRefsDir() string {
	exe, err := os.Executable()
	if err == nil {
		candidate := filepath.Join(filepath.Dir(exe), "references")
		if info, err := os.Stat(candidate); err == nil && info.IsDir() {
			return candidate
		}
	}

	cwd, err := os.Getwd()
	if err == nil {
		candidate := filepath.Join(cwd, "references")
		if info, err := os.Stat(candidate); err == nil && info.IsDir() {
			return candidate
		}
	}

	return ""
}

func loadMeta(dir string) (ReferencesMeta, error) {
	metaFile := filepath.Join(dir, "references.json")
	data, err := os.ReadFile(metaFile)
	if err != nil {
		return ReferencesMeta{}, err
	}

	var meta ReferencesMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return ReferencesMeta{}, err
	}
	return meta, nil
}

func listImageFiles(dir string) ([]string, error) {
	var files []string
	exts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true, ".gif": true, ".svg": true}
	if err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(d.Name()))
		if !exts[ext] {
			return nil
		}
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			rel = d.Name()
		}
		rel = filepath.ToSlash(rel)
		files = append(files, rel)
		return nil
	}); err != nil {
		return nil, err
	}
	sort.Strings(files)
	return files, nil
}

func mergeReferences(files []string, metaRefs []Reference, dir string) []Reference {
	metaMap := make(map[string]Reference)
	for _, r := range metaRefs {
		metaMap[r.File] = r
	}

	var refs []Reference
	for _, f := range files {
		if m, ok := metaMap[f]; ok {
			refs = append(refs, m)
		} else {
			refs = append(refs, Reference{
				File: f,
				Path: filepath.Join(dir, f),
			})
		}
	}
	return refs
}

func buildDefaultBrief(refs []Reference) string {
	if len(refs) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("VISUAL REFERENCES LOADED:\n")
	b.WriteString(fmt.Sprintf("The user has provided %d design reference images. ", len(refs)))
	b.WriteString("These represent the visual direction and aesthetic quality expected.\n\n")

	described := 0
	for _, r := range refs {
		if r.Description != "" {
			described++
		}
	}

	if described > 0 {
		b.WriteString("Described references:\n")
		for _, r := range refs {
			if r.Description != "" {
				b.WriteString(fmt.Sprintf("- %s: %s", r.File, r.Description))
				if r.Style != "" {
					b.WriteString(fmt.Sprintf(" [Style: %s]", r.Style))
				}
				if r.Elements != "" {
					b.WriteString(fmt.Sprintf(" [Elements: %s]", r.Elements))
				}
				b.WriteString("\n")
			}
		}
	} else {
		b.WriteString("NOTE: References lack descriptions. To activate full art direction context, create a 'references.json' file in the references/ directory with descriptions for each image.\n")
		b.WriteString("Format: {\"references\": [{\"file\": \"image.jpg\", \"description\": \"...\", \"style\": \"...\", \"elements\": \"...\"}], \"art_brief\": \"...\"}\n")
	}

	return b.String()
}

func AddDescription(file, description, style, elements string) string {
	refsDir := findRefsDir()
	if refsDir == "" {
		return "No references/ directory found. Create one and add image files first."
	}

	metaFile := filepath.Join(refsDir, "references.json")
	var meta ReferencesMeta

	data, err := os.ReadFile(metaFile)
	if err == nil {
		json.Unmarshal(data, &meta)
	}

	found := false
	for i, r := range meta.References {
		if r.File == file {
			meta.References[i].Description = description
			if style != "" {
				meta.References[i].Style = style
			}
			if elements != "" {
				meta.References[i].Elements = elements
			}
			found = true
			break
		}
	}

	if !found {
		meta.References = append(meta.References, Reference{
			File:        file,
			Path:        filepath.Join(refsDir, file),
			Description: description,
			Style:       style,
			Elements:    elements,
		})
	}

	out, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error saving: %v", err)
	}

	if err := os.WriteFile(metaFile, out, 0644); err != nil {
		return fmt.Sprintf("Error writing file: %v", err)
	}

	return fmt.Sprintf("Reference '%s' updated with description: %s", file, description)
}
