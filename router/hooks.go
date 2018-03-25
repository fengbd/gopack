package router

import (
	"net/http"

	"github.com/fengbd/gopack/util/project"

	"github.com/fengbd/gopack/model"
	"github.com/gin-gonic/gin"
)

// HooksGroup is a router group
var HooksGroup = Router.Group("/hooks")

func push(ctx *gin.Context) {
	body := &model.PushBody{}

	if err := ctx.BindJSON(body); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}

	_, err := project.New(
		body.Repository.Name,
		body.Branch(),
		body.Repository.URL,
		body.Pusher,
	)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
	})
}

func init() {
	HooksGroup.
		POST("/push", push)
}
