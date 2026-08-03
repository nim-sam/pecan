package pecan

import "github.com/nim-sam/pecan/internal"

var r *internal.Registry

// Runs upon package import. Will instantiate the command registry
// and takes care of user input parsing

func init() {
	r = internal.NewRegistry()
	_ = r
}

// Command Constructor

func NewCommand() *internal.Command {
	new_cmd := &internal.Command{}
	new_cmd.WithRegistry(r)
	return new_cmd
}

// Application Constructor

func App(name string) *internal.App {
	a := &internal.App{
		Name:     name,
		Registry: r,
	}
	return a
}
