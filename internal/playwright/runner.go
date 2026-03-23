package playwright

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ArctisDev/frontgate/internal/types"
)

const scriptTemplate = `
const fs = require("fs");
const path = require("path");

async function run() {
  const input = JSON.parse(process.argv[2]);
  let playwright;
  try {
    playwright = require("playwright");
  } catch (err) {
    console.error(JSON.stringify({ error: "playwright package not found in current Node resolution path" }));
    process.exit(2);
  }

  const browser = await playwright.chromium.launch({ headless: true });
  const page = await browser.newPage({
    viewport: {
      width: input.viewport_width || 1440,
      height: input.viewport_height || 960
    }
  });
  await page.goto(input.url, { waitUntil: "networkidle" });
  if (input.wait_for) {
    await page.waitForSelector(input.wait_for, { timeout: 10000 });
  }

  const metrics = await page.evaluate(() => {
    const all = Array.from(document.querySelectorAll("*"));
    const visible = all.filter((el) => {
      const style = window.getComputedStyle(el);
      const rect = el.getBoundingClientRect();
      return style.display !== "none" && style.visibility !== "hidden" && rect.width > 0 && rect.height > 0;
    });

    let maxDepth = 0;
    let repeatedWrappers = 0;
    let overflowing = 0;
    let oversizedSidebarPX = 0;
    const fontSizes = new Set();
    const gaps = new Set();
    const paddings = new Set();
    const notes = [];

    for (const el of visible) {
      let depth = 0;
      let current = el;
      while (current.parentElement) {
        depth += 1;
        current = current.parentElement;
      }
      maxDepth = Math.max(maxDepth, depth);

      const style = window.getComputedStyle(el);
      fontSizes.add(style.fontSize);
      const gap = style.gap;
      if (gap && gap !== "normal" && gap !== "0px") gaps.add(gap);
      const padding = [style.paddingTop, style.paddingRight, style.paddingBottom, style.paddingLeft].join(" ");
      if (padding !== "0px 0px 0px 0px") paddings.add(padding);

      if (el.scrollWidth > el.clientWidth + 2) overflowing += 1;

      if (["ASIDE", "NAV"].includes(el.tagName)) {
        const rect = el.getBoundingClientRect();
        if (rect.width > oversizedSidebarPX) oversizedSidebarPX = Math.round(rect.width);
      }

      if (
        el.tagName === "DIV" &&
        el.children.length === 1 &&
        el.textContent.trim() === el.children[0].textContent.trim()
      ) {
        repeatedWrappers += 1;
      }
    }

    if (oversizedSidebarPX > 360) {
      notes.push("Oversized sidebar detected");
    }
    if (overflowing > 0) {
      notes.push("Overflowing elements detected");
    }

    return {
      total_elements: visible.length,
      max_depth: maxDepth,
      repeated_wrappers: repeatedWrappers,
      overflowing_elements: overflowing,
      unique_font_sizes: Array.from(fontSizes).sort(),
      unique_gaps: Array.from(gaps).sort(),
      unique_paddings: Array.from(paddings).sort().slice(0, 20),
      oversized_sidebar_px: oversizedSidebarPX,
      notes
    };
  });

  if (input.screenshot_path) {
    const dir = path.dirname(input.screenshot_path);
    fs.mkdirSync(dir, { recursive: true });
    await page.screenshot({ path: input.screenshot_path, fullPage: true });
  }
  await browser.close();
  console.log(JSON.stringify({
    url: input.url,
    screenshot_path: input.screenshot_path || "",
    ...metrics
  }));
}

run().catch((err) => {
  console.error(JSON.stringify({ error: err.message }));
  process.exit(1);
});
`

func Render(ctx context.Context, req types.RenderRequest) (*types.PlaywrightReport, error) {
	if req.URL == "" {
		return nil, fmt.Errorf("render_request.url is required")
	}

	tmpDir, err := os.MkdirTemp("", "frontgate-playwright-*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmpDir)

	scriptPath := filepath.Join(tmpDir, "collect.js")
	if err := os.WriteFile(scriptPath, []byte(scriptTemplate), 0o600); err != nil {
		return nil, err
	}

	cmdName := "node"
	if strings.TrimSpace(req.Command) != "" {
		cmdName = req.Command
	}
	payload, _ := json.Marshal(req)
	cmd := exec.CommandContext(ctx, cmdName, scriptPath, string(payload))
	if req.WorkingDir != "" {
		cmd.Dir = req.WorkingDir
	}
	stdout, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok && len(ee.Stderr) > 0 {
			return nil, fmt.Errorf("playwright execution failed: %s", strings.TrimSpace(string(ee.Stderr)))
		}
		return nil, fmt.Errorf("playwright execution failed: %w", err)
	}

	var raw struct {
		URL                 string   `json:"url"`
		ScreenshotPath      string   `json:"screenshot_path"`
		TotalElements       int      `json:"total_elements"`
		MaxDepth            int      `json:"max_depth"`
		RepeatedWrappers    int      `json:"repeated_wrappers"`
		OverflowingElements int      `json:"overflowing_elements"`
		UniqueFontSizes     []string `json:"unique_font_sizes"`
		UniqueGaps          []string `json:"unique_gaps"`
		UniquePaddings      []string `json:"unique_paddings"`
		OversizedSidebarPX  int      `json:"oversized_sidebar_px"`
		Notes               []string `json:"notes"`
	}
	if err := json.Unmarshal(stdout, &raw); err != nil {
		return nil, fmt.Errorf("parse playwright output: %w", err)
	}

	return &types.PlaywrightReport{
		URL:                 raw.URL,
		ScreenshotPath:      raw.ScreenshotPath,
		TotalElements:       raw.TotalElements,
		MaxDepth:            raw.MaxDepth,
		RepeatedWrappers:    raw.RepeatedWrappers,
		OverflowingElements: raw.OverflowingElements,
		UniqueFontSizes:     raw.UniqueFontSizes,
		UniqueGaps:          raw.UniqueGaps,
		UniquePaddings:      raw.UniquePaddings,
		OversizedSidebarPX:  raw.OversizedSidebarPX,
		Notes:               raw.Notes,
	}, nil
}
