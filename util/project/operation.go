package project

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/IdanLoo/gopack/util/cmd"
	"github.com/IdanLoo/gopack/util/config"
)

// operation is some shell command in gopack.yaml
type operation struct {
	Build []string
	Run   []string
}

// NewOperation is constructor of operation
func newOperation(path string) (*operation, error) {
	if !config.IsExist(path) {
		return nil, os.ErrNotExist
	}

	op := &operation{}

	if err := config.Parse(path, op); err != nil {
		return nil, err
	}

	return op, nil
}

func (p *Project) opertaion() (*operation, error) {
	configPath := p.Path.Src + "gopack.yaml"
	return newOperation(configPath)
}

// Build to run build commands
func (p *Project) Build() error {
	var (
		in  *info
		op  *operation
		err error
	)

	if in, err = parseInfoOfDir(p.Path.Src); err != nil {
		return err
	}

	if op, err = p.opertaion(); err != nil {
		return err
	}

	tm := time.Unix(in.PushTime, 0).Format("0102030405")
	branchWithoutSlash := strings.Replace(in.Branch, "/", "-", -1)

	dst := fmt.Sprintf(
		"%s%s-%s-%s/",
		p.Path.Bin, in.Pusher.Username, branchWithoutSlash, tm,
	)

	cmds := []string{
		"mkdir " + dst,
		"cd " + p.Path.Src,
	}
	cmds = append(cmds, op.Build...)

	if err = p.exec(in.Branch, dst, cmds); err != nil {
		return err
	}

	return in.SaveToDir(dst)
}

func (p *Project) exec(branch, dst string, cmds []string) error {
	c := cmd.New(p.env(branch, dst)...)
	c.Append(cmds...)

	return c.Run()
}
