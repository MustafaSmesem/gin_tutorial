package controller

import (
	"fmt"
	"github.com/MustafaSmesem/gin_tutorial/model"
	"github.com/MustafaSmesem/gin_tutorial/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	GetAll() []model.Video
	GetById(id int) (model.Video, error)
	Save(ctx *gin.Context) model.Video
}

type controller struct {
	service service.VideoService
}

func (c *controller) GetAll() []model.Video {
	return c.service.FindAll()
}

func (c *controller) GetById(id int) (model.Video, error) {
	return c.service.FindById(id)
}

func (c *controller) Save(ctx *gin.Context) model.Video {
	var video model.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		fmt.Println("cannot get object body")
		return model.Video{}
	}
	return c.service.Save(video)
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}
