package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"projcli/django"
)

var (
	app   *cli.App
	tasks []string
)

func init() {
	tasks = []string{"newdjango"}
	app = cli.NewApp()
	app.Name = "Project CLI"
	app.Usage = "Manage your Projects"
	app.EnableBashCompletion = true
	setupActions()
}

func setupActions() {
	app.Commands = []cli.Command{
		{
			Name:    "newdjango",
			Aliases: []string{"ndj"},
			Usage:   "Create a new Django Application: projcli newdjango {appName}",
			Action:  django.NewDjango,
			BashComplete: func(c *cli.Context) {
				// This will complete if no args are passed
				if len(c.Args()) > 0 {
					return
				}
				for _, t := range tasks {
					fmt.Println(t)
				}
			},
		},
		{
			Name:    "migrationsdjango",
			Aliases: []string{"mdj"},
			Usage:   "Make Django Migrations",
			Action:  django.MigrationsDjango,
			BashComplete: func(c *cli.Context) {
				// This will complete if no args are passed
				if len(c.Args()) > 0 {
					return
				}
				for _, t := range tasks {
					fmt.Println(t)
				}
			},
		},
		{
			Name:    "migratedjango",
			Aliases: []string{"midj"},
			Usage:   "Migrate Django Application",
			Action:  django.MigrateDjango,
			BashComplete: func(c *cli.Context) {
				// This will complete if no args are passed
				if len(c.Args()) > 0 {
					return
				}
				for _, t := range tasks {
					fmt.Println(t)
				}
			},
		},
	}
}

func main() {
	app.Run(os.Args)
}
