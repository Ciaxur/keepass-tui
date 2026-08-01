#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)"
source "$SCRIPT_DIR/common.sh"

if [ $# -lt 1 ] || [ -z "$1" ]; then
  echo "Usage: $(basename "$0") <path-to-kdbx>" >&2
  exit 1
fi
DATABASE_KEY_PATH="$1"

require_deps || exit 1

# Grab key credentials
read_password "Enter key's password: "

# Verify credentials are gtg.
verify_credentials "$PASSWORD" "$DATABASE_KEY_PATH" || {
  echo "Failed to unlock database" >&2; exit 1;
}

# Cancelling the picker is not an error.
SELECTION="$(get_entry_selection "$PASSWORD" "$DATABASE_KEY_PATH")" || exit 0
[ -n "$SELECTION" ] || exit 0

echo "$PASSWORD" | keepassxc-cli show "$DATABASE_KEY_PATH" "$SELECTION"
