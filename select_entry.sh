#!/usr/bin/env bash
set -e

SCRIPT_DIR="$(realpath $(dirname $0))"
COMBINE_PATHS_BIN="$SCRIPT_DIR/bin/combine_dirpaths"

[ "$1" == "" ] && { echo "Please provide path to a key"; exit 1; }
DATABASE_KEY_PATH="$1"

# Grab key credentials
echo -n "Enter key's password: "
read -s PASSWORD

# Use fuzzy search to grab a selection from the database
DATABASE_PATHS=`echo "$PASSWORD" | keepassxc-cli ls -R "$DATABASE_KEY_PATH"`
SELECTION=$("$COMBINE_PATHS_BIN" <(echo "$DATABASE_PATHS") | fzf)

echo "$PASSWORD" | keepassxc-cli show "$DATABASE_KEY_PATH" "$SELECTION"

