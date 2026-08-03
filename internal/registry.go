package internal

import "strings"

type Registry struct {
	// Holds the dummy head of the command tree
	// which is upon struct creation
	root map[string]*Command
}

func NewRegistry() *Registry {
	return &Registry{root: make(map[string]*Command)}
}

func (r *Registry) Register(cmd *Command) {
	r.root[cmd.Path()] = cmd
}

func (r *Registry) Find(args []string) *Command {
	cmd_key := strings.Join(args, " ")
	return r.root[cmd_key]
}
