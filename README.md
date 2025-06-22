# Riqo

<center>
<img height="128px" src="docs/logo.png">
</center>

Riqo is a command-line interface (CLI) tool designed to enhance productivity by providing intelligent suggestions and managing CLI commands efficiently.

## Features

- **Suggest**: Generate CLI commands based on natural language input.
- **History**: Manage command history (view, clear, and search).
- **Docs**: Automatically organize and update documentation based on command usage.
- **Realtime**: Provide real-time command suggestions.
- **LLM Integration**: Leverage local LLMs for intelligent command suggestions.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/riqo.git
   cd riqo
   ```

2. Build the binary:
   ```bash
   ./scripts/dev_run.sh build
   ```

## Usage

Run the following command to see available options:
```bash
./scripts/dev_run.sh help
```

### Example Commands

- **Suggest a command**:
  ```bash
  ./scripts/dev_run.sh suggest "create a repository"
  ```

- **View command history**:
  ```bash
  ./scripts/dev_run.sh history view
  ```

- **Clear command history**:
  ```bash
  ./scripts/dev_run.sh history clear
  ```

- **Update documentation**:
  ```bash
  ./scripts/dev_run.sh docs
  ```

- **Real-time suggestions**:
  ```bash
  ./scripts/dev_run.sh realtime
  ```

## License

This project is licensed under the MIT License.
