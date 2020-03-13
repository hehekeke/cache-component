package main

import (
	"cache-component/originOrm"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//testOriginOrmCreate()
	//TestCacheDataInsert()
	//testOriginOrmFindByPk()
	//testOriginOrmFindByPkWithCache()

	//testOriginOrmUpdate()
	//testOriginOrmUpdateWithCache()
	//testOriginOrmDel()
	testOriginOrmDelWithCache()

	//TestRedis()

}

//插入 =================
func testOriginOrmCreate() {
	user := originOrm.User{
		Id:    0,
		Name:  "1",
		Age:   1,
		Sex:   1,
		Phone: "1",
	}
	user.Insert()
}

func TestCacheDataInsert()  {
	user := originOrm.User{
		Id:    0,
		Name:  "1",
		Age:   1,
		Sex:   1,
		Phone: "1",
	}
	user.InsertWithCache()
}

//查找 =================

func testOriginOrmFindByPk() {
	user := originOrm.User{
		Id: 8,
	}
	users := user.Query()
	fmt.Println(users)
}

func testOriginOrmFindByPkWithCache() {
	user := originOrm.User{
		Id: 1,
	}
	users := user.QueryWithCache()
	fmt.Println(users)
}

//更新 =================
func testOriginOrmUpdate() {
	user := originOrm.User{
		Id: 8,
	}
	user.Update()
}
func testOriginOrmUpdateWithCache() {
	user := originOrm.User{
		Id: 1,
	}
	user.UpdateWithCache()
}

//删除 =================

func testOriginOrmDel() {
	user := originOrm.User{
		Id: 8,
	}
	user.Del()
}

func testOriginOrmDelWithCache() {
	user := originOrm.User{
		Id: 1,
	}
	user.DelWithCache()
}

//按照之后的Orm
func testCacheDataOrm() {
	user := originOrm.User{
		Id:    0,
		Name:  "",
		Age:   0,
		Sex:   0,
		Phone: "",
	}
	user.Insert()
}


func TestRedis()  {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang1111")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}