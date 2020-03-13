package cacheDataOrm

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func SetRedis(key string, data interface{}) error {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return err
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang1111")
	return err
}

func GetFromRedis(key string) (interface{}, error) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return "", err
	}
	defer c.Close()

	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		return "", err
	}
	return value, err

}
