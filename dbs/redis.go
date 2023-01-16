package dbs

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

var cxt = context.Background()

func GetRedisObject() redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:30001",
		Password: "",
		DB:       0,
	})
	return *rdb
}

func RedisRead(key string, redisClient redis.Client) (string, error) {
	p, err := redisClient.Get(cxt, key).Result()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return p, nil
}

func RedisWrite(key string, object interface{}, redisClient redis.Client) error {
	var (
		err     error
		jsonStr []byte
	)
	jsonStr, err = json.Marshal(object)
	err = redisClient.Set(cxt, key, jsonStr, 0).Err()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
