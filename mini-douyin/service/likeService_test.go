package service

import (
	"log"
	"testing"
)

func TestLikeServiceImpl_FavoriteAction(t *testing.T) {
	err := GetLikeServImpInstance().FavoriteAction(7, 1, 1)
	if err != nil {
		log.Println(err)
	}
	log.Println("success")
}

func TestGetVideoLikedCount(t *testing.T) {
	likeCnt, err := GetLikeServImpInstance().GetVideoLikedCount(1)
	if err != nil {
		log.Default()
	}
	log.Println(likeCnt)
}

func TestGetUserLikeCount(t *testing.T) {
	likeCnt, err := GetLikeServImpInstance().GetUserLikeCount(1)
	if err != nil {
		log.Default()
	}
	log.Println(likeCnt)
}

func TestLikeServiceImpl_IsLikedByUser(t *testing.T) {
	liked, err := GetLikeServImpInstance().IsLikedByUser(7, 1)
	if err != nil {
		log.Default()
	}
	log.Println(liked)
}

func TestLikeServiceImpl_GetUserLikedCnt(t *testing.T) {
	count, err := GetLikeServImpInstance().GetUserLikedCnt(7)
	if err != nil {
		log.Default()
	}
	log.Println(count)
}

func TestLikeServiceImpl_GetLikesList(t *testing.T) {
	list, _ := GetLikeServImpInstance().GetLikesList(7)
	log.Println(list)
}
