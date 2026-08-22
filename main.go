package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/justwinstuff/hakarutrade/cmd"
)

const ver = "a1.0.0"

var CLI struct {
	Version kong.VersionFlag `short:"v" help:"Print version"`
	Setup   struct{}         `cmd:"" help:"Setup HakaruTrade"`
	Serve   struct{}         `cmd:"" aliases:"s,start" help:"Start HakaruTrade" default:"1"`
}

func main() {
	ctx := kong.Parse(&CLI, kong.Vars{
		"version": ver,
	})
	switch ctx.Command() {
	case "setup":
		cmd.Setup()
	case "serve":
		fmt.Println("Serve")
	default:
		ctx.PrintUsage(false)
	}
}
