# Go Import Cycle Demonstration

This project is a simple example designed to demonstrate Go's "import cycle not allowed" error. It showcases how Go prevents circular dependencies between packages, which is a key feature for maintaining a clean and modular architecture.

## Project Structure

- **`cmd/example/main.go`**: The entry point of the application. It imports and uses the `app` package.
- **`internal/app/app.go`**: Contains the core application logic. It imports the `utils` package.
- **`internal/utils/utils.go`**: Intended for generic utility functions. However, it incorrectly tries to import the `app` package, creating a circular dependency.

## The Import Cycle

The dependency graph looks like this:
`main` -> `app` -> `utils` -> `app`

This cycle (`app` depends on `utils`, and `utils` depends on `app`) is forbidden in Go.

## How to Reproduce

Attempt to run the project from the root directory:

```bash
go run cmd/example/main.go
```

You will see an error message similar to this:

```text
package github.com/philwo/importcycle/cmd/example
	imports github.com/philwo/importcycle/internal/app from main.go
	imports github.com/philwo/importcycle/internal/utils from app.go
	imports github.com/philwo/importcycle/internal/app from utils.go: import cycle not allowed
```
