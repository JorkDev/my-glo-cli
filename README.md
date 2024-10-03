
# My Go CLI

[![Go Version](https://img.shields.io/github/go-mod/go-version/JorkDev/my-go-cli)](https://golang.org/doc/go1.18)
[![License](https://img.shields.io/github/license/JorkDev/my-go-cli)](LICENSE)

## Description

`My Go CLI` is a command-line tool built in Go that showcases advanced development techniques such as configuration management, logging, error handling, Dockerization, and unit testing. This project is built step-by-step, with each commit demonstrating a new feature or concept. The goal is to provide a comprehensive guide to building an advanced CLI in Go from scratch, progressing from simple to complex.

---

## Features

- Basic command handling using Cobra
- Configuration management with Viper
- Structured logging with Zap
- Advanced error handling
- Unit testing with Go's testing package
- Dockerized CLI for easy distribution

---

## Installation

To get started with `My Go CLI`, clone this repository and ensure you have Go installed.

```bash
git clone https://github.com/JorkDev/my-go-cli.git
cd my-go-cli
go mod tidy
```

---

## Usage

Run the CLI directly from the terminal:

```bash
go run main.go
```

### Available Commands

- **`mycli echo [text]`**: Echoes the input text back to the terminal.
- **`mycli version`**: Prints the current version of the CLI.

---

## Project Structure

```
my-go-cli/
├── cmd/                  # Command definitions
│   ├── root.go
├── config/               # Configuration files and setup
│   └── config.go
├── internal/             # Internal packages, such as logging
│   └── logger.go
├── config.yaml           # Default configuration file
├── main.go               # Main entry point of the CLI
├── README.md             # Project description and instructions
├── go.mod                # Go module file
└── LICENSE               # License file
```

---

## Contributing

If you'd like to contribute to `My Go CLI`, feel free to fork the repository, make some changes, and submit a pull request. All contributions are welcome!

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

## Roadmap

This project is divided into several phases. Each phase adds more complexity and advanced features:

1. **Basic CLI Skeleton**: Initialize a basic Cobra-based CLI.
2. **Command Handling**: Add commands and flags with Cobra.
3. **Configuration Management**: Integrate Viper for managing configuration files.
4. **Logging**: Implement structured logging using Zap.
5. **Error Handling**: Add comprehensive error handling.
6. **Unit Testing**: Add unit tests for commands.
7. **Dockerization**: Package the CLI with Docker for easy deployment.

---

## Contact

For any questions or support, please reach out via GitHub Issues.
