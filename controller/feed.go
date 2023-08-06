package controller

import (
	"douyin/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type FeedResponse struct {
	Response
	VideoList []models.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})

}
