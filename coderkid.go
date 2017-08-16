package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	fmt.Println("Welcome to the Coderkid client.")
	app := cli.NewApp()
	app.Name = "Coderkid"
	app.Usage = "A client for interfacing with coderkid.com"
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
			Usage:   "Download specified almanac from coderkid.com",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "read",
			Aliases: []string{"r"},
			Usage:   "Read specified almanac from coderkid.com",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	app.Run(os.Args)
}
