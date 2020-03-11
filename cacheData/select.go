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
}

func (user *User) Update() {
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	db.Table("user").Create(user)
}

func (user *User) Del() {
	db.Delete(&user)

	//delete from user where id > 11;
	db.Delete(&User{}, "id > ?", 11)
}

func (user *User) query() (u []User) {
	db.Find(&u)

	db.Find(&u, "id > ? and age > ?", 2, 12)

	db.Where("id > ?", 2).Find(&u)

	db.Where("name in (?)", []string{"xiaoming", "xiaohong"}).Find(&u)

	db.First(&u)

	db.First(&u, "where sex = ?", 1)

	db.Last(&u)

	return u
}
