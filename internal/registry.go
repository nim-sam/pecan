package internal

type Registry struct {
	// Holds the dummy head of the command tree
	// which is upon struct creation
	root map[string]*Command
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) Register(cmd *Command) {

}
