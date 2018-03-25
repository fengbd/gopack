package router

import (
	"net/http"

	"github.com/IdanLoo/gopack/util/project"
	"github.com/gin-gonic/gin"
)

// ProjectsGroup is a router group
var ProjectsGroup = apiGroup.Group("/projects")

func init() {
	ProjectsGroup.
		GET("", allProjects).
		GET("/:name", allBinaries).
		GET("/:name/build", buildProject).
		POST("/:name/:binary/run", runBinary)
}

func allProjects(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"projects": project.All(),
	})
}

func allBinaries(ctx *gin.Context) {
	name := ctx.Param("name")

	proj, err := project.Of(name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "无此项目",
		})
		return
	}

	bins := proj.Binaries()
	ctx.JSON(http.StatusOK, gin.H{
		"binaries": bins,
	})
}

func buildProject(ctx *gin.Context) {
	name := ctx.Param("name")

	proj, err := project.Of(name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "无此项目",
		})
		return
	}

	if err := proj.Build(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "编译失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
	})
}

func runBinary(ctx *gin.Context) {
	name := ctx.Param("name")
	binary := ctx.Param("binary")

	proj, err := project.Of(name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "无此项目",
		})
		return
	}

	bin, err := proj.BinaryWith(binary)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "无此二进制包",
		})
		return
	}

	if err := bin.Run(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
	})
}
