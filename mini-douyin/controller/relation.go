package controller

import (
	"log"
	"mini-douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RelationActionResp struct {
	Response
}

type UserListResponse struct {
	Response
	UserList []service.User `json:"user_list"`
}

type FriendUserListResponse struct {
	Response
	FriendUserList []service.FriendUser `json:"user_list"`
}

func RelationAction(c *gin.Context) {
	userId := c.GetInt64("userId")
	toUserId, err1 := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, err2 := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err1 != nil || err2 != nil || actionType < 1 || actionType > 2 {
		c.JSON(http.StatusOK, RelationActionResp{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "请求参数格式错误",
			},
		})
		return
	}
	followService := service.GetFollowServiceInstance()
	var err error
	switch {
	case 1 == actionType:
		_, err = followService.FollowAction(userId, toUserId)
	case 2 == actionType:
		_, err = followService.CancelFollowAction(userId, toUserId)
	}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, RelationActionResp{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "服务器错误",
			},
		})
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "操作成功"})
}

func FollowList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "请求参数格式错误",
			},
		})
		return
	}
	followService := service.GetFollowServiceInstance()
	followList, err := followService.GetFollowings(userId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "服务器错误",
			},
		})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "操作成功",
		},
		UserList: followList,
	})
}

func FollowerList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "请求参数格式错误",
			},
		})
		return
	}
	followService := service.GetFollowServiceInstance()
	followList, err := followService.GetFollowers(userId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "获取粉丝列表失败",
			},
		})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "获取粉丝列表成功",
		},
		UserList: followList,
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusOK, FriendUserListResponse{
			Response{
				StatusCode: -1,
				StatusMsg:  "请求参数格式错误",
			},
			nil,
		})
		return
	}

	fsi := service.GetFollowServiceInstance()
	followers, err1 := fsi.GetFriends(userId)
	if err1 != nil {
		c.JSON(http.StatusOK, FriendUserListResponse{
			Response{
				StatusCode: -1,
				StatusMsg:  "获取好友列表失败",
			},
			nil,
		})
		return
	}

	c.JSON(http.StatusOK, FriendUserListResponse{
		Response{
			StatusCode: 0,
			StatusMsg:  "获取好友列表成功",
		},
		followers,
	})
}
