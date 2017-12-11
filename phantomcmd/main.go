package main

import (
	// standard
	// "fmt"

	// local
	"github.com/jheise/phantomlimb/phantomapi"

	// external
	"github.com/alecthomas/kingpin"
)

var (
	// plugin
	plugin = kingpin.Arg("plugin", "Which plugin to run"). Required().String()
	target = kingpin.Arg("target", "Website to target").Required().String()

	// api
	api *phantomapi.PhantomApi
)

func main() {
	kingpin.Parse()

	api = phantomapi.NewPhantomApi("localhost:9797")

	if *plugin == "screenshot" {
		processScreenShot(*target)
	}
}
