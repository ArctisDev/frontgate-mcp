#!/usr/bin/env python3
import argparse
import json
import pathlib
import subprocess
import sys
import tempfile


def send_request(proc, req):
    body = json.dumps(req).encode()
    proc.stdin.write(f"Content-Length: {len(body)}\r\n\r\n".encode() + body)
    proc.stdin.flush()

    header = b""
    while b"\r\n\r\n" not in header:
        chunk = proc.stdout.read(1)
        if not chunk:
            raise RuntimeError("server closed stdout before responding")
        header += chunk
    headers, _, rest = header.partition(b"\r\n\r\n")

    length = 0
    for line in headers.decode().split("\r\n"):
        if line.lower().startswith("content-length:"):
            length = int(line.split(":", 1)[1].strip())
            break
    if length <= 0:
        raise RuntimeError(f"invalid response headers: {headers!r}")

    payload = rest + proc.stdout.read(length - len(rest))
    return json.loads(payload)


def sample_project():
    root = pathlib.Path(tempfile.mkdtemp(prefix="frontgate-smoke-"))
    files = {
        "package.json": json.dumps(
            {
                "dependencies": {
                    "next": "15.0.0",
                    "react": "19.0.0",
                    "tailwindcss": "4.0.0",
                    "class-variance-authority": "0.7.0",
                    "tailwind-merge": "2.0.0",
                    "lucide-react": "0.1.0",
                    "@radix-ui/react-dialog": "1.0.0",
                }
            },
            indent=2,
        ),
        "components.json": '{"style":"default"}',
        "app/globals.css": ":root { --background: 0 0% 100%; --radius: 0.75rem; }",
        "app/layout.tsx": 'export default function RootLayout({ children }) { return <div className="min-h-screen"><header className="sticky top-0" /><Sidebar />{children}</div> }',
        "app/deploy/page.tsx": '"use client"\nexport default function DeployPage(){ return <div className="px-6 py-8 gap-4"><aside>Sidebar</aside><form /><Card /></div> }',
        "components/ui/button.tsx": "export function Button(){ return null }",
        "components/ui/input.tsx": "export function Input(){ return null }",
        "components/ui/card.tsx": "export function Card(){ return null }",
        "components/ui/sidebar.tsx": "export function Sidebar(){ return null }",
    }
    for rel, content in files.items():
        path = root / rel
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(content)
    return root


def main():
    parser = argparse.ArgumentParser(description="Smoke test the frontgate MCP server.")
    parser.add_argument(
        "--server",
        default=str(pathlib.Path(__file__).resolve().parents[1] / "frontgate-mcp"),
        help="Path to the frontgate MCP binary.",
    )
    parser.add_argument(
        "--project",
        default="",
        help="Optional project path to analyze. If omitted, a temporary Next.js sample is created.",
    )
    args = parser.parse_args()

    project_path = pathlib.Path(args.project).resolve() if args.project else sample_project()
    proc = subprocess.Popen(
        [str(pathlib.Path(args.server).resolve())],
        stdin=subprocess.PIPE,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
    )

    try:
        init = send_request(
            proc,
            {
                "jsonrpc": "2.0",
                "id": 1,
                "method": "initialize",
                "params": {
                    "protocolVersion": "2025-03-26",
                    "capabilities": {},
                    "clientInfo": {"name": "frontgate-smoke", "version": "0.1"},
                },
            },
        )

        tools = send_request(
            proc,
            {"jsonrpc": "2.0", "id": 2, "method": "tools/list", "params": {}},
        )

        analysis = send_request(
            proc,
            {
                "jsonrpc": "2.0",
                "id": 3,
                "method": "tools/call",
                "params": {
                    "name": "analyze_ui_context",
                    "arguments": {
                        "project_path": str(project_path),
                        "task": "Refatore a pagina de deploy",
                    },
                },
            },
        )

        structured = analysis["result"]["structuredContent"]
        summary = {
            "initialize": init["result"]["serverInfo"],
            "tool_names": [tool["name"] for tool in tools["result"]["tools"]],
            "analysis_summary": {
                "project_path": str(project_path),
                "framework": structured["detected_framework"],
                "target_file_hints": structured["target_file_hints"],
                "shell_patterns": structured.get("shell_patterns", []),
                "route_inventory_count": len(structured.get("route_inventory", [])),
                "primitive_catalog": structured.get("primitive_catalog", {}),
                "visual_risks": structured["visual_risks"],
            },
        }
        print(json.dumps(summary, indent=2))
    finally:
        proc.terminate()
        try:
            proc.wait(timeout=2)
        except subprocess.TimeoutExpired:
            proc.kill()


if __name__ == "__main__":
    try:
        main()
    except Exception as exc:
        print(json.dumps({"error": str(exc)}), file=sys.stderr)
        sys.exit(1)
