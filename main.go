package main

import (
	"github.com/atotto/clipboard"
	"github.com/jasonharrison/usps/usps"
	"gopkg.in/urfave/cli.v1"
	"os"
	"sort"
)

// Formats a USPS tracking number in the clipboard.
func formatcmd(c *cli.Context) error {
	text, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}
	clipboard.WriteAll(usps.Format(text))
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "USPS Tools"
	app.Usage = "Provides functions for working with USPS tracking numbers."
	app.Version = "git"
	app.Commands = []cli.Command{
		{
			Name:    "format",
			Aliases: []string{"f"},
			Usage:   "Formats a USPS tracking number in the clipboard.",
			Action:  formatcmd,
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}
