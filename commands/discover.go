package commands

import (
	"fmt"
	"os"

	"github.com/github/hub/github"
	"github.com/github/hub/utils"
)

var cmdDiscover = &Command{
	Run:   discover,
	Usage: "discover USERNAME",
	Short: "List all the public GitHub repos of the given user. ",
	Long:  "Fetch and show the list of all public GitHub repos of the given user.",
}

func init() {
	CmdRunner.Use(cmdDiscover)
}

func discover(cmd *Command, args *Args) {
	if args.ParamsSize() < 1 {
		utils.Check(fmt.Errorf("Error: you must specify a username."))
	}

	client := github.NewClient(github.GitHubHost)
	repos, err := client.Repositories(args.GetParam(0))
	if err != nil {
		utils.Check(err)
	}

	for _, repo := range repos {
		fmt.Println(repo.Name)
	}

	os.Exit(0)
}
