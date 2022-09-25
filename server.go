package main

import (
	"github.com/MustafaSmesem/gin_tutorial/controller"
	"github.com/MustafaSmesem/gin_tutorial/middlewares"
	"github.com/MustafaSmesem/gin_tutorial/service"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	videoService    = service.New()
	videoController = controller.New(videoService)
)

func setupLogOutput() {
	f, err := os.Create("gin.log")
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(context *gin.Context) {
			context.JSON(200, videoController.GetAll())
		})
		apiRoutes.GET("/video/:id", func(context *gin.Context) {
			id, _ := strconv.Atoi(context.Param("id"))
			video, err := videoController.GetById(id)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(200, video)
			}
		})
		apiRoutes.POST("/video", func(context *gin.Context) {
			video, err := videoController.Save(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(200, video)
			}
		})
	}

	viewRoutes := server.Group("/")
	{
		viewRoutes.GET("", func(context *gin.Context) {
			context.HTML(http.StatusOK, "videos.page.tmpl", gin.H{"title": "Home Page"})
		})

		viewRoutes.GET("videos", videoController.ShowAll)
	}

	server.Static("/css", "./public/css")
	server.LoadHTMLGlob("./public/templates/*.tmpl")

	var port = os.Getenv("PORT")
	if port == "" {
		port = "3200"
	}
	err := server.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
