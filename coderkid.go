package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"

	"github.com/MaxTheMonster/coderkid/config"
	"github.com/MaxTheMonster/coderkid/tools"
)

const (
	Version = "0.1.0"

	descDownload = "Download specified almanac from coderkid.com"
	descRead     = "Read specified almanac from coderkid.com"
	descTest     = "Finds current task and runs according tests."
	descConfig   = "Generates a config based on user inputted values."
)

func main() {
	fmt.Println("Welcome to the Coderkid client.")
	app := cli.NewApp()
	app.Name = "Coderkid"
	app.Usage = "A client for interfacing with coderkid.com"
	app.Version = Version
	// app.Flags = []cli.Flag {
	// 	cli.StringFlag{
	// 		Name: "almanac",
	// 		Usage: "Almanac relating to the almanacs on coderkid.com",

	// 	},
	// }
	app.Commands = []cli.Command{
		{
			Name:    "download",
			Aliases: []string{"dl"},
			Usage:   descDownload,
			Action:  tools.Download,
		},
		{
			Name:    "read",
			Aliases: []string{"r"},
			Usage:   descRead,
			Action:  tools.Read,
		},
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   descTest,
			Action:  tools.Test,
		},
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   descConfig,
			Action:  config.Configure,
		},
	}

	app.Run(os.Args)
}
