package model

// Pusher is structure of pusher
type Pusher struct {
	Username string `json:"username" binding:"required"`
	Avatar   string `json:"avatar_url" binding:"required"`
}
