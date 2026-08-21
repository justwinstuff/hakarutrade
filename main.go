package main

import (
	"github.com/alecthomas/kong"
	"github.com/justwinstuff/hakarutrade/cmd"
)

const ver = "a1.0.0"

var CLI struct {
	Version kong.VersionFlag `short:"v" help:"Print version."`
	Setup struct{} `cmd:"" help:"Setup HakaruTrade"`
}

func main() {
	ctx := kong.Parse(&CLI, kong.Vars{
		"version": ver,
	})
	switch ctx.Command() {
	case "setup":
		cmd.DownloadMT5()
		cmd.CreateConfig()
	default:
		ctx.PrintUsage(false)
	}
}
