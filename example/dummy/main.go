package main

import "github.com/nim-sam/pecan"

func main() {
	println("test")

	c := pecan.NewCommand().WithName("your mom")

	d := pecan.NewCommand().WithName("your dad").WithParent(c)

	e := pecan.NewCommand().WithName("your sister").WithParent(d)

	path := pecan.ExtractCmdPath(e)
	for i := 0; i < 3; i++ {
		println(path[i])
	}
}
