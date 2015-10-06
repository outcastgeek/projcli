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
	flagSet.StringVar(&configName, "conf", defaultDjConfig, "Create a new Django Application: projcli {cmd} -conf")
}

func usageFunc() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "  -h   : show help usage")
}

func main() {
	flag.Usage = usageFunc

	// flag.Parse()

	// if flag.NFlag() == 0 {
	// 	// usageFunc()
	//   os.Exit(1)
	// }

	// if flag.NFlag() == 0 {
	// 	flag.Usage()
	// 	os.Exit(1)
	// }

	// cmdName := flag.Args()[0]
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
