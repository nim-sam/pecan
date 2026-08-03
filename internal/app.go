package internal

import "os"

type App struct {
	Name     string
	Registry *Registry
}

func (a *App) Launch() (*Command, error) {
	cmd, err := a.Registry.Find(os.Args[1:])

	if err != nil {
		return nil, err
	}

	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return cmd, nil
}
