package model

import (
	"fmt"

	"github.com/lexkong/log"

	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DataBase struct {
	Self *gorm.DB
}

// 声明DB为全局变量
var (
	DB *DataBase
)

func (db *DataBase) Init() {
	DB = &DataBase{
		Self: SetUpDB(viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.host"),
			viper.GetString("db.name")),
	}

	Migrate(DB.Self)
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
		log.Fatalf(err, "DB setup fail")
	}

	db.LogMode(viper.GetBool("gormlog"))
	// db.DB().SetMaxOpenConns(20000)
	db.DB().SetMaxIdleConns(0)

	return db
}

func (db *DataBase) CloseDB() {
	DB.Self.Close()
}
