package project

import (
	"os"

	"github.com/IdanLoo/gopack/util/git"
)

// Path is a structure of project path
type Path struct {
	Root string `json:"root"`
	Src  string `json:"src"`
	Bin  string `json:"bin"`
}

// NewPath to construct a Pro
func NewPath(root, src, bin string) *Path {
	return &Path{
		root, src, bin,
	}
}

func (p *Project) createDir() error {
	for _, path := range []string{p.Path.Root, p.Path.Src, p.Path.Bin} {
		if err := createDir(path); err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) cleanSrc() error {
	return os.RemoveAll(p.Path.Src)
}

func (p *Project) clone(branch string) error {
	if err := p.cleanSrc(); err != nil {
		return err
	}

	return git.Clone(p.URL, branch, p.Path.Src)
}
