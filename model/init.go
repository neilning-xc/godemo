package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 声明DB为全局变量
var Db *gorm.DB

func Init() {
	Db = SetUpDB("homestead", "password", "192.168.10.10:3306", "db_godemo")
	Migrate(Db)
}

func Migrate(db *gorm.DB) {
	// 创建user表
	db.AutoMigrate(&User{})
}

func SetUpDB(username, password, addr, name string) *gorm.DB {

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Fatalf("DB setup fail")
	}

	return db
}