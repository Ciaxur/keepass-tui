#!/usr/bin/env bash
set -e


SCRIPT_DIR="$(realpath $(dirname $0))"
source "$SCRIPT_DIR/common.sh"

[ "$1" == "" ] && { echo "Please provide path to a key"; exit 1; }
DATABASE_KEY_PATH="$1"

# Grab key credentials
echo -n "Enter key's password: "
read -s PASSWORD

SELECTION=$(get_entry_selection "$PASSWORD" "$DATABASE_KEY_PATH")
echo "$PASSWORD" | keepassxc-cli show "$DATABASE_KEY_PATH" "$SELECTION"
echo "$PASSWORD" | keepassxc-cli clip "$DATABASE_KEY_PATH" "$SELECTION"

