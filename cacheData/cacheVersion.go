package cacheDataOrm

import (
	"encoding/json"
	"fmt"
	"runtime"
)

/**
获得全局的versio版本
*/
func GetGlobal() string {
	return "111"
}

/**
获得不适用全局版本的话 ， 返回是0
*/
func GetNoUserGlobal() string {
	return "0"
}

// 根据传入的函数名 ，方法名 + 参数  ，得出唯一的key ，
func GetReisKey(key interface{}) string {
	pc, file, _, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	var keyBytes []byte
	keyBytes, _ = json.Marshal(key)

	var getMd5OriginKey = fmt.Sprintf("%s%s%s", f.Name(), file, string(keyBytes))
	return Md5V(getMd5OriginKey)
}

func GetReisKeyByPk(tableName string, key interface{}) string {
	return fmt.Sprintf("find_by_pk:%s:%d", tableName, key)
}
