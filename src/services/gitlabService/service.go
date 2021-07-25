package gitlabService

import (
	"errors"
	"fmt"
	cachePkg "github.com/ArminGodiz/gitlab-api/src/cache"
	"net/http"
)

var (
	cache cachePkg.Cache
)

func SetCache(port int) {
	cache = cachePkg.GetNewCache(port)
}

func SaveToken(id string, token string) error {
	return cache.Set(id, token)
}

func GetDetails(projId, uid string) (*http.Response, error) {
	// check token existence
	token, err := cache.Get(uid)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("token not found")
	}
	url := "https://gitlab.com/api/v4/projects/" + projId + "/repository/tree?access_token=" + token
	resp, err2 := http.Get(url)
	return resp, err2
}
