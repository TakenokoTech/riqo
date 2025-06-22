# Development Guide

## Linter Setup
To install and use the linter (`golangci-lint`), follow these steps:

1. **Install golangci-lint**:
   ```bash
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   ```

2. **Run the linter**:
   ```bash
   golangci-lint run
   ```

## Formatter Setup
To format the code, use the following command:

1. **Use `gofmt`**:
   ```bash
   gofmt -w .
   ```

2. **Alternative: Use `golangci-lint` for formatting**:
   ```bash
   golangci-lint run --fix
   ```

## Running the CLI
To build and run the CLI tool, follow these steps:

1. **Build the binary**:
   ```bash
   go build -o riqo
   ```

2. **Run the CLI**:
   ```bash
   ./riqo suggest "Show repository details"
   ```

   Example output:
   ```
   Processing input: Show repository details
   Suggested GitHub CLI command: gh repo view
   ```

3. **View Command History**:
   ```bash
   ./riqo history view
   ```

   Example output:
   ```
   Command History:
   suggest Show repository details
   ```

4. **Clear Command History**:
   ```bash
   ./riqo history clear
   ```

   Example output:
   ```
   History cleared.
   ```

## Measuring Test Coverage
To measure test coverage and analyze the results, use the following commands:

1. **Generate a coverage profile**:
   ```bash
   go test ./cmd/... -coverprofile=coverage.out
   ```

2. **Analyze the coverage profile**:
   ```bash
   go tool cover -func=coverage.out
   ```

   This will display a summary of the coverage for each function in the codebase.

## Development Script
To simplify the build and run process during development, use the provided script:

1. **Run the script**:
   ```bash
   ./scripts/dev_run.sh <command> [args...]
   ```

   Example:
   ```bash
   ./scripts/dev_run.sh help
   ```

   This script will:
   - Build the Riqo binary using `go build -o riqo`.
   - Execute the specified Riqo command with any provided arguments.

## Notes
- Ensure your Go version is up-to-date to avoid compatibility issues.
- Add these commands to your workflow to maintain code quality.
