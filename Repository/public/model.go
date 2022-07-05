package public

import (
	"time"
)

type PublicInterface interface {
	SetCache(keys string, data interface{}, expired time.Duration) (err error)
	GetCache(keys string, data interface{}) (err error)
	SetHMCache(keys string, data map[string]interface{}, expired time.Duration) (err error)
	GetHMCache(keys, field string, data interface{}) (err error)
	HGetAllCache(keys string) (res map[string]string, err error)
	HDelCache(keys string, field ...string) (err error)
	DelCache(keys ...string) (err error)
}

type public struct{}

func InitPublic() PublicInterface {
	return &public{}
}
