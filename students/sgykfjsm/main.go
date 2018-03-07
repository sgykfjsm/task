package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gophercises/task/students/sgykfjsm/storage"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "task"
	app.Usage = "task is a CLI for managing your TODOs."
	app.Version = "0.1.0"

	filePath := "task.db"
	bucketName := "task"
	db, err := storage.NewBoltDBStorage(filePath, bucketName)
	if err != nil {
		log.Fatal(err)
	}

	app.Commands = []cli.Command{}
	addCommand := cli.Command{
		Name:  "add",
		Usage: `Add a new task to your TODO list`,
		Action: func(c *cli.Context) error {
			args := c.Args()
			if len(args) == 0 {
				log.Fatal("ERROR: Please give some description to add task")
			}
			description := strings.Join(args, " ")
			task, err := db.Add(description)
			if err != nil {
				return err
			}
			fmt.Printf("Added %q to your task list\n", task.Description)

			return nil
		},
	}

	doCommand := cli.Command{
		Name:  "do",
		Usage: `Mark a task on your TODO list as complete`,
		Action: func(c *cli.Context) error {
			args := c.Args()
			if len(args) == 0 {
				log.Fatal("ERROR: Please give task id")
			}

			i, err := strconv.Atoi(args.Get(0))
			if err != nil {
				log.Fatal(err)
			}

			task, err := db.Find(i)
			if err != nil {
				log.Fatal(err)
			}

			task.Finished = true
			if err := db.Put(task); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("You have completed the %q task.\n", task.Description)
			return nil
		},
	}

	listCommand := cli.Command{
		Name:  "list",
		Usage: `List all of your incomplete tasks`,
		Action: func(c *cli.Context) error {
			tasks, err := db.FindAll()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("You have the following tasks:")
			for i, task := range tasks {
				if !task.Finished {
					fmt.Printf("%d. %s\n", i+1, task.Description)
				}
			}
			return nil
		},
	}

	app.Commands = append(app.Commands, addCommand, doCommand, listCommand)

	app.Run(os.Args)
}
