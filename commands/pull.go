package commands

import (
	"errors"
	"fmt"
	"github.com/MichaelMure/git-bug/bug"
	"github.com/MichaelMure/git-bug/repository"
)

func runPull(repo repository.Repo, args []string) error {
	if len(args) > 1 {
		return errors.New("Only pulling from one remote at a time is supported")
	}

	remote := "origin"
	if len(args) == 1 {
		remote = args[0]
	}

	if err := repo.PullRefs(remote, bug.BugsRefPattern+"*", bug.BugsRemoteRefPattern+"*"); err != nil {
		return err
	}
	return nil
}

// showCmd defines the "push" subcommand.
var pullCmd = &Command{
	Usage: func(arg0 string) {
		fmt.Printf("Usage: %s pull [<remote>]\n", arg0)
	},
	RunMethod: runPull,
}
