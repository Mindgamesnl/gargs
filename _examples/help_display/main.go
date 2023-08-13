package main

import (
	"gargs"
	"os"
)

type options struct {
	Output string `short:"o" desc:"Output file name" defaultValue:"output.txt"`
}

func main() {
	opts := &options{}
	_, helpTriggered := gargs.Handle(os.Args, opts)
	if helpTriggered {
		return
	}

	// rest of your code...
}
