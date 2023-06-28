# Factorial CLI

[![Go Build](https://github.com/wantaekchoi/factorial-go/actions/workflows/go.yml/badge.svg)](https://github.com/wantaekchoi/factorial-go/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/wantaekchoi/factorial-go)](https://goreportcard.com/report/github.com/wantaekchoi/factorial-go)

This is a command-line interface (CLI) tool written in Go that calculates the factorial of a given integer.

## Usage

To use the factorial CLI, follow the instructions below:

1. Clone the repository:

   ```
   git clone https://github.com/your-username/factorial-cli.git
   ```

1. Navigate to the project directory:

   ```
   cd factorial-cli
   ```

1. Build the CLI tool:

   ```
   go build factorial.go
   ```

1. Run the CLI tool with the desired number and options:

   ```
   ./factorial-go -number <number> -base <base>
   ```

Replace `<number>` with the integer value for which you want to calculate the factorial. Replace `<base>` with the desired number base for the input (default is 10).

Examples:

- Calculate the factorial of 5 (base 10):

  ```
  ./factorial-go -number 5
  ```

- Calculate the factorial of FF (base 16):
  ```
  ./factorial-go -number FF -base 16
  ```

5. The factorial of the given number will be displayed.

## Options

The following options are available:

- `-number`, `-n`: Integer value to calculate factorial (required).
- `-base`, `-b`: Number base for the input (default: 10).
