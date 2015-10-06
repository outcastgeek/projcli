package main

import (
	"flag"
	"fmt"
	"os"
	"projcli/django"
)

const (
	defaultDjConfig = "django.yaml"
)

var (
	configName string
	flagSet    *flag.FlagSet
)

func init() {
	flagSet = flag.NewFlagSet("configuration", flag.ContinueOnError)
	flagSet.StringVar(&configName, "djconf", defaultDjConfig, "Manage Django Application: projcli {djcmd} -djconf")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "  -h   : show help usage")
		flagSet.PrintDefaults()
	}
}

func main() {

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	cmdName := os.Args[1]
	flagSet.Parse(os.Args[2:])

	switch cmdName {
	case "djnew":
		django.New(configName)
	case "djmigrations":
		django.Migrations(configName)
	case "djmigrate":
		django.Migrate(configName)
	case "djapp":
		django.App(configName)
	default:
		flag.Usage()
	}
}
