package util

import (
	"errors"
	"fmt"
	"github.com/MustafaSmesem/gin_tutorial/model"
)

func FindVideoById(slice []model.Video, id int) (model.Video, error) {
	for _, element := range slice {
		if element.Id == id {
			return element, nil
		}
	}
	return model.Video{}, errors.New(fmt.Sprintf("Video with id %v doesn't exist", id))
}
