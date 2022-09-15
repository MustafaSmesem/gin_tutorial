package model

type Person struct {
	FirstName string `json:"first_name" binding:"required" validate:"good-name"`
	LastName  string `json:"last_name" binding:"required"`
	Age       int    `json:"age" binding:"lte=130,gte=15"`
	Email     string `json:"email" binding:"required,email"`
}

type Video struct {
	Title    string `json:"title" binding:"min=10,max=200"`
	Url      string `json:"url" binding:"required,url"`
	Duration int    `json:"duration"`
	Name     string `json:"name"`
	Id       int    `json:"id"`
	Author   Person `json:"author"`
}
