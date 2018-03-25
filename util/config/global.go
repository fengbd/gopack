package config

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

// GlobalConfig is global config structure
type GlobalConfig struct {
	Workspace string
}

// Global is global config, save as ~/.gopack
var (
	configPath, _    = homedir.Expand("~/.gopack")
	defaultWorkspace = "~/gopack-projects/"
	Global           = newGlobalConfig()
)

func newGlobalConfig() *GlobalConfig {
	var (
		global = &GlobalConfig{}
		err    error
	)

	if !IsExist(configPath) {
		if err = createConfigFile(); err != nil {
			log.Fatal(err)
		}
	}

	if err = Parse(configPath, global); err != nil {
		log.Fatal(err)
	}

	if global.Workspace == "" {
		if global.Workspace, err = configureWorkspace(); err != nil {
			log.Fatal(err)
		}
	}

	if err = Save(configPath, global); err != nil {
		log.Fatal(err)
	}

	return global
}

func createConfigFile() error {
	_, err := os.Create(configPath)
	return err
}

func configureWorkspace() (string, error) {
	var workspace string

	fmt.Printf("Workspace Path (default is %s): ", defaultWorkspace)
	if len, err := fmt.Scanln(&workspace); len == 0 || err != nil {
		workspace = defaultWorkspace
	}

	workspace, _ = homedir.Expand(workspace)
	if workspace[len(workspace)-1] != '/' {
		workspace += "/"
	}

	return workspace, nil
}
