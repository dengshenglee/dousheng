package dao

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestInsertLikeInfo(t *testing.T) {
	err := InsertLikeInfo(Like{
		UserId:    1,
		VideoId:   1,
		Liked:     1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		t.Error(err)
	}
}
func TestUpdateLikeInfo(t *testing.T) {
	err := UpdateLikeInfo(1, 1, 1)
	if err != nil {
		log.Println(err)
	}
}

func TestGetLikeListByUserId(t *testing.T) {
	list, err := GetLikeListByUserId(1)
	if err != nil {
		log.Print(err.Error())
	}
	for _, v := range list {
		fmt.Printf("%d\n", v)
	}
}

func TestVideoLikedCount(t *testing.T) {
	likeCnt, err := VideoLikedCount(1)
	if err != nil {
		log.Print(err.Error())
	}
	fmt.Printf("Like Count：%d\n", likeCnt)
}

func TestGetLikeCountByUser(t *testing.T) {
	likeCnt, err := GetLikeCountByUser(1)
	if err != nil {
		log.Print(err.Error())
	}
	fmt.Printf("Like Count：%d\n", likeCnt)
}

func TestIsLikedByUser(t *testing.T) {
	flag, err := IsLikedByUser(1, 1)
	if err != nil {
		log.Default()
	}
	log.Println(flag)
}

func TestGetUserVideoLikedByOther(t *testing.T) {
	likedList, err := GetUserVideoLikedByOther(1)
	if err != nil {
		log.Default()
	}
	log.Println(likedList)
}
