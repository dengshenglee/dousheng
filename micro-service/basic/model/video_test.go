package model

import (
	"testing"
	"time"
)

func TestCreateVideo(t *testing.T) {
	err := CreateVideo(&Video{
		VTitle:        "hahaha",
		VUrl:          "xxx",
		CoverUrl:      "xxx",
		FavoriteCount: 0,
		CommentCount:  0,
		AuthorID:      1,
	})
	if err != nil {
		t.Fatal("create err", err)
	}
	t.Log("create success!")
}

func TestGetVideoById(t *testing.T) {
	v, err := GetVideoById(1)
	if err != nil {
		t.Fatal("get by id err", err)
	}
	t.Logf("get success with %s\n", v.VTitle)
}

func TestGetVideoOrderByTime(t *testing.T) {
	videos, err := GetVideoOrderByTime(time.Now())
	if err != nil {
		t.Fatal("get by time err", err)
	}
	for _, v := range videos {
		t.Logf("No.%d video's title:%s\n", v.ID, v.VTitle)
	}
}

func TestGetVideoAfterTime(t *testing.T) {
	videos, err := GetVideoAfterTime(time.Now())
	if err != nil {
		t.Fatal("get by time err", err)
	}
	for _, v := range videos {
		t.Logf("No.%d video's title:%s\n", v.ID, v.VTitle)
	}
}
func TestGetVideoByUser(t *testing.T) {
	user, err := GetUserById(1)
	if err != nil {
		t.Fatal("get by user err", err)
	}
	v, err := GetVideosByUser(user)
	if err != nil {
		t.Fatal("get by user err", err)
	}
	t.Logf("get success with %d videos\n", len(v))
}

func TestAddFavorite(t *testing.T) {
	user, err := GetUserById(1)
	if err != nil {
		t.Fatal("get by user err", err)
	}
	for _, v := range user.FavoriteVideos {
		t.Log("before: favorite video", v.VTitle)
	}
	video, err := GetVideoById(3)
	if err != nil {
		t.Fatal("get by video err", err)
	}
	err = AddFavoriteVideo(user, video)
	if err != nil {
		t.Fatal("add favorite video err", err)
	}
	for _, v := range user.FavoriteVideos {
		t.Log("after: favorite video", v.VTitle)
	}
}

func TestUnFavoriteVideo(t *testing.T) {
	user, err := GetUserById(1)
	if err != nil {
		t.Fatal("get by user err", err)
	}
	for _, v := range user.FavoriteVideos {
		t.Log("before: favorite video", v.VTitle)
	}
	video, err := GetVideoById(1)
	if err != nil {
		t.Fatal("get by video err", err)
	}
	err = UnFavoriteVideo(user, video)
	if err != nil {
		t.Fatal("unfavorite video err", err)
	}
	for _, v := range user.FavoriteVideos {
		t.Log("after: favorite video", v.VTitle)
	}
}

func TestGetUserFavoriteVideos(t *testing.T) {
	user, err := GetUserById(1)
	if err != nil {
		t.Fatal("get by user err", err)
	}
	videos, err := GetUserFavoriteVideos(user)
	if err != nil {
		t.Fatal("get by user favorite videos err", err)
	}
	for _, v := range videos {
		t.Log("favorite video", v.VTitle)
	}
}

func TestIsFavoriteVideo(t *testing.T) {
	user, err := GetUserById(1)
	if err != nil {
		t.Fatal("get by user err", err)
	}
	video, err := GetVideoById(1)
	if err != nil {
		t.Fatal("get by video err", err)
	}
	if IsFavoriteVideo(user, video) {
		t.Log("is favorite video")
	} else {
		t.Fatal("not favorite video")
	}
}
