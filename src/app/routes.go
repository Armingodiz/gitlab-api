package app

import (
	"github.com/ArminGodiz/gitlab-api/src/controllers"
	gController "github.com/ArminGodiz/gitlab-api/src/controllers/gitlabController"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping) // for cloud to hit it and understand our web server is still running
	router.POST("/token", gController.SetToken)
	router.GET("/projects/:id/repository_tree", gController.GetDetails)
}
