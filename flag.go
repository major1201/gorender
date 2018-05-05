package main

import (
	"github.com/major1201/goutils"
	"github.com/urfave/cli"
)

func getApp() *cli.App {
	app := cli.NewApp()
	app.Name = "gorender"
	app.HelpName = app.Name
	app.Usage = "go template cli client"
	app.ArgsUsage = "[template file]"
	app.Version = AppVer
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "help, h",
			Usage: "show help",
		},
		cli.VersionFlag,
		cli.StringSliceFlag{
			Name:  "arguments, a",
			Usage: "set extra arguments to be rendered to the template, overrides all json/yaml/toml file arguments",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "write output to a file instead of to stdout",
		},
		cli.BoolFlag{
			Name:  "in-place, i",
			Usage: "write output in place of the template file",
		},
		cli.StringFlag{
			Name:  "json",
			Usage: "use a json argument file",
		},
		cli.StringFlag{
			Name:  "yaml, yml",
			Usage: "use a yaml argument file",
		},
		cli.StringFlag{
			Name:  "toml",
			Usage: "use a toml argument file",
		},
		cli.BoolFlag{
			Name:  "html",
			Usage: "enable html template engine which automatically secures HTML output against certain attacks",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.Bool("help") {
			cli.ShowAppHelpAndExit(c, 0)
		}
		verifyFlags(c)
		runApp(c)
		return nil
	}
	app.HideHelp = true
	return app
}

func verifyFlags(c *cli.Context) {
	// 1. only one argument should be sent in
	if c.NArg() != 1 || !goutils.IsFile(c.Args().First()) {
		cli.ShowAppHelpAndExit(c, 1)
	}
	// 2. -o can't be used with -i
	if c.Bool("in-place") && c.IsSet("output") {
		logger.Error(`"-i" can't be used with "-o"`)
		cli.ShowAppHelpAndExit(c, 1)
	}
}
