#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
WEB_PUBLIC_DIR="$ROOT_DIR/web/public"
GOROOT_DIR="$(go env GOROOT)"

mkdir -p "$WEB_PUBLIC_DIR"

GOOS=js GOARCH=wasm go build \
  -o "$WEB_PUBLIC_DIR/senbay.wasm" \
  "$ROOT_DIR/cmd/senbay-wasm"

WASM_EXEC_CANDIDATES=(
  "$GOROOT_DIR/misc/wasm/wasm_exec.js"
  "$GOROOT_DIR/lib/wasm/wasm_exec.js"
)

WASM_EXEC_SRC=""

for candidate in "${WASM_EXEC_CANDIDATES[@]}"; do
  if [[ -f "$candidate" ]]; then
    WASM_EXEC_SRC="$candidate"
    break
  fi
done

if [[ -z "$WASM_EXEC_SRC" ]]; then
  echo "wasm_exec.js was not found under $GOROOT_DIR" >&2
  echo "Checked:" >&2
  printf '  - %s\n' "${WASM_EXEC_CANDIDATES[@]}" >&2
  exit 1
fi

cp "$WASM_EXEC_SRC" "$WEB_PUBLIC_DIR/wasm_exec.js"

echo "Built $WEB_PUBLIC_DIR/senbay.wasm"
echo "Copied $WASM_EXEC_SRC to $WEB_PUBLIC_DIR/wasm_exec.js"
