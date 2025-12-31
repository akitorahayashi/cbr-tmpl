# cbr-tmpl

Minimal Go CLI template using Cobra with item CRUD functionality.

## Installation

### Install with go install

```sh
go install github.com/akitorahayashi/cbr-tmpl/cmd/cbr-tmpl@latest
```

After installation:

```sh
cbr-tmpl --version
cbr-tmpl --help
cbr-tmpl add note1 -c "My first note"
cbr-tmpl list
cbr-tmpl delete note1
```

## Development

### Setup

```sh
git clone https://github.com/akitorahayashi/cbr-tmpl.git
cd cbr-tmpl
go mod download
```

### Run

```sh
go run ./cmd/cbr-tmpl --help
go run ./cmd/cbr-tmpl add note1 -c "Hello"
go run ./cmd/cbr-tmpl a note2 -c "World"  # alias
go run ./cmd/cbr-tmpl list
go run ./cmd/cbr-tmpl ls                   # alias
go run ./cmd/cbr-tmpl delete note1
go run ./cmd/cbr-tmpl rm note2             # alias
```

### Test and Lint

```sh
just test    # run tests
just fix     # auto-format and fix issues
just check   # static checks (format, lint, vet)
```

**Required tools:**
- `goimports`: `go install golang.org/x/tools/cmd/goimports@latest`
- `golangci-lint`: https://golangci-lint.run/welcome/install/

## Project Structure

```
cbr-tmpl/
├── cmd/
│   └── cbr-tmpl/
│       └── main.go           # entry point
├── internal/
│   ├── storage.go            # Storage interface + errors
│   ├── filesystem.go         # FilesystemStorage implementation
│   ├── filesystem_test.go    # storage unit tests
│   └── cmd/
│       ├── root.go           # root command
│       ├── root_test.go      # CLI integration tests
│       ├── add.go            # add command
│       ├── list.go           # list command
│       ├── delete.go         # delete command
│       └── styles.go         # Lipgloss styles
├── .github/
│   ├── actions/setup/        # reusable setup action
│   └── workflows/            # CI workflows
├── justfile                  # task runner
├── go.mod
└── README.md
```

## Commands

| Command | Alias | Description |
|---------|-------|-------------|
| `cbr-tmpl add <id> -c <content>` | `a` | Add a new item |
| `cbr-tmpl list` | `ls` | List all items |
| `cbr-tmpl delete <id>` | `rm` | Delete an item |

## Storage

Items are stored in `~/.config/cbr-tmpl/items/` as individual `.txt` files.
