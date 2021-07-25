package gitlabController

import (
	"errors"
	gitService "github.com/ArminGodiz/gitlab-api/src/services/gitlabService"
	"github.com/gin-gonic/gin"
	"net/http"
)

type postData struct {
	userId            string `json:"userId"`
	gitlabAccessToken string `json:"gitlabAccessToken"`
}

func SetToken(c *gin.Context) {
	var data postData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid json body"))
		return
	}
	if err := gitService.SaveToken(data.userId, data.gitlabAccessToken); err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("error in saving cache"))
	}
	c.JSON(http.StatusOK, nil)
}

func GetDetails(c *gin.Context) {
	projId := c.Param("id")
	values, _ := c.Request.Header["userId"]
	if len(values) == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	resp, err := gitService.GetDetails(projId, values[0])
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
