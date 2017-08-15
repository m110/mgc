package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mgc"
	app.Usage = "the mongodb console"
	app.Action = cli.ShowAppHelp

	app.Commands = []cli.Command{
		{
			Name:    "find",
			Aliases: []string{"f"},
			Usage:   "find object in collection",
			Action:  findCommand,
		},
	}

	instance, err := readline.NewEx(&readline.Config{
		Prompt:          color.GreenString("» "),
		HistoryFile:     "/tmp/history.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})

	if err != nil {
		fmt.Println("Failed to initialize readline: ", err)
		os.Exit(1)
	}
	defer instance.Close()

	loop(app, instance)
}

func loop(app *cli.App, instance *readline.Instance) int {
	for {
		line, err := instance.Readline()
		if err == readline.ErrInterrupt {
			continue
		} else if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading input: ", err)
			continue
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		args := append([]string{"mgc"}, strings.Fields(line)...)

		if err := app.Run(args); err != nil {
			fmt.Println("App error:", err)
		}
	}

	return 0
}

func findCommand(ctx *cli.Context) error {
	fmt.Println("Find:", ctx.Args().Get(0), ctx.Args().Get(1))
	return nil
}
