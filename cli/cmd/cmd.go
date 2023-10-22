package cmd

import (
	"flag"
	"os"

	"github.com/Diegiwg/dwork-web/cli/log"
	"github.com/Diegiwg/dwork-web/cli/menu"
)

var initCmd = flag.NewFlagSet("init", flag.ExitOnError)
var initProjectName = initCmd.String("name", "", "Project Name")

func InitImpl(args *[]string) {
	initCmd.Parse(*args)

	log.Info("Running init...")

	_, err := os.Stat(".project")
	if err == nil {
		log.Fatal("The working directory already has a Project in progress.")
	}

	// folders: api, page
	err = os.MkdirAll("api", 0755)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = os.MkdirAll("page", 0755)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Project Name
	if *initProjectName == "" {
		*initProjectName = menu.Question("Project Name: ")
	}

	// Save .project
	file, err := os.OpenFile(".project", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	str := "name=" + *initProjectName
	if _, err := file.WriteString(str); err != nil {
		log.Fatal(err.Error())
	}

	log.Success("Created a new Project in the current directory.")
}

func Commander() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: dwork <command>")
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "init":
		InitImpl(&args)

	default:
		log.Fatal("Unknown command: " + cmd)
	}

}
