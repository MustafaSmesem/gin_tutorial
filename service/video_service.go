package service

import (
	"fmt"
	"github.com/MustafaSmesem/gin_tutorial/model"
	"github.com/MustafaSmesem/gin_tutorial/util"
)

type VideoService interface {
	Save(video model.Video) model.Video
	FindAll() []model.Video
	FindById(id int) (model.Video, error)
}

type videoService struct {
	videos []model.Video
	lastId int
}

func New() VideoService {
	return &videoService{lastId: 0}
}

func (vs *videoService) Save(video model.Video) model.Video {
	video.Id = vs.lastId
	vs.lastId = vs.lastId + 1
	vs.videos = append(vs.videos, video)
	fmt.Println(vs.videos)
	return video
}
func (vs *videoService) FindAll() []model.Video {
	return vs.videos
}

func (vs *videoService) FindById(id int) (model.Video, error) {
	return util.FindVideoById(vs.videos, id)
}
