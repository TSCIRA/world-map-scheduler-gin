package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// viewディレクトリを指定
	router.LoadHTMLGlob("views/*")

	router.GET("/", func(c *gin.Context) {
		ctrl := task.NewTask()

		// ControllerのGetAllメソッドを使って全てのタスクを取得
		tasks := ctrl.GetAll()

		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"tasks": tasks,
		})
	})

	router.POST("/", func(c *gin.Context) {
		text := c.PostForm("text")

		ctrl := task.NewTask()

		// ControllerのCreateメソッドを使ってタスクを新規作成
		ctrl.Create(text)

		tasks := ctrl.GetAll()

		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"tasks": tasks,
		})
	})

	router.Run(":8080")
}