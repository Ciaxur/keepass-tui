# Keepass CLI Wrappers
This repository contains wrapper scripts for a more interactive experience with [keepassxc-cli](https://www.mankier.com/1/keepassxc-cli).

# Dependencies
- [fzf](https://github.com/junegunn/fzf): Fuzzy finder CLI tool
- [golang](https://go.dev/): Golang binary for compiling small programs

# Scripts
## show_entry
Uses `fzf` to display a TUI for interactive search through the given keepass database entries, then displays the content of that entry.

## copy_pass
Uses `fzf` to display a TUI for interactive search through the given keepass database entries, then displays the content of that entry and copies the password into the clipboard.

