package model

import (
	"strings"
)

// PushBody is the body of push event sent by git server
type PushBody struct {
	Ref        string      `json:"ref" binding:"required"`
	Repository *Repository `json:"repository" binding:"required"`
	Pusher     *Pusher     `json:"pusher" binding:"required"`
}

// Branch to get the branch name of this ref
func (p *PushBody) Branch() string {
	refs := strings.Split(p.Ref, "/")[2:]
	name := strings.Join(refs, "/")
	return name
}
