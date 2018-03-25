package project

import (
	"os"

	"github.com/IdanLoo/gopack/model"
	"github.com/IdanLoo/gopack/util/config"
)

const infoFilename = ".info"

type info struct {
	Branch   string        `json:"branch"`
	Pusher   *model.Pusher `json:"pusher"`
	PushTime int64         `json:"pushTime"`
}

func newInfo(branch string, pusher *model.Pusher, pushTime int64) *info {
	return &info{
		branch,
		pusher,
		pushTime,
	}
}

func parseInfoOfDir(path string) (*info, error) {
	if path += infoFilename; !config.IsExist(path) {
		return nil, os.ErrNotExist
	}

	info := &info{}
	if err := config.Parse(path, info); err != nil {
		return nil, err
	}

	return info, nil
}

func (in *info) SaveToDir(path string) error {
	path += infoFilename
	return config.Save(path, in)
}

func (p *Project) env(branch, dst string) []string {
	return []string{
		"dst=" + dst,
		"src=" + p.Path.Src,
		"branch=" + branch,
	}
}
