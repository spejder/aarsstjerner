package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/mattn/go-isatty"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

var (
	//go:embed LICENSE.md
	license string
	// Version is the version string to be set at compile time via command line.
	version string
)

//nolint:funlen
func main() {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		log.SetFlags(0)
	}

	app := cli.NewApp()
	app.Name = "aarsstjerner"
	app.Usage = "Lav liste over årsstjerner"
	app.EnableBashCompletion = true
	app.Authors = []*cli.Author{
		{
			Name:  "Arne Jørgensen",
			Email: "arne@arnested.dk",
		},
	}
	app.Version = getVersion()
	app.Copyright = fmt.Sprintf("MIT License, run `%s license` to view", app.Name)

	configPath := ""
	userConfigDir, err := os.UserConfigDir()

	if err == nil {
		configPath = userConfigDir + "/aarsstjerner.yaml"
	}

	app.Flags = []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "username",
			Value:   "",
			Usage:   "The Medlemsservice username",
			EnvVars: []string{"MS_USERNAME"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "1pass",
			Value:   "",
			Usage:   "A 1Password secret reference for the Medlemsservice password",
			EnvVars: []string{"MS_1PASS"},
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:    "slack",
			Value:   90, //nolint:gomnd
			Usage:   "Days of slack in the calculation",
			EnvVars: []string{"AARSSTJERNER_SLACK"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "ms-url",
			Value:   "https://medlem.dds.dk",
			Usage:   "The Medlemsservice URL",
			EnvVars: []string{"MS_URL"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:    "ms-database",
			Value:   "dds",
			Usage:   "The Medlemsservice database name",
			EnvVars: []string{"MS_DATABASE"},
		}),
		&cli.StringFlag{
			Name:      "config",
			Value:     configPath,
			Usage:     "Read config from `FILE`",
			EnvVars:   []string{"AARSSTJERNER_CONFIG_FILE"},
			TakesFile: true,
		},
		&cli.BoolFlag{
			Name:    "all",
			Value:   false,
			Usage:   "Include all not just those who needs new",
			EnvVars: []string{"AARSSTJERNER_ALL"},
		},
		&cli.BoolFlag{
			Name:    "fake-names",
			Value:   false,
			Usage:   "Display fake names (for demo purposes)",
			EnvVars: []string{"AARSSTJERNER_FAKE_NAMES"},
			Hidden:  true,
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:   "browser",
			Usage:  "Open årsstjerner in browser (default)",
			Action: runBrowser,
		},
		{
			Name:   "markdown",
			Usage:  "List årsstjerner as markdown",
			Action: list,
		},
		{
			Name:   "term",
			Usage:  "List årsstjerner in terminal",
			Action: term,
		},
		{
			Name:  "license",
			Usage: "View the license",
			Action: func(c *cli.Context) error {
				fmt.Fprintln(os.Stdout, license)

				return nil
			},
		},
		{
			Name:   "edit-config",
			Usage:  "Open config file in editor",
			Action: editConfig,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "editor",
					Value:   "vi",
					Usage:   "Use `EDITOR` to edit config file (create it of it doesn't exist)",
					EnvVars: []string{"EDITOR"},
				},
			},
		},
	}

	app.DefaultCommand = "browser"

	app.Before = func(ctx *cli.Context) error {
		if _, err := os.Stat(configPath); !os.IsNotExist(err) {
			initConfig := altsrc.InitInputSourceWithContext(app.Flags, altsrc.NewYamlSourceFromFlagFunc("config"))

			err = initConfig(ctx)
			if err != nil {
				return fmt.Errorf("reading config file: %w", err)
			}
		}

		return nil
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
