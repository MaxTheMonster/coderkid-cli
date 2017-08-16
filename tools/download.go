package tools

import (
	"fmt"
	"os"

	app "github.com/urfave/cli"
)

// Downlaod downloads the entire almanac specified from coderkid.com
func Download(c *app.Context) error {
	args := c.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Use: coderkid download <almanac_name>")
		os.Exit(1)
	}
	almanac := args[0]
	fmt.Println("Downloading", almanac)
	return nil
}
