package public

import (
	"encoding/json"
	"errors"
	"fmt"
	"redis-api/Connections"
	"time"
)

//set
func (p public) SetCache(keys string, data interface{}, expired time.Duration) (err error) {
	conn := Connections.RedisConnection
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = conn.Set(keys, string(dataBytes), expired).Err()
	return
}

//get
func (p public) GetCache(keys string, data interface{}) (err error) {
	conn := Connections.RedisConnection
	dataCache, err := conn.Get(keys).Result()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(dataCache), &data)
	return
}

//hmset
func (p public) SetHMCache(keys string, data map[string]interface{}, expired time.Duration) (err error) {
	conn := Connections.RedisConnection
	err = conn.HMSet(keys, data).Err()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = conn.Expire(keys, expired).Err()
	return
}

//hmget
func (p public) GetHMCache(keys, field string, data interface{}) (err error) {
	conn := Connections.RedisConnection
	dataRedis, err := conn.HMGet(keys, field).Result()
	if err != nil {
		return
	}
	fmt.Println("DATA : ", dataRedis)
	if len(dataRedis) == 0 || dataRedis[0] == nil {
		err = errors.New("Data redis empty")
		return
	}
	err = json.Unmarshal([]byte(dataRedis[0].(string)), &data)
	return
}

//hgetall
func (p public) HGetAllCache(keys string) (res map[string]string, err error) {
	conn := Connections.RedisConnection
	res, err = conn.HGetAll(keys).Result()
	return
}

//hdel
func (p public) HDelCache(keys string, field ...string) (err error) {
	conn := Connections.RedisConnection
	err = conn.HDel(keys, field...).Err()
	return
}

//del
func (p public) DelCache(keys ...string) (err error) {
	conn := Connections.RedisConnection
	err = conn.Del(keys...).Err()
	return
}
