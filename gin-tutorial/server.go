package main

import (
	"gin-tutorial/controller"
	"gin-tutorial/service"
	"github.com/gin-gonic/gin"
)

var(
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "ok!!",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
