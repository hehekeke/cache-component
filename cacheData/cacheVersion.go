package cacheDataOrm

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
