# gargs ![gargs](https://github.com/Mindgamesnl/gargs/actions/workflows/go.yml/badge.svg)

![image](https://github.com/Mindgamesnl/gargs/assets/10709682/76c00c3c-7417-4a0d-bc72-6913769bfe40)

**gargs** is a lightweight, flexible Go package for command-line argument parsing. It provides an intuitive API to define and access command line options with built-in support for custom MOTDs (Message of The Day).

# Features
- **ORM like**: Automatically parse and map your arguments to a model. No more endless switches or elfsifs in your application startup
- **Default values**: Set default values for your arguments
- **Simple API**: Easily define and parse command-line arguments with automatic parsing an dhelp menus
- **Built-in MOTD**: Display a custom message when the application starts.
- **Support for Positional Parameters**: Beyond just flags, gargs supports parsing positional arguments as well.
- **Type Support**: Parse various data types including strings, integers, and booleans.

# Installation
```bash
go get github.com/Mindgamesnl/gargs
```

# Quick Start
Here's a basic example to showcase how you can use `gargs`:
```go
package main

import (
    "fmt"
    "github.com/Mindgamesnl/gargs"
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
```

This will produce the following output:
```bash
$ ./gargs-demo -h
Welcome to gargs!
Your lightweight solution for command-line argument parsing.

Usage: gargs-demo [options..] 

  -c string
        Color of the output (shorthand) (default "red")
  -color string
        Color of the output (default "red")
  -d int
        Delay in seconds (shorthand) (default 5)
  -delay int
        Delay in seconds (default 5)
```

or when you run the application with the following arguments:
```bash
$ ./gargs-demo -c blue -d 10
Color: blue
Delay: 10
```

Or default values from the struct:
```bash
$ ./gargs-demo
Color: red
Delay: 5
```
