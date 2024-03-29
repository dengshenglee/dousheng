package service

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestVideoServiceImpl_PublishList(t *testing.T) {
	videoList, err := GetVideoServiceInstance().PublishList(1)
	if err != nil {
		log.Default()
	}
	fmt.Println(videoList)
}

func TestVideoServiceImpl_Feed(t *testing.T) {
	videoList, nextTime, err := GetVideoServiceInstance().Feed(time.Now(), 1)
	if err != nil {
		log.Default()
	}
	fmt.Println(nextTime)
	fmt.Println(videoList)
}
