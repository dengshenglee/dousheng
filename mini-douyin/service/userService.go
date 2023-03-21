package service

import "mini-douyin/dao"

type UserService interface {
	GetUserBasicInfoById(id int64) dao.UserBasicInfo
	GetUserBasicInfoByName(name string) dao.UserBasicInfo
	// get user info by id (not login)
	GetUserLoginInfoById(id int64) (User, error)
	// get user info by id (login)
	GetUserLoginInfoByIdWithCurId(id int64, curId int64) (User, error)
	InsertUser(user *dao.UserBasicInfo) bool
}

type User struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}
