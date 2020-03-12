package cacheDataOrm

import "github.com/jinzhu/gorm"

var db *gorm.DB

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
	db.SingularTable(true)
}

func (user *User) Insert() {
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	db.Table("user").Create(user)
	//db.Table("user").Create(user)
}

func (user *User) Update() {
	user = &User{Id: user.Id, Name: "xiaoming"}
	db.Model(&user).Update(user)
}

func (user *User) Del() {
	db.Where("id =  ?", user.Id).Delete(&user)

}

func (user *User) Query() (u []User) {
	db.Where("id = ?", user.Id).Find(&u)
	return u
}