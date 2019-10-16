package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"world-map-scheduler-gin/models"
)

type Handler struct {
	Db         *gorm.DB
}

type BossName struct {
	Id   int    `json:"boss_id"`
	Name string `json:"boss_name"`
}

func (h *Handler) GetAllBoss(c *gin.Context) {
	var bosses []models.Boss
	h.Db.Find(&bosses)

	restaurant := BossName{
		Id: 3,
		Name: "bbbbbb",
	}

	c.JSON(200, restaurant)
}
