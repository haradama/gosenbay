#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
WEB_PUBLIC_DIR="$ROOT_DIR/web/public"

mkdir -p "$WEB_PUBLIC_DIR"

GOOS=js GOARCH=wasm go build \
  -o "$WEB_PUBLIC_DIR/senbay.wasm" \
  "$ROOT_DIR/cmd/senbay-wasm"

cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" "$WEB_PUBLIC_DIR/wasm_exec.js"

echo "Built $WEB_PUBLIC_DIR/senbay.wasm"
echo "Copied wasm_exec.js"
