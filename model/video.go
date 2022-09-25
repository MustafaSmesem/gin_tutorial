package model

import (
	"fmt"
	"strings"
)

type Person struct {
	FirstName string `json:"first_name" binding:"required" validate:"good-name"`
	LastName  string `json:"last_name" binding:"required"`
	Age       int    `json:"age" binding:"lte=130,gte=15"`
	Email     string `json:"email" binding:"required,email"`
}

type Video struct {
	Title       string `json:"title" binding:"min=10,max=200"`
	Url         string `json:"url" binding:"required,url"`
	Duration    int    `json:"duration"`
	Description string `json:"description" binding:"max=500"`
	Name        string `json:"name"`
	Id          int    `json:"id"`
	Author      Person `json:"author"`
}

func (p Person) Name() string {
	return p.FirstName + " " + strings.ToUpper(p.LastName)
}

func (v Video) Time() string {
	var h, m, s int
	s = v.Duration
	if s/60 > 0 {
		m = s / 60
		s = s % 60
	}
	if m/60 > 0 {
		h = m / 60
		m = m % 60
	}
	hs := ""
	if h > 0 {
		hs = fmt.Sprintf("%v hours, ", h)
	}
	return fmt.Sprintf("%v%v minutes and %v seconds", hs, m, s)
}
