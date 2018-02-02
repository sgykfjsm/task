package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "task"
	app.Usage = "task is a CLI for managing your TODOs."
	app.Version = "0.1.0"

	app.Commands = []cli.Command{}
	addCommand := cli.Command{
		Name:  "add",
		Usage: `Add a new task to your TODO list`,
		Action: func(c *cli.Context) error {
			fmt.Println("Add")
			return nil
		},
	}
	doCommand := cli.Command{
		Name:    "do",
		Usage:   `Mark a task on your TODO list as complete`,
		Action: func(c *cli.Context) error {
			fmt.Println("do")
			return nil
		},
	}
	listCommand := cli.Command{
		Name:    "list",
		Usage:   `List all of your incomplete tasks`,
		Action: func(c *cli.Context) error {
			fmt.Println("list")
			return nil
		},
	}
	app.Commands = append(app.Commands, addCommand, doCommand, listCommand)

	app.Run(os.Args)
}
