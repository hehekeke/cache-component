package cacheDataOrm

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type DBWithCache struct {
	DB          *gorm.DB
	IsUserCache bool
	TableName   string
}

func (dBWithCache *DBWithCache) SetCache() *DBWithCache {
	dBWithCache.IsUserCache = true
	return dBWithCache
}

func (dBWithCache *DBWithCache) Table(table string) *DBWithCache {
	dBWithCache.IsUserCache = true
	return dBWithCache
}

/**
增
*/
func (dBWithCache DBWithCache) Create(data interface{}) {
	res := dBWithCache.DB.Table(dBWithCache.TableName).Create(data)
	if dBWithCache.IsUserCache == false {
		return
	}
	var newObject = data
	res.Row().Scan(newObject)
	getReisKey := GetReisKey(data)

	redisData := CacheRedisData{}
	redisData.CacheVersion = "1.0"
	redisData.RealData = newObject

	SetRedis(getReisKey, redisData)
}

/**
删
*/
func (dBWithCache DBWithCache) Delete(value interface{}, where ...interface{}) {
	dBWithCache.DB.Where(where).Delete(value)
	if dBWithCache.IsUserCache == true {
		fmt.Println("说明需要添加缓存了")
	}
}

/**
改
*/
func (dBWithCache DBWithCache) Update(attrs interface{}) {
	dBWithCache.DB.Model(&attrs).Table("user").Update(attrs)
	if dBWithCache.IsUserCache == true {
		fmt.Println("说明需要添加缓存了")
	}
}

/**
查
*/
func (dBWithCache DBWithCache) Query(value interface{}, where ...interface{}) {
	dBWithCache.DB.Find(value, where)
	if dBWithCache.IsUserCache == true {
		fmt.Println("说明需要添加缓存了")
	}
}
