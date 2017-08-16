package tools

import (
	"fmt"
	"os"

	app "github.com/urfave/cli"
)

// Read() reads from either the local source, or reads from the web
// for the almanac specified
func Read(c *app.Context) error {
	args := c.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Use: coderkid read <almanac_name>")
		os.Exit(1)
	}
	almanac := args[0]
	fmt.Println("Reading", almanac)
	return nil
}
