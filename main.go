package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"test/originOrm"
)

func main() {
	//testOriginOrmCreate()
	//testOriginOrmFindByPk()
	//testOriginOrmUpdate()
	testOriginOrmDel()
}

//按照之前Orm
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

func testOriginOrmFindByPk() {
	user := originOrm.User{
		Id: 8,
	}
	users := user.Query()
	fmt.Println(users)
}

//按照之前Orm
func testOriginOrmUpdate() {
	user := originOrm.User{
		Id: 8,
	}
	user.Update()
}

func testOriginOrmDel() {
	user := originOrm.User{
		Id: 8,
	}
	user.Del()
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
