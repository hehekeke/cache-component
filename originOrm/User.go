package originOrm

import (
	cacheDataOrm "cache-component/cacheData"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

var dbWithCache cacheDataOrm.DBWithCache


type User struct {
	Id    int
	Name  string
	Age   int
	Sex   byte
	Phone string
}

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//设置全局表名禁用复数
	dbWithCache.DB = db
	db.SingularTable(true)
}

/**
增
 */
func (user *User) Insert() {
	db.Table("user").Create(user)
}
/**
增 cache
*/
func (user *User) InsertWithCache() {
	dbWithCache.SetCache().Table("user").Create(user)
}
/**
更新  update
 */
func (user *User) Update() {
	user = &User{Id: user.Id, Name: "xiaoming"}
	db.Model(&user).Table("user").Update(user)
}

func (user *User) UpdateWithCache() {
	user = &User{Id: user.Id, Name: "xiaoming"}
	dbWithCache.SetCache().Table("user").Update(user)
}
/**
删除
*/
func (user *User) Del() {
	db.Delete(&user,"id =  ?", user.Id)

}

func (user *User) DelWithCache() {
	dbWithCache.SetCache().Table("user").Delete(&user,"id =  ?", user.Id)

}
/**
查找
*/
func (user *User) Query() (u []User) {
	db.Where("id = ?", user.Id).Find(&u)
	return u
}

func (user *User) QueryWithCache() (u []User) {
	dbWithCache.SetCache().Table("user").Query(&u,"id =  ?", user.Id)
	return u
}
