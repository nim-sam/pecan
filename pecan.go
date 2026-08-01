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
	r.Register(new_cmd)
	return new_cmd
}
