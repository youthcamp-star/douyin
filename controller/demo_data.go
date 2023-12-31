package controller

import (
	"douyin/models"
	"time"
	"douyin/service"
	"gorm.io/gorm"
)

var DemoVideo = models.Video{
	Model: gorm.Model{
		ID: 1,
	},
	Title: "Bear",
	AuthorID: DemoUser.ID,
	PlayUrl:  "https://www.w3schools.com/html/movie.mp4",
	CoverUrl: "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
}
var DemoVideos = []models.Video{
	{
		Model: gorm.Model{
			ID: 1,
		},
		AuthorID: DemoUser.ID,
		PlayUrl:  "https://www.w3schools.com/html/movie.mp4",
		CoverUrl: "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
	},
}

var DemoComments = []models.Comment{
	{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Date(2023, 05, 01, 0, 0, 0, 0, time.Local),
		},
		Content: "Test Comment",
		UserId:  DemoUser.ID,
	},
}

var DemoUser = models.User{
	Model: gorm.Model{
		ID: 11,
	},
	Name: "TestUser",
	Password: "123456",
}

func Init(){
	service.CreateUser(&DemoUser)
}