# cbr-tmpl

## Overview

Minimal Cobra CLI template with item CRUD functionality. Demonstrates Go-idiomatic dependency injection via constructor arguments, interface-based storage abstraction, and clean architecture.

## CLI Commands

| Command | Alias | Description |
|---------|-------|-------------|
| `cbr-tmpl add <id> -c <content>` | `a` | Add a new item |
| `cbr-tmpl list` | `ls` | List all items |
| `cbr-tmpl delete <id>` | `rm` | Delete an item |

## Package Structure

```
internal/
├── storage.go              # Storage interface + error types
├── filesystem.go           # FilesystemStorage implementation
├── filesystem_test.go      # Unit tests (same package as implementation)
└── cmd/
    ├── root.go             # Root command + DI setup
    ├── root_test.go        # CLI integration tests with inline mock
    ├── add.go              # add/a command
    ├── list.go             # list/ls command
    ├── delete.go           # delete/rm command
    └── styles.go           # Lipgloss style definitions
```

## Design Rules

### Dependency Injection

Use constructor arguments for DI (no global variables):

```go
func main() {
    storage := internal.NewFilesystemStorage("")
    rootCmd := cmd.NewRootCmd(storage)
    rootCmd.Execute()
}

func NewAddCmd(storage internal.Storage) *cobra.Command {
    // Closure captures storage
    return &cobra.Command{
        RunE: func(cmd *cobra.Command, args []string) error {
            return storage.Add(args[0], content)
        },
    }
}
```

### Interface-Based Abstraction

Define interfaces for testability:

```go
type Storage interface {
    Add(id, content string) error
    List() ([]string, error)
    Delete(id string) error
    Exists(id string) bool
    Get(id string) (string, error)
}
```

### Error Handling

Use custom error types with `errors.As()`:

```go
var existsErr *internal.ItemExistsError
if errors.As(err, &existsErr) {
    // handle specific error
}
```

### Testing

- Test files live in the same package as implementation (`_test.go`)
- Use `t.TempDir()` for filesystem tests
- Mock implementations defined inline in test files
- No external testing libraries (use standard `testing` package)

### Adding New Commands

1. Create command constructor in `internal/cmd/<name>.go`
2. Accept `Storage` interface as argument
3. Register in `NewRootCmd()` with aliases
4. Add tests in `internal/cmd/root_test.go`

### Development

- `just run <args>`: Run CLI in dev mode
- `just test`: Run all tests
- `just fix`: Auto-format and fix issues (goimports + golangci-lint --fix)
- `just check`: Static checks (format check + lint + vet)

## Go Idioms Used

- Constructor-based DI (no globals)
- Flat package structure (avoid over-nesting)
- Tests colocated with implementation
- `errors.As()` for type assertions
- `t.TempDir()` for test isolation
- Standard library preferred over external deps
