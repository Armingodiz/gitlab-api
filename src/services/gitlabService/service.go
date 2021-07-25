package gitlabService

import (
	cachPkg "github.com/ArminGodiz/gitlab-api/src/cache"
	"net/http"
)

var (
	cache cachPkg.Cache
)

func SetCache(port int) {
	cache = cachPkg.GetNewCache(port)
}

func SaveToken(id string, token string) error {
	return cache.Set(id, token)
}

func GetDetails(projId, uid string) (*http.Response, error) {
	// check token existence
	token, err := cache.Get(uid)
	if err != nil {
		return nil, nil
	}
	resp, err2 := http.Get("https://gitlab.com/api/v4/projects/" + projId + "/repository/tree?access_token=" + token)
	return resp, err2
}
