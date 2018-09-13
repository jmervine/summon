package command

import (
	"github.com/codegangsta/cli"
)

var Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "p, provider",
		Usage: "Path to provider for fetching secrets",
	},
	cli.StringFlag{
		Name:  "e, environment",
		Usage: "Specify section/environment to parse from secrets.yaml",
	},
	cli.StringFlag{
		Name:  "f",
		Value: "secrets.yml",
		Usage: "Path to secrets.yml",
	},
	cli.StringSliceFlag{
		Name:  "D",
		Value: &cli.StringSlice{},
		Usage: "var=value causes substitution of value to $var",
	},
	cli.StringFlag{
		Name:  "yaml",
		Usage: "secrets.yml as a literal string",
	},
	cli.StringFlag{
		Name:  "output, o",
		Usage: "write results out to a file in 'env' format and exit without executing a command",
	},
	cli.StringSliceFlag{
		Name:  "ignore, i",
		Value: &cli.StringSlice{},
		Usage: "Ignore the specified key if is isn't accessible or doesn't exist",
	},
	cli.BoolFlag{
		Name:  "ignore-all, I",
		Usage: "Ignore inaccessible or missing keys",
	},
}
