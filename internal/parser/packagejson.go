package parser

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type PackageJSON struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	Scripts         map[string]string `json:"scripts"`
}

func LoadPackageJSON(projectPath string) (PackageJSON, error) {
	var pkg PackageJSON
	body, err := os.ReadFile(filepath.Join(projectPath, "package.json"))
	if err != nil {
		return pkg, err
	}
	err = json.Unmarshal(body, &pkg)
	return pkg, err
}
