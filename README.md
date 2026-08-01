# Keepass CLI Wrappers
This repository contains wrapper scripts for a more interactive experience with [keepassxc-cli](https://www.mankier.com/1/keepassxc-cli).

# Dependencies
- [keepassxc-cli](https://www.mankier.com/1/keepassxc-cli): KeePassXC command line client
- [fzf](https://github.com/junegunn/fzf): Fuzzy finder CLI tool
- [golang](https://go.dev/): Golang binary for compiling small programs

# Build
The scripts depend on the `combine_dirpaths` helper, which is not checked in.
Build it before first use:

```sh
make        # builds bin/combine_dirpaths
make test   # runs the parser tests
```

# Scripts
All scripts take the path to a `.kdbx` database as their only argument and
prompt for its password once.

## show_entry
Uses `fzf` to display a TUI for interactive search through the given keepass database entries, then displays the content of that entry.

```sh
./show_entry.sh ~/secrets.kdbx
```

## copy_pass
Same as `show_entry`, but also copies the entry's password to the clipboard.

```sh
./copy_pass.sh ~/secrets.kdbx
```

## interactive-cli
Menu-driven session that keeps the database unlocked, so you can look up several
entries without re-entering the password.

```sh
./interactive-cli.sh ~/secrets.kdbx
```

# Helpers
## combine_dirpaths
`keepassxc-cli ls -R` prints entries as an indented tree, but `keepassxc-cli show`
expects a full `Group/Subgroup/entry` path. This Go program reads the tree on
stdin and writes the flattened paths on stdout.
