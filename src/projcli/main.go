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
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "  -h   : show help usage")
	}
	flagSet = flag.NewFlagSet("configuration", flag.ContinueOnError)
	flagSet.StringVar(&configName, "conf", defaultDjConfig, "Create a new Django Application: projcli {cmd} -conf")
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
		django.NewDjango(configName)
	case "djmigrations":
		django.MigrationsDjango(configName)
	case "djmigrate":
		django.MigrateDjango(configName)
	default:
		flag.Usage()
	}
}
