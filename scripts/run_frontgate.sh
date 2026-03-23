#!/usr/bin/env bash
set -eu

ROOT="/home/silva/Documentos/GitHub/better-frontend-mcp"
LOG_FILE="/tmp/frontgate-mcp.log"

{
  echo "[$(date '+%Y-%m-%d %H:%M:%S')] starting frontgate wrapper"
  echo "cwd_before=$PWD"
  echo "root=$ROOT"
} >> "$LOG_FILE"

cd "$ROOT"
exec "$ROOT/frontgate-mcp"
