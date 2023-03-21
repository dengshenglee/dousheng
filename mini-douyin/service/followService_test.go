package service

import (
	"fmt"
	"log"
	"testing"
)

func TestFollowServiceImp_GetFollowings(t *testing.T) {
	followings, err := GetFollowServiceInstance().GetFollowings(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(followings)
}

func TestFollowServiceImp_GetFollowers(t *testing.T) {
	followers, err := GetFollowServiceInstance().GetFollowers(7)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(followers)
}

func TestFollowServiceImp_GetFollowingCnt(t *testing.T) {

	userIdCnt, err := GetFollowServiceInstance().GetFollowingCnt(7)
	if err != nil {
		log.Default()
	}
	fmt.Println(userIdCnt)
}

func TestFollowServiceImp_GetFollowerCnt(t *testing.T) {

	userIdCnt, err := GetFollowServiceInstance().GetFollowerCnt(1)
	if err != nil {
		log.Default()
	}
	fmt.Println(userIdCnt)
}

func TestFollowServiceImp_CheckIsFollowing(t *testing.T) {
	result, err := GetFollowServiceInstance().CheckIsFollowing(1, 7)
	if err != nil {
		log.Default()
	}
	fmt.Println(result)
}

func TestFollowServiceImp_FollowAction(t *testing.T) {
	result, err := GetFollowServiceInstance().FollowAction(1, 7)
	if err != nil {
		log.Default()
	}
	fmt.Println(result)
}

func TestCancelFollow(t *testing.T) {
	result, err := GetFollowServiceInstance().CancelFollowAction(1, 7)
	if err != nil {
		log.Default()
	}
	fmt.Println(result)
}
