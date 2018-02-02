package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

type Task struct {
	ID          int
	Description string
	Finished    bool
}

type Storage interface {
	Add(string)(int, error)
	Do(int) error
	List()
}

type FileStorage struct{
	*os.File
}

func NewFileStorage(filePath string) (*FileStorage, error) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileStorage{f}, nil
}

// Add adds new task. If succeeded, Add returns task ID with Nil. If not, Add returns error and its ID is zero
func (f *FileStorage) Add(description string) (taskID int, err error) {
	return
}

// Do mark the task specified with `taskID` as completed. If the task doesn't exist, Do return error.
func (f *FileStorage) Do(taskID int) (err error) {
	return
}

// List prints the list of TODOs unmarked as finished.
func (f *FileStorage) List() {
	return
}

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
		Name:  "do",
		Usage: `Mark a task on your TODO list as complete`,
		Action: func(c *cli.Context) error {
			fmt.Println("do")
			return nil
		},
	}
	listCommand := cli.Command{
		Name:  "list",
		Usage: `List all of your incomplete tasks`,
		Action: func(c *cli.Context) error {
			fmt.Println("list")
			return nil
		},
	}
	app.Commands = append(app.Commands, addCommand, doCommand, listCommand)

	app.Run(os.Args)
}
