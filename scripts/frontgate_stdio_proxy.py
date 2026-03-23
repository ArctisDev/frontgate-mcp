#!/usr/bin/env python3
import os
import selectors
import subprocess
import sys
import time
from pathlib import Path


ROOT = Path("/home/silva/Documentos/GitHub/better-frontend-mcp")
SERVER = ROOT / "frontgate-mcp"
LOG = Path("/tmp/frontgate-mcp-proxy.log")


def log(message: str) -> None:
    with LOG.open("a") as fh:
        fh.write(f"[{time.strftime('%Y-%m-%d %H:%M:%S')}] {message}\n")


def preview(chunk: bytes) -> str:
    snippet = chunk[:160]
    return snippet.decode(errors="replace").replace("\r", "\\r").replace("\n", "\\n")


def main() -> int:
    log(f"proxy starting cwd={os.getcwd()} server={SERVER}")
    proc = subprocess.Popen(
        [str(SERVER)],
        cwd=str(ROOT),
        stdin=subprocess.PIPE,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        bufsize=0,
    )

    sel = selectors.DefaultSelector()
    sel.register(sys.stdin.buffer, selectors.EVENT_READ, ("stdin", proc.stdin))
    sel.register(proc.stdout, selectors.EVENT_READ, ("stdout", sys.stdout.buffer))
    sel.register(proc.stderr, selectors.EVENT_READ, ("stderr", None))

    while True:
        if proc.poll() is not None:
            log(f"server exited code={proc.returncode}")
            return proc.returncode or 0

        for key, _ in sel.select(timeout=0.5):
            source, target = key.data
            chunk = key.fileobj.read1(65536) if hasattr(key.fileobj, "read1") else key.fileobj.read(65536)
            if not chunk:
                continue

            if source == "stdin":
                log(f"stdin_bytes={len(chunk)} preview={preview(chunk)}")
                target.write(chunk)
                target.flush()
            elif source == "stdout":
                log(f"stdout_bytes={len(chunk)} preview={preview(chunk)}")
                target.write(chunk)
                target.flush()
            elif source == "stderr":
                log(f"stderr={chunk.decode(errors='replace').strip()}")


if __name__ == "__main__":
    try:
        sys.exit(main())
    except Exception as exc:
        log(f"proxy exception={exc}")
        raise
