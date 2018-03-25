package project

import (
	"io/ioutil"
	"os"

	"github.com/IdanLoo/gopack/util/config"
)

// Binary is the production after build
type Binary struct {
	project *Project
	Name    string `json:"name"`
}

func newBinary(project *Project, name string) *Binary {
	return &Binary{
		project,
		name,
	}
}

// Binaries to get all binaries of project
func (p *Project) Binaries() []*Binary {
	files, err := ioutil.ReadDir(p.Path.Bin)
	if err != nil {
		return nil
	}

	bins := make([]*Binary, 0)
	for _, file := range files {
		if file.IsDir() {
			bin := newBinary(p, file.Name())
			bins = append(bins, bin)
		}
	}

	return bins
}

// BinaryWith to get binary with name
func (p *Project) BinaryWith(name string) (*Binary, error) {
	path := p.Path.Bin + name + "/"
	if !config.IsExist(path) {
		return nil, os.ErrNotExist
	}

	return newBinary(p, name), nil
}

// Run to Run this binary
func (b *Binary) Run() error {
	op, err := b.project.opertaion()
	if err != nil {
		return err
	}

	path, err := b.path()
	if err != nil {
		return err
	}

	cmds := []string{"cd $dst"}
	cmds = append(cmds, op.Run...)

	return b.project.exec("nope", path, cmds)
}

func (b *Binary) path() (string, error) {
	path := b.project.Path.Bin + b.Name + "/"
	if !config.IsExist(path) {
		return "", os.ErrNotExist
	}

	return path, nil
}
