package main

import (
	"dousheng/pb/core"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	test := &core.User{
		Id:              1,
		Name:            "",
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		Avatar:          "",
		BackgroundImage: "",
		Signature:       "",
		TotalFavorited:  0,
		WorkCount:       0,
		FavoriteCount:   0,
	}
	//序列化
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("proto encode error: ", err)
		return
	}

	//反序列化
	newTest := &core.User{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("proto decode error: ", err)
		return
	}

	if test.GetId() != newTest.GetId() {
		log.Fatalf("data mismatch id %d != %d", test.GetId(), newTest.GetId())
		return
	} else {
		log.Println("data matched!")
	}

}
