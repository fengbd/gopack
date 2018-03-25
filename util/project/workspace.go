package project

import (
	"log"
	"os"

	"github.com/fengbd/gopack/util/config"
)

var workspace = config.Global.Workspace

func init() {
	if !config.IsExist(workspace) {
		if err := createDir(workspace); err != nil {
			log.Fatal(err)
		}
	}
}

func createDir(path string) error {
	return os.Mkdir(path, os.ModeDir|os.ModePerm)
}
