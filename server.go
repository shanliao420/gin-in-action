package main

import (
	"Gin_In_Action/handler"
	"Gin_In_Action/resposity"
	"Gin_In_Action/util"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	if err := Init(); err != nil {
		os.Exit(-1)
	}

	router := gin.Default()

	router.Use(gin.Logger())

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/community/page/get/:id", func(context *gin.Context) {
		topicId := context.Param("id")
		data := handler.QueryPageInfo(topicId)
		context.JSON(200, data)
	})

	router.POST("/community/post/do", func(context *gin.Context) {
		uid, _ := context.GetPostForm("uid")
		topicId, _ := context.GetPostForm("topic_id")
		content, _ := context.GetPostForm("content")

		data := handler.PublishPost(uid, topicId, content)

		context.JSON(200, data)
	})

	err := router.Run()
	if err != nil {
		return
	}
}

func Init() error {
	if err := resposity.Init(); err != nil {
		return err
	}
	if err := util.InitLogger(); err != nil {
		return err
	}
	return nil
}
