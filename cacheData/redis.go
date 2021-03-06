package cacheDataOrm

import (
	"cache-component/mcpack2"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type CacheRedisData struct {
	CacheVersion string
	RealData     interface{}
}

func SetRedis(key string, data CacheRedisData) error {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return err
	}
	defer c.Close()
	marshal, _ := mcpack2.Marshal(data)
	_, err = c.Do("SET", key, marshal)
	return err
}

func GetFromRedis(key string) (interface{}, error) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return "", err
	}
	defer c.Close()

	value, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		return "", err
	}
	return value, err

}
