package main

import (
	"log"
	"path/filepath"

	"github.com/CodeGophercises/cli-task-manager/cmd"
	"github.com/CodeGophercises/cli-task-manager/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	dir, _ := homedir.Dir()
	userHome := filepath.Join(dir, "tasks.db")
	checkError(db.InitPath(userHome))
	checkError(cmd.Execute())
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Something is wrong: %s", err)
	}
}
