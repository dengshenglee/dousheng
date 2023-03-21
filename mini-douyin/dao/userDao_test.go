package dao

import (
	"log"
	"testing"
)

func TestInsertUser(t *testing.T) {
	user := UserBasicInfo{Name: "ut2", Password: "unit test"}
	flag := InsertUser(&user)
	log.Println(flag)
}

func TestGetUserBasicInfoById(t *testing.T) {
	res, err := GetUserBasicInfoById(1)
	if err == nil {
		log.Println(res)
	}
	log.Println(res.CreatedAt)
}

func TestGetUserBasicInfoByName(t *testing.T) {
	res, err := GetUserBasicInfoByName("ut")
	if err == nil {
		log.Println(res)
	}
}
