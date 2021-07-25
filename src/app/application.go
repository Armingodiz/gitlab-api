package app

import (
	"github.com/ArminGodiz/gitlab-api/src/services/gitlabService"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	// router will be in charge of creating different goRoutines for different request we want to handle .
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	config := GetConfig()
	gitlabService.SetCache(config.CachePort)
	err := router.Run(":" + strconv.Itoa(config.ListenerPort))
	if err != nil {
		panic(err)
	}
}
