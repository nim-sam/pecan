package main

import "github.com/nim-sam/pecan"

func test_func(args []string) error {
	println("Hello, world!")
	return nil
}

func main() {

	c := pecan.NewCommand().WithName("FirstArg").Register()

	d := pecan.NewCommand().WithName("SecondArg").WithParent(c).Register()

	pecan.NewCommand().WithName("ThirdArg").WithParent(d).WithHandler(test_func).Register()

	// path := e.Path()

	pecan.App("Test application").Launch()
}
