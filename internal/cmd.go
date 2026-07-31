package internal

import "os"

// Handler function type
type HandlerFunction func([]string) error

// Command struct
type Command struct {
	// Command metadata
	name  string
	desc  string
	usage string

	// A command can have a parent or
	// mulitple children
	parent   *Command
	children []*Command

	// Handler method (method ran by the command)
	handler HandlerFunction
}

// Command Getters

func (cmd *Command) Name() string {
	return cmd.name
}

func (cmd *Command) Desc() string {
	return cmd.desc
}

func (cmd *Command) Usage() string {
	return cmd.usage
}

func (cmd *Command) Parent() *Command {
	return cmd.parent
}

func (cmd *Command) Children() []*Command {
	return cmd.children
}

func (cmd *Command) Handler() HandlerFunction {
	return cmd.handler
}

func (cmd *Command) Path() []string {
	temp := cmd
	var cmd_names []string

	for temp != nil {
		cmd_names = append(cmd_names, temp.Name())
		temp = temp.Parent()
	}

	// Reversing order of command parent traversal to obtain
	// full command path
	var out []string
	for i := len(cmd_names); i >= 0; i-- {
		out = append(out, cmd_names[i])
	}

	return out
}

// Command Setters

func (cmd *Command) WithName(name string) *Command {
	cmd.name = name
	return cmd
}

func (cmd *Command) WithDesc(desc string) *Command {
	cmd.desc = desc
	return cmd
}

func (cmd *Command) WithUsage(usage string) *Command {
	cmd.usage = usage
	return cmd
}

func (cmd *Command) WithParent(parent *Command) *Command {
	cmd.parent = parent
	// Automatically adding current command to children of parent
	parent.children = append(parent.children, cmd)
	return cmd
}

func (cmd *Command) WithHandler(handler HandlerFunction) *Command {
	cmd.handler = handler
	return cmd
}

// Command execution methods

func (cmd *Command) Run() (*Command, error) {
	if err := cmd.handler(os.Args); err != nil {
		return nil, err
	}
	return cmd, nil
}
