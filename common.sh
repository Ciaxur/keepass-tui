#!/usr/bin/env bash

# BASH_SOURCE, not $0: this file is sourced, so $0 is the caller's path.
COMMON_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)"
COMBINE_PATHS_BIN="$COMMON_DIR/bin/combine_dirpaths"

# Fail with a useful message instead of an empty picker.
function require_deps() {
  local missing=0
  for cmd in keepassxc-cli fzf; do
    command -v "$cmd" > /dev/null || { echo "Missing dependency: $cmd" >&2; missing=1; }
  done
  if [ ! -x "$COMBINE_PATHS_BIN" ]; then
    echo "Missing $COMBINE_PATHS_BIN - run 'make' to build it" >&2
    missing=1
  fi
  return "$missing"
}

# Prompt for the database password without echoing it. Sets $PASSWORD.
function read_password() {
  echo -n "$1" >&2
  read -rs PASSWORD
  echo >&2
}

# Catch a wrong password before we show an empty entry list.
function verify_credentials() {
  echo "$1" | keepassxc-cli ls "$2" > /dev/null
}

# Use fuzzy search to grab a selection from the database.
# Returns non-zero when the listing fails or the user cancels the picker.
function get_entry_selection() {
  local password="$1" database="$2" listing
  listing="$(echo "$password" | keepassxc-cli ls -R "$database")" || return 1
  printf '%s\n' "$listing" | "$COMBINE_PATHS_BIN" | fzf
}
