package git

import (
	"log"
	"os/exec"
)

var git string

func init() {
	var err error
	git, err = exec.LookPath("git")

	if err != nil {
		log.Fatal(err)
	}
}
