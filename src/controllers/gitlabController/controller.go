package gitlabController

import (
	"bytes"
	"errors"
	"fmt"
	gitService "github.com/ArminGodiz/gitlab-api/src/services/gitlabService"
	"github.com/gin-gonic/gin"
	"net/http"
)

type postData struct {
	UserId            string `json:"userId"`
	GitlabAccessToken string `json:"gitlabAccessToken"`
}

func SetToken(c *gin.Context) {
	var data postData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid json body"))
		return
	}
	if err := gitService.SaveToken(data.UserId, data.GitlabAccessToken); err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("error in saving cache"))
	}
	c.JSON(http.StatusOK, nil)
}

func GetDetails(c *gin.Context) {
	projId := c.Param("id")
	val := c.Request.Header.Get("userId")
	resp, err := gitService.GetDetails(projId, val)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	buf := new(bytes.Buffer)
	_, err2 := buf.ReadFrom(resp.Body)
	if err2 != nil {
		c.Status(400)
		return
	}
	c.Status(200)
	fmt.Fprint(c.Writer, buf.String())
}
