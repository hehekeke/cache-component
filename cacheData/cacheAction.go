package cacheDataOrm

import (
	"cache-component/mcpack2"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
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
	dBWithCache.TableName = table
	return dBWithCache
}

/**
增
*/
func (dBWithCache DBWithCache) Create(data interface{}) *DBWithCache {
	res := dBWithCache.DB.Table(dBWithCache.TableName).Create(data)

	if dBWithCache.IsUserCache == false {
		return &dBWithCache
	}
	var newObject = data
	res.Row().Scan(newObject)

	return &dBWithCache
}
func (dBWithCache *DBWithCache) SetRedisByPk(pk interface{}, data interface{}) {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Array, reflect.Slice:
		fmt.Println("数组返回")
	case reflect.Struct:
	default:
		fmt.Println(reflect.TypeOf(data).Kind())
		fmt.Println("以上类型都不是")
		return
	}
	_, isGetData := dBWithCache.GetRedisByPk(pk)
	if isGetData == true {
		return
	}
	reisKey := GetReisKeyByPk(dBWithCache.TableName, pk)
	redisData := CacheRedisData{}
	redisData.CacheVersion = "1.0"
	redisData.RealData = data
	SetRedis(reisKey, redisData)
}

func (dBWithCache *DBWithCache) GetRedisByPk(pk interface{}) (data interface{}, isGetData bool) {
	reisKey := GetReisKeyByPk(dBWithCache.TableName, pk)
	fromRedis, err := GetFromRedis(reisKey)
	if err != nil {
		fmt.Println(err.Error())
	}
	if fromRedis != "" {
		return fromRedis, true
	}
	return "", false
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

type User struct {
	Id    int64
	Name  string
	Age   string
	Sex   string
	Phone string
}

func (dBWithCache *DBWithCache) QueryPkId(value interface{}, id interface{}) *DBWithCache {
	data, isGetData := dBWithCache.GetRedisByPk(id)
	if isGetData == true {
		redisData := CacheRedisData{}
		mcpack2.Unmarshal(data.([]byte), &redisData)
		dBWithCache.DB.Table("user").Row().Scan(&value)
		return dBWithCache
	}
	dBWithCache.DB.Find(value, " id = ?", id)

	return dBWithCache
}
