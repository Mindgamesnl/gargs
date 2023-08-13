package gargs

import (
	"testing"
)

func TestHandle(t *testing.T) {
	// Define some options for testing
	type options struct {
		Color string `short:"c" desc:"Color of the output" defaultValue:"red"`
		Delay int    `short:"d" desc:"Delay in seconds" defaultValue:"5"`
	}

	opts := &options{}

	// Test valid arguments
	err, helpTriggered := Handle([]string{"binary", "-c", "blue", "-d", "10", "arg1", "arg2"}, opts, "<arg1>", "<arg2>")
	if err != nil {
		t.Errorf("Handle() error = %v", err)
		return
	}
	if opts.Color != "blue" {
		t.Errorf("expected Color: blue, got: %s", opts.Color)
	}
	if opts.Delay != 10 {
		t.Errorf("expected Delay: 10, got: %d", opts.Delay)
	}
	if len(Params()) != 2 || GetParam("arg1") == "" || GetParam("arg2") == "" {
		t.Errorf("Positional parameters not parsed correctly")
	}

	// Test when help is triggered
	_, helpTriggered = Handle([]string{"binary", "-h"}, opts)
	if !helpTriggered {
		t.Errorf("expected helpTriggered: true, got: false")
	}

	// Test positional arguments
	err, _ = Handle([]string{"binary", "value1", "value2"}, opts, "<param1>", "<param2>")
	if err != nil {
		t.Errorf("Handle() error = %v", err)
		return
	}
	if GetParam("param1") != "value1" {
		t.Errorf("expected param1: value1, got: %s", GetParam("param1"))
	}
	if GetParam("param2") != "value2" {
		t.Errorf("expected param2: value2, got: %s", GetParam("param2"))
	}

	// Test default values
	err, _ = Handle([]string{"binary"}, opts)
	if err != nil {
		t.Errorf("Handle() error = %v", err)
		return
	}
	if opts.Color != "red" {
		t.Errorf("expected default Color: red, got: %s", opts.Color)
	}
	if opts.Delay != 5 {
		t.Errorf("expected default Delay: 5, got: %d", opts.Delay)
	}

	// Test shorthand flags
	err, _ = Handle([]string{"binary", "-c", "green", "-d", "3"}, opts)
	if err != nil {
		t.Errorf("Handle() error = %v", err)
		return
	}
	if opts.Color != "green" {
		t.Errorf("expected Color using shorthand: green, got: %s", opts.Color)
	}
	if opts.Delay != 3 {
		t.Errorf("expected Delay using shorthand: 3, got: %d", opts.Delay)
	}

	// Test missing required positional parameters
	err, _ = Handle([]string{"binary", "value1"}, opts, "<param1>", "<param2>")
	if err == nil {
		t.Errorf("Expected an error for missing positional parameter but got none")
	}
}
