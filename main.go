package main

import (
	"fmt"
	"os"

	"github.com/ezetter/task/cmd"
	"github.com/ezetter/task/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	homeDir, _ := homedir.Dir()
	path := fmt.Sprintf("%s/.gotasks", homeDir)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
	db.Init(path)
	defer db.Close()
	cmd.Execute()
}
