package main

import (
	"fmt"
	"gargs"
	"os"
)

type options struct {
	Verbose bool `short:"v" desc:"Display verbose output" defaultValue:"false"`
}

func main() {
	opts := &options{}
	err, _ := gargs.Handle(os.Args, opts, "[file]")
	if err != nil {
		return
	}

	if opts.Verbose {
		fmt.Println("Verbose mode is ON.")
	}

	if filename := gargs.GetParam("file"); filename != "" {
		fmt.Printf("Processing file: %s\n", filename)
	} else {
		fmt.Println("No file specified.")
	}
}
