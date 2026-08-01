#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)"
source "$SCRIPT_DIR/common.sh"

function print_menu() {
  cat << EOF
Menu:
1) Copy entry password
2) Show entry details
0) Exit
EOF

}

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

# Enter interactive session
print_menu
while read -r input; do
  echo

  case "$input" in
    1|2)
      # A cancelled picker returns to the menu instead of killing the session.
      if SELECTION="$(get_entry_selection "$PASSWORD" "$DATABASE_KEY_PATH")" && [ -n "$SELECTION" ]; then
        echo "$PASSWORD" | keepassxc-cli show "$DATABASE_KEY_PATH" "$SELECTION" \
          || echo "Failed to show entry" >&2
        if [ "$input" == "1" ]; then
          echo "$PASSWORD" | keepassxc-cli clip "$DATABASE_KEY_PATH" "$SELECTION" \
            || echo "Failed to copy password" >&2
        fi
      fi
      ;;
    0)
      exit 0
      ;;
    *)
      echo "Unknown entry!"
      ;;
  esac

  echo
  print_menu
done
