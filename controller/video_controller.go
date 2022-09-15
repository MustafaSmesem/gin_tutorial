package controller

import (
	"github.com/MustafaSmesem/gin_tutorial/model"
	"github.com/MustafaSmesem/gin_tutorial/service"
	validators "github.com/MustafaSmesem/gin_tutorial/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
)

type VideoController interface {
	GetAll() []model.Video
	GetById(id int) (model.Video, error)
	Save(ctx *gin.Context) (model.Video, error)
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

func (c *controller) Save(ctx *gin.Context) (model.Video, error) {
	var video model.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return model.Video{}, err
	}
	err = validate.Struct(video)
	if err != nil {
		return model.Video{}, err
	}
	return c.service.Save(video), nil
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	err := validate.RegisterValidation("good-name", validators.ValidateGoodName)
	if err != nil {
		log.Fatalf("Validator cannot registered:%v", err)
	}
	return &controller{
		service: service,
	}
}
