package main

import (
	"fmt"
	"gargs"
	"os"
)

type options struct {
	Color string `short:"c" desc:"Color of the output" defaultValue:"red"`
	Delay int    `short:"d" desc:"Delay in seconds" defaultValue:"5"`
}

func main() {
	opts := &options{}

	// Custom MOTD
	gargs.SetMOTD(`
Welcome to gargs!
Your lightweight solution for command-line argument parsing.
`)

	_, helpTriggered := gargs.Handle(os.Args, opts)
	if helpTriggered {
		return
	}

	fmt.Println("Color:", opts.Color)
	fmt.Println("Delay:", opts.Delay)
}
