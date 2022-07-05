package Connections

import "github.com/go-redis/redis"

var RedisConnection *redis.Client

func RedisConn() (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:49153",
		Password: "redispw",
		DB:       0,
	})
	return
}

func InitiateConnection() {
	RedisConnection = RedisConn()
}
