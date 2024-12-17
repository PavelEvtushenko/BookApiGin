package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//подключаемся к базе данных

var (
	db *gorm.DB
)

func Connect() {

	dsn := "root:GGGwwwppp@tcp(localhost:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
