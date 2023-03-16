package dao

import (
	"log"
	"testing"
)

func TestInsertComment(t *testing.T) {
	var comment Comment = Comment{
		UserId:     1,
		VideoId:    1,
		ActionType: 1,
		Content:    "这条评论来自单元测试A",
	}
	commentRes, err := InsertComment(comment)
	if err != nil {
		log.Println(err)
	}
	log.Println("返回评论", commentRes)
}

func TestDeleteComment(t *testing.T) {
	err := DeleteComment(5)
	if err == nil {
		log.Println("delete comment success")
	}
}

func TestGetCommentList(t *testing.T) {
	commentList, err := GetCommentList(1)
	if err == nil {
		log.Println(commentList)
	}
}

func TestGetCommentCnt(t *testing.T) {
	count, err := GetCommentCnt(1)
	if err == nil {
		log.Println(count)
	}
}
