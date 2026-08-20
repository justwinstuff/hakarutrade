package main

import (
	"github.com/alecthomas/kong"
)

const ver = "a1.0.0"

var CLI struct {
	Version kong.VersionFlag `short:"v" help:"Print version."`
}

func main() {
	ctx := kong.Parse(&CLI, kong.Vars{
		"version": ver,
	})
	switch ctx.Command() {
	default:
		ctx.PrintUsage(false)
	}
}