package repo

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestAnalyzeDetectsNextTailwindAndShadcn(t *testing.T) {
	root := t.TempDir()
	mustWrite(t, filepath.Join(root, "package.json"), `{
	  "dependencies": {
	    "next": "15.0.0",
	    "react": "19.0.0",
	    "tailwindcss": "4.0.0",
	    "@radix-ui/react-dialog": "1.0.0",
	    "class-variance-authority": "0.7.0",
	    "tailwind-merge": "2.0.0",
	    "lucide-react": "0.1.0"
	  },
	  "scripts": {
	    "dev": "next dev"
	  }
	}`)
	mustWrite(t, filepath.Join(root, "components.json"), `{"style":"default"}`)
	mustWrite(t, filepath.Join(root, "app", "globals.css"), `:root { --background: 0 0% 100%; --radius: 0.75rem; }`)
	mustWrite(t, filepath.Join(root, "app", "layout.tsx"), `
	export default function RootLayout({ children }) {
	  return <div className="min-h-screen"><header className="sticky top-0" /><Sidebar />{children}</div>
	}
	`)
	mustWrite(t, filepath.Join(root, "app", "deploy", "page.tsx"), `
	"use client"
	export default function DeployPage() {
	  return <div className="px-6 py-8 gap-4"><aside>Sidebar</aside><form /><div className="grid grid-cols-2"><Card /></div></div>
	}
	`)
	mustWrite(t, filepath.Join(root, "components", "ui", "button.tsx"), `export function Button(){ return null }`)
	mustWrite(t, filepath.Join(root, "components", "ui", "card.tsx"), `export function Card(){ return null }`)
	mustWrite(t, filepath.Join(root, "components", "ui", "input.tsx"), `export function Input(){ return null }`)
	mustWrite(t, filepath.Join(root, "components", "ui", "sidebar.tsx"), `export function Sidebar(){ return null }`)

	got, err := Analyze(context.Background(), root, "Refatore a pagina de deploy")
	if err != nil {
		t.Fatalf("analyze failed: %v", err)
	}
	if got.DetectedFramework != "Next.js" {
		t.Fatalf("expected Next.js, got %q", got.DetectedFramework)
	}
	if !got.UsesTailwind || !got.UsesShadcn || !got.UsesRadix || !got.UsesNextAppRouter {
		t.Fatalf("unexpected detection flags: %#v", got)
	}
	if got.DesignTokens["background"] == "" {
		t.Fatalf("expected design tokens, got %#v", got.DesignTokens)
	}
	if len(got.TargetFileHints) == 0 {
		t.Fatalf("expected target file hints")
	}
	if len(got.RouteInventory) == 0 || got.RouteInventory[0].Route == "" {
		t.Fatalf("expected route inventory, got %#v", got.RouteInventory)
	}
	if len(got.PrimitiveCatalog.Buttons) == 0 || len(got.PrimitiveCatalog.Inputs) == 0 {
		t.Fatalf("expected primitive catalog to be populated, got %#v", got.PrimitiveCatalog)
	}
	if len(got.ShellPatterns) == 0 {
		t.Fatalf("expected shell patterns, got %#v", got.ShellPatterns)
	}
}

func mustWrite(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write failed: %v", err)
	}
}
