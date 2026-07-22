package main

import (
	"fmt"

	"github.com/nim-sam/pecan/pkg/command"
)

func main() {

	cmd_one := command.NewCommand().
		WithName("Parent Command").
		WithDesc("Parent Description")

	cmd_two := command.NewCommand().
		WithName("Test command").
		WithDesc("Description").
		WithUsage("Argument in the <cmd>").
		WithParent(cmd_one)

	fmt.Println(cmd_two.Parent())

}
