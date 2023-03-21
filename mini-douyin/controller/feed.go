package controller

import (
	"math"
	"mini-douyin/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []service.Video `json:"video_list"`
	NextTime  int64           `json:"next_time"`
}

func Feed(c *gin.Context) {
	latestTime := c.Query("latest_time")
	var convTime time.Time
	if latestTime != "0" {
		t, _ := strconv.ParseInt(latestTime, 10, 64)
		if t > math.MaxInt32 {
			convTime = time.Now()
		} else {
			convTime = time.Unix(t, 0)
		}
	} else {
		convTime = time.Now()
	}

	userId := c.GetInt64("userId")
	videoService := service.GetVideoServiceInstance()
	videos, nextTime, err := videoService.Feed(convTime, userId)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 1, StatusMsg: "刷新视频流失败"},
			VideoList: nil,
			NextTime:  nextTime.Unix(),
		})
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "刷新视频流成功!"},
		VideoList: videos,
		NextTime:  nextTime.Unix(),
	})
}
