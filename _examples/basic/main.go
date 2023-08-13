package main

import (
	"fmt"
	"gargs"
	"os"
)

type options struct {
	Color string `short:"c" desc:"Color of the item" defaultValue:"blue"`
	Count int    `short:"n" desc:"Number of items" defaultValue:"1"`
}

func main() {
	opts := &options{}
	err, _ := gargs.Handle(os.Args, opts, "<item>")
	if err != nil {
		return
	}

	fmt.Printf("Color of %s: %s\n", gargs.GetParam("item"), opts.Color)
	fmt.Printf("Number of %s: %d\n", gargs.GetParam("item"), opts.Count)
}
