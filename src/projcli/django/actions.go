package django

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
	"os"
	"projcli/utils"
)

var workDir string

func init() {
	wd, err := os.Getwd()
	if err != nil {
		utils.HandleErr(err)
	}
	workDir = wd
}

func setup() {
	viper.AddConfigPath(workDir)
	viper.SetConfigName("django")
	viper.SetConfigType("yaml")
	viper.Set("Verbose", true)
	err := viper.ReadInConfig()
	if err != nil {
		utils.HandleErr(err)
	}
}

func NewDjango(c *cli.Context) {
	setup()
	projectName := viper.Get("project")
	// projectName := c.Args()[0]
	fmt.Println("Creating a new Django Application with name: ", projectName)
	cmd := "django-admin.py"
	args := []string{"startproject", projectName.(string)}
	utils.RunCmd(cmd, args)
	// if len(c.Args()) > 0 {
	// 	projectName := c.Args()[0]
	// 	fmt.Println("Creating a new Django Application with name: ", projectName)
	// 	cmd := "django-admin.py"
	// 	args := []string{"startproject", projectName.(string)}
	// 	utils.RunCmd(cmd, args)
	// }
}

func MigrationsDjango(c *cli.Context) {
	fmt.Println("Making Migrations...")
	cmd := workDir + "/manage.py"
	args := []string{"makemigrations"}
	utils.RunCmd(cmd, args)
}

func MigrateDjango(c *cli.Context) {
	fmt.Println("Migrating Database...")
	cmd := workDir + "/manage.py"
	args := []string{"migrate"}
	utils.RunCmd(cmd, args)
}
