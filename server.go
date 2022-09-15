package main

import (
	"github.com/MustafaSmesem/gin_tutorial/controller"
	"github.com/MustafaSmesem/gin_tutorial/middlewares"
	"github.com/MustafaSmesem/gin_tutorial/service"
	"github.com/gin-gonic/gin"
	ginDump "github.com/tpkeeper/gin-dump"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const PORT = ":1234"

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
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), ginDump.Dump())

	server.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello world from gin!!",
		})
	})
	server.GET("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.GetAll())
	})
	server.GET("/video/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		video, err := videoController.GetById(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(200, video)
		}
	})
	server.POST("/video", func(context *gin.Context) {
		video, err := videoController.Save(context)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(200, video)
		}
	})

	err := server.Run(PORT)
	if err != nil {
		log.Fatal(err)
	}
}
