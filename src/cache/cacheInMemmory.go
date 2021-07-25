package cache

import (
	"errors"
)

type InMemoryCache struct {
	MCache map[string]string
}

func (mc *InMemoryCache) Set(key string, value string) error {
	mc.MCache[key] = value
	if mc.MCache[key] != value {
		return errors.New("error in setting token")
	}
	return nil
}
func (mc *InMemoryCache) Get(key string) (string, error) {
	value, ok := mc.MCache[key]
	if !ok {
		return "", errors.New("error in getting token")
	}
	return value, nil
}
