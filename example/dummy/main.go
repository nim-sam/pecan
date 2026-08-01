package main

import "github.com/nim-sam/pecan"

func main() {

	c := pecan.NewCommand().WithName("FirstArg")

	d := pecan.NewCommand().WithName("SecondArg").WithParent(c)

	e := pecan.NewCommand().WithName("ThirdArg").WithParent(d)

	path := e.Path()

	println(path)
}
