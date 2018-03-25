package project

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/IdanLoo/gopack/model"
	"github.com/IdanLoo/gopack/util/config"
	"github.com/IdanLoo/gopack/util/git"
)

// Project is a map to repository
type Project struct {
	Name   string        `json:"name"`
	URL    string        `json:"url"`
	Path   *Path         `json:"path"`
	Pusher *model.Pusher `json:"pusher"`
}

// New to construct a Project object and make dirs.
// If project not exist, then create.
func New(name, branch, url string, pusher *model.Pusher) (*Project, error) {
	proj := newProject(name, url, pusher)

	if !config.IsExist(proj.Path.Root) {
		if err := proj.createDir(); err != nil {
			return nil, err
		}
	}

	if err := proj.clone(branch); err != nil {
		return nil, err
	}

	in := newInfo(branch, pusher, time.Now().Unix())
	if err := in.SaveToDir(proj.Path.Src); err != nil {
		return nil, err
	}

	return proj, proj.Build()
}

// Of to get project of name
func Of(name string) (*Project, error) {
	path := workspace + name + "/"

	if !config.IsExist(path) {
		return nil, os.ErrNotExist
	}

	url, err := git.OriginURL(path)
	if err != nil {
		return nil, err
	}

	in, err := parseInfoOfDir(path + "src/")
	if err != nil {
		return nil, err
	}

	return newProject(name, url, in.Pusher), nil
}

// All to get all projects)
func All() []*Project {
	files, err := ioutil.ReadDir(workspace)

	if err != nil {
		return []*Project{}
	}

	projs := make([]*Project, 0)

	for _, file := range files {
		if file.IsDir() {
			proj, err := Of(file.Name())

			if err != nil {
				panic(err)
			}

			projs = append(projs, proj)
		}
	}

	return projs
}

func newProject(name, url string, pusher *model.Pusher) *Project {
	root := workspace + name + "/"
	path := NewPath(root, root+"src/", root+"bin/")

	return &Project{
		name,
		url,
		path,
		pusher,
	}
}
