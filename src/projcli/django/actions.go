package django

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
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

func setup(configName string) {
	extension := filepath.Ext(configName)
	_configName := configName[0 : len(configName)-len(extension)]
	viper.AddConfigPath(workDir)
	viper.SetConfigName(_configName)
	viper.SetConfigType("yaml")
	viper.Set("Verbose", true)
	err := viper.ReadInConfig()
	if err != nil {
		utils.HandleErr(err)
	}
}

func NewDjango(configName string) {
	setup(configName)
	projectName := viper.Get("project")
	fmt.Println("Creating a new Django Application with name: ", projectName)
	cmd := "django-admin.py"
	args := []string{"startproject", projectName.(string)}
	utils.RunCmd(cmd, args)
	fmt.Println("Copying: ", configName)
	cmd = "cp"
	args = []string{workDir + "/" + configName, workDir + "/" + projectName.(string)}
	utils.RunCmd(cmd, args)
	// if len(c.Args()) > 0 {
	// 	projectName := c.Args()[0]
	// 	fmt.Println("Creating a new Django Application with name: ", projectName)
	// 	cmd := "django-admin.py"
	// 	args := []string{"startproject", projectName.(string)}
	// 	utils.RunCmd(cmd, args)
	// }
}

func MigrationsDjango(configName string) {
	setup(configName)
	fmt.Println("Making Migrations...")
	cmd := workDir + "/manage.py"
	args := []string{"makemigrations"}
	utils.RunCmd(cmd, args)
}

func MigrateDjango(configName string) {
	setup(configName)
	fmt.Println("Migrating Database...")
	cmd := workDir + "/manage.py"
	args := []string{"migrate"}
	utils.RunCmd(cmd, args)
}
