package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"

	"github.com/alecthomas/kong"
)

type CLI struct {
	Init InitCmd `cmd:"" help:"Initialise a Golang module and Git in the current directory."`
	New  NewCmd  `cmd:"" help:"Create a subdirectory and initialise it."`
}

type InitCmd struct {
	ModuleName string `arg:"" help:"The module name (passed to go mod init)."`

	NoGit bool `help:"Don't initialise Git." short:"g"`
	Bin   bool `help:"Create a main.go file." short:"b"`
}

func (c *InitCmd) Run(ctx *kong.Context) error {
	// Initalise Git
	if c.NoGit {
		fmt.Fprintln(os.Stderr, "not initialising Git")
	} else {
		git := exec.Command("git", "init")

		out, err := git.CombinedOutput()
		if err != nil {
			fmt.Fprint(os.Stderr, "error initialising Git:", string(out))
			return err
		}
		fmt.Fprint(os.Stderr, string(out))
	}

	// Create a Go module
	goMod := exec.Command("go", "mod", "init", c.ModuleName)
	out, err := goMod.CombinedOutput()
	if err != nil {
		fmt.Fprint(os.Stderr, "error initialising a Go module:", string(out))
		return err
	}
	fmt.Fprint(os.Stderr, string(out))

	if !c.Bin {
		return nil
	}

	helloWorld := []byte(`package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}`)

	// Create a main.go file and add some 'hello world' code
	err = os.WriteFile("main.go", helloWorld, 0644)
	if err != nil {
		fmt.Fprint(os.Stderr, "error creating or writing to main.go:", string(out))
		return err
	}
	fmt.Fprint(os.Stderr, "created main.go")

	return nil
}

type NewCmd struct {
	DirName    string `arg:"" help:""`
	ModuleName string `arg:"" help:"The module name (passed to go mod init)."`

	NoGit bool `help:"Don't initialise Git." short:"g"`
	Bin   bool `help:"Create a main.go file." short:"b"`
}

func (c *NewCmd) Run(ctx *kong.Context) error {
	// Make the target directory
	err := os.Mkdir(c.DirName, 0755)
	if err != nil && errors.Is(err, fs.ErrExist) {
		fmt.Fprintln(os.Stderr, "error creating target directory:", err)
		return err
	}

	// Move into that directory
	err = os.Chdir(c.DirName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error changing directory:", err)
		return err
	}

	// Run the `init` command
	ic := InitCmd{
		ModuleName: c.ModuleName,
		NoGit:      c.NoGit,
		Bin:        c.Bin,
	}

	return ic.Run(ctx)
}

func main() {
	var cli CLI

	ctx := kong.Parse(&cli,
		kong.Name("go-init"),
		kong.Description("A simple CLI that creates or initialises a Golang module and Git with one command."),
		kong.UsageOnError(),
	)

	err := ctx.Run()
	if err != nil {
		os.Exit(1)
	}
}
