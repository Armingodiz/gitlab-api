package app

import (
	"github.com/ArminGodiz/gitlab-api/src/services/gitlabService"
	"github.com/gin-gonic/gin"
)

var (
	// router will be in charge of creating different goRoutines for different request we want to handle .
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	gitlabService.SetCache(-1)
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
