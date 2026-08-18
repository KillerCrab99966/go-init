# Go Init

A simple CLI written in Go with Kong that creates or initialises a Golang module and Git with one command.

## Installation

### With Go CLI and Github repo (Recommended)

Run `go install github.com/KillerCrab99966/go-init` to install.

To run the tool from anywhere, add `$(go env GOPATH)/bin` to your `PATH` environment variable.

### From local repo
1. Clone this repo
2. Install Go
3. Install with `go install /path/to/local/repo`
4. Add `$(go env GOPATH)/bin` to your `PATH` environment variable to run the tool from anywhere

## Usage

```
go-init <command>

Flags:
  -h, --help    Show context-sensitive help.

Commands:
  init <module-name> [flags]
    Initialise a Golang module and Git in the current directory.

  new <dir-name> <module-name> [flags]
    Create a subdirectory and initialise it.
```

Run `go-init <command> --help` for more information on a specific command.
