package cmd

import (
	"os/exec"
	"strings"
)

// Cmd is a command class.
type Cmd struct {
	list []string
}

// New to construct a Cmd object
func New(cmds ...string) *Cmd {
	return &Cmd{cmds}
}

// Run to run the cmd
func (c *Cmd) Run() error {
	cmd := c.executable()
	return cmd.Run()
}

// Output to get the output of the cmd
func (c *Cmd) Output() ([]byte, error) {
	cmd := c.executable()
	return cmd.Output()
}

// Append to append some commands to list
func (c *Cmd) Append(cmd ...string) {
	c.list = append(c.list, cmd...)
}

func (c *Cmd) executable() *exec.Cmd {
	cmdStr := strings.Join(c.list, " && ")
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	return cmd
}
