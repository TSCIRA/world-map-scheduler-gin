package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"world-map-scheduler-gin/models"
)

func Init() *gorm.DB {
	DBMS     := "mysql"
	USER     := "root"
	PASS     := ""
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME   := "world_map_scheduler_db"
	OPTION 	 := "charset=utf8mb4&parseTime=True&loc=Local"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		fmt.Println(err)
		panic("データベースへの接続に失敗しました")
	}
	db.LogMode(true)
	db.AutoMigrate(&models.Boss{})

	return db
}