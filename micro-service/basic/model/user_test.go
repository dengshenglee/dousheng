package model

import (
	"testing"

	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	err := CreateUser(&User{
		Model:           gorm.Model{},
		UserName:        "Adam",
		UserPassword:    "123456",
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		Avatar:          "xxx",
		BackgroundImage: "xxx",
		Signature:       "xxx",
		Videos:          []Video{},
		FavoriteVideos:  []Video{},
		Comments:        []Comment{},
		Fllow_Users:     []User{},
	})
	if err != nil {
		t.Log(err)
	}
	t.Log("Success")
}

func TestGetUserByName(t *testing.T) {
	user, err := GetUserByName("Adam")
	if err != nil {
		t.Log(err)
	}
	t.Logf("Get Success of No.%d %s\n", user.ID, user.UserName)
}

func TestGetUserById(t *testing.T) {
	user, err := GetUserById(1)
	if err != nil {
		t.Log(err)
	}
	t.Logf("Get Success of No.%d %s\n", user.ID, user.UserName)
}
