package model

// Repository is structure of repository
type Repository struct {
	Name string `json:"name" binding:"required"`
	URL  string `json:"clone_url" binding:"required"`
}
