package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"world-map-scheduler-gin/controllers"
)

func Router(dbConn *gorm.DB) {
	handler := controllers.Handler{
		Db: dbConn,
	}

	router := gin.Default()

	// Boss一覧取得
	router.GET("/boss",  handler.GetAllBoss)

	router.Run(":3000")
}