package controller

import (
	"mini-douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteActionResponse struct {
	Response
}

type GetFavouriteListResponse struct {
	Response
	VideoList []service.Video `json:"video_list"`
	// 81542800
}

func FavoriteAction(c *gin.Context) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	likeService := service.GetLikeServImpInstance()
	err := likeService.FavoriteAction(c.GetInt64("userId"), videoId, int32(actionType))
	if err == nil {
		c.JSON(http.StatusOK, FavoriteActionResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "favourite action success",
			},
		})
	} else {
		c.JSON(http.StatusOK, FavoriteActionResponse{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "favourite action fail",
			},
		})
	}
}

// FavoriteList 获取点赞列表
func FavoriteList(c *gin.Context) {
	strUserId := c.Query("user_id")
	//likeCnt:=dao.VideoLikedCount()
	userId, _ := strconv.ParseInt(strUserId, 10, 64)
	Fni := service.GetLikeServImpInstance()
	// 返回视频列表信息
	videoList, err := Fni.GetLikesList(userId)
	if err == nil {
		c.JSON(http.StatusOK, GetFavouriteListResponse{
			Response: Response{StatusCode: 0, StatusMsg: "get favouriteList success"}, VideoList: videoList,
		})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: -1, StatusMsg: "get favouriteList fail "})
	}
}
