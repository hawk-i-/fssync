package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/hawk-i-/fssync/core"
)

//VERSION Application version to be injected during build time
var VERSION = "0.0.1"

var (
	author     = "Gurbakhshish Singh"
	configFile = kingpin.Flag("config", "Config file").Short('c').Default("config.yaml").String()
)

func main() {
	kingpin.Version(VERSION)
	kingpin.CommandLine.Author(author)
	kingpin.Parse()

	config, err := core.NewConfig(*configFile)

	kingpin.FatalIfError(err, "Unable to load config")

	fmt.Println(config.Version)
}
