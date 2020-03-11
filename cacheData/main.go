package cacheDataOrm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, _ := gorm.Open("mysql", "user:password@tcp(IP:port)/dbname?charset=utf8&parseTime=True&loc=Local")
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	defer db.Close()
}
