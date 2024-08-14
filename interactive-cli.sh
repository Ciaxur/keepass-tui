#!/usr/bin/env bash
set -e


SCRIPT_DIR="$(realpath $(dirname $0))"
source "$SCRIPT_DIR/common.sh"

function print_menu() {
  cat << EOF
Menu:
1) Copy entry password
2) Show entry details
0) Exit
EOF

}

[ "$1" == "" ] && { echo "Please provide path to a key"; exit 1; }
DATABASE_KEY_PATH="$1"

# Grab key credentials
echo -n "Enter key's password: "
read -s PASSWORD


# Enter interactive session
print_menu
while read input; do
  echo

  case $input in
    1)
      # Grab entry from DB
      SELECTION=$(get_entry_selection "$PASSWORD" "$DATABASE_KEY_PATH")

      # Show & copy entry.
      echo "$PASSWORD" | keepassxc-cli show "$DATABASE_KEY_PATH" "$SELECTION"
      echo "$PASSWORD" | keepassxc-cli clip "$DATABASE_KEY_PATH" "$SELECTION"
      ;;
    2)
      # Grab entry from DB
      SELECTION=$(get_entry_selection "$PASSWORD" "$DATABASE_KEY_PATH")

      # Show entry.
      echo "$PASSWORD" | keepassxc-cli show "$DATABASE_KEY_PATH" "$SELECTION"
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


