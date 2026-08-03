package internal

import (
	"errors"
	"strings"
)

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

func (r *Registry) Find(args []string) (*Command, error) {
	cmd_key := strings.Join(args, " ")
	if value, ok := r.root[cmd_key]; ok {
		return value, nil
	}
	return nil, errors.New("Command not found")
}
