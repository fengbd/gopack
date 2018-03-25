package router

import (
	"github.com/gin-gonic/gin"
)

// Router is base router object
var (
	Router   = gin.Default()
	apiGroup = Router.Group("/api")
)
