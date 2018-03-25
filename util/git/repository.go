package git

import (
	"fmt"
	"strings"

	"github.com/IdanLoo/gopack/util/cmd"
)

// Clone a repository
func Clone(url, branch, dest string) error {
	c := cmd.New(fmt.Sprintf("%s clone -b %s %s %s", git, branch, url, dest))
	return c.Run()
}

// OriginURL to get the origin remote url project folder
func OriginURL(path string) (string, error) {
	cmds := []string{
		"cd " + path + "/src/",
		git + " remote get-url origin",
	}
	c := cmd.New(cmds...)

	bytes, err := c.Output()
	if err != nil {
		return "", err
	}

	url := string(bytes)
	url = strings.Trim(url, "\n")

	return url, nil
}
