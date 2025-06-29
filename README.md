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
   go build -o riqo
   ```

## Usage

Run the following command to see available options:
```bash
riqo help
```

### Example Commands

- **Suggest a command**:
  ```bash
  riqo suggest "create a repository"
  ```

- **View command history**:
  ```bash
  riqo history view
  ```

- **Clear command history**:
  ```bash
  riqo history clear
  ```

- **Update documentation**:
  ```bash
  riqo docs
  ```

- **Real-time suggestions**:
  ```bash
  riqo realtime
  ```

## License

This project is licensed under the MIT License.
