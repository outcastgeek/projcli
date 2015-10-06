package django

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
)

func NewDjango(c *cli.Context) {
	if len(c.Args()) > 0 {
		projectName := c.Args()[0]
		fmt.Println("Creating a new Django Application with name: ", projectName)
		cmd := "django-admin.py"
		args := []string{"startproject", projectName}
		if err := exec.Command(cmd, args...).Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Successfully created Django Application with name: ", projectName)
	}
}

func MigrationsDjango(c *cli.Context) {
	currDir, err := os.Getwd()
	if err != nil {
		handleErr(err)
	}
	fmt.Println("Making Migrations...")
	cmd := currDir + "/manage.py"
	args := []string{"makemigrations"}
	if err = exec.Command(cmd, args...).Run(); err != nil {
		handleErr(err)
	}
	fmt.Println("Successfully generated Migrations")
}

func MigrateDjango(c *cli.Context) {
	currDir, err := os.Getwd()
	if err != nil {
		handleErr(err)
	}
	fmt.Println("Migrating Database...")
	cmd := currDir + "/manage.py"
	args := []string{"migrate"}
	if err = exec.Command(cmd, args...).Run(); err != nil {
		handleErr(err)
	}
	fmt.Println("Successfully migrated Database!")
}

func handleErr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
