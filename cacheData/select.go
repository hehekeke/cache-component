package cacheDataOrm

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

var dbWithCache *DBWithCache

type User struct {
	Id    int
	Name  string
	Age   int
	Sex   byte
	Phone string
}

func init() {
	var err error
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	dbWithCache.DB = db
	//设置全局表名禁用复数
	db.SingularTable(true)
}

func (user *User) Insert() {
	dbWithCache.SetCache().Table("user").Create(user)
	// db.Table("user").Create(user)
}

func (user *User) Update() {

	user = &User{Id: user.Id, Name: "xiaoming"}
	dbWithCache.SetCache().Update(user)
	//db.Model(&user).Update(user)
}

func (user *User) Del() {
	dbWithCache.SetCache().Delete(user,fmt.Sprintf("id = %d",user.Id))
	//db.Where("id =  ?", user.Id).Delete(&user)
}

func (user *User) Query() (u []User) {
	dbWithCache.SetCache().Query(u,fmt.Sprintf("id = %d",user.Id))
	//db.Where("id = ?", user.Id).Find(&u)
	return u
}