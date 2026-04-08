#! /usr/bin/env bash
# Regenerates go.work from the current module layout.
# Run from the repo root. Reads the Go version from go.mod.

set -euo pipefail
trap 'echo error: line $LINENO >&2' ERR

go_version=$(awk '/^go [0-9]/{print $2}' go.mod)
if [[ -z "$go_version" ]]; then
  echo "error: could not parse go version from go.mod" >&2
  exit 1
fi

# Find all sub-module directories (dirs containing go.mod, excluding root)
mods=()
while IFS= read -r modfile; do
  dir="${modfile%/go.mod}"
  [[ "$dir" == "." ]] && continue
  dir="${dir#./}"
  # keep only top-level dirs (no slash in remaining path)
  if [[ "$dir" != */* ]]; then
    mods+=("$dir")
  fi
done < <(find . -name go.mod -not -path '*/vendor/*' | sort)

printf '// generated do not edit.\ngo %s\n\nuse (\n\t.\n' "$go_version" > go.work
for mod in "${mods[@]}"; do
  printf '\t%s\n' "$mod" >> go.work
done
printf ')\n' >> go.work

echo "wrote go.work (go $go_version, ${#mods[@]} modules)" >&2
