#!/usr/bin/env bash

SCRIPT_DIR="$(realpath $(dirname $0))"
COMBINE_PATHS_BIN="$SCRIPT_DIR/bin/combine_dirpaths"

# Use fuzzy search to grab a selection from the database
function get_entry_selection() {
  PASSWORD="$1"
  DATABASE_KEY_PATH="$2"
  DATABASE_PATHS=`echo "$PASSWORD" | keepassxc-cli ls -R "$DATABASE_KEY_PATH"`
  "$COMBINE_PATHS_BIN" <(echo "$DATABASE_PATHS") | fzf
}

