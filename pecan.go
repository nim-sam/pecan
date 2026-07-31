package pecan

import "github.com/nim-sam/pecan/internal"

// Runs upon package import. Will instantiate the command registry
// and takes care of user input parsing

func init() {
	r := internal.NewRegistry()
}

// Command Constructor

func NewCommand() *internal.Command {
	return &internal.Command{}
}
