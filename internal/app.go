package internal

import "os"

type App struct {
	Name     string
	Registry *Registry
}

func (a *App) Launch() error {
	cmd := a.Registry.Find(os.Args[1:])
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
