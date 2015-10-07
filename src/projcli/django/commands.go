package django

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
	"projcli/utils"
	"time"
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

func New(configName string) {
	pipChan := make(chan string, 1)
	go func() {
		pipTxt, err := utils.HttpGet("bootstrap.pypa.io/get-pip.py")
		if err != nil {
			utils.HandleErr(err)
		}
		pipChan <- pipTxt
	}()
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
	select {
	case pipTxt := <-pipChan:
		pipPath := workDir + "/" + projectName.(string) + "/get-pip.py"
		requirementsPath := workDir + "/" + projectName.(string) + "/requirements.txt"
		ioutil.WriteFile(pipPath, []byte(pipTxt), 0755)
		ioutil.WriteFile(requirementsPath, []byte(Requirements), 0644)
		fmt.Println("Installing pip...")
		cmd = pipPath
		args = []string{}
		utils.RunCmd(cmd, args)
		fmt.Println("Installing requirements...")
		cmd = "pip"
		args = []string{"install", "-r", requirementsPath}
		utils.RunCmd(cmd, args)
	case <-time.After(1 * time.Second):
		fmt.Println("Could not install pip")
	}
}

func Migrations(configName string) {
	setup(configName)
	fmt.Println("Making Migrations...")
	cmd := workDir + "/manage.py"
	args := []string{"makemigrations"}
	utils.RunCmd(cmd, args)
}

func Migrate(configName string) {
	setup(configName)
	fmt.Println("Migrating Database...")
	cmd := workDir + "/manage.py"
	args := []string{"migrate"}
	utils.RunCmd(cmd, args)
}

func App(configName string) {
	setup(configName)
	applications := viper.Get("applications").([]interface{})
	numOfApps := len(applications)
	done := make(chan bool, numOfApps)
	for _, app := range applications {
		for appName, _ := range app.(map[interface{}]interface{}) {
			go func(appName string) {
				fmt.Println("Creating Application: ", appName)
				cmd := workDir + "/manage.py"
				args := []string{"startapp", appName}
				utils.RunCmd(cmd, args)
				done <- true
			}(appName.(string))
		}
	}
	for _ = range applications {
		<-done
	}
}
