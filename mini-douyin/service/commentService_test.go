package service

import (
	"fmt"
	"mini-douyin/dao"
	"testing"
)

func TestCommentServiceImpl_GetCommentCnt(t *testing.T) {
	count, err := GetCommentServiceInstance().GetCommentCnt(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(count)
}

func TestCommentServiceImpl_GetCommentList(t *testing.T) {
	list, err := GetCommentServiceInstance().GetCommentList(1, 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(list)
}

func TestCommentServiceImpl_DeleteComment(t *testing.T) {
	err := GetCommentServiceInstance().DeleteComment(8)
	if err != nil {
		t.Error(err)
	}
}

func TestCommentServiceImpl_CommentAction(t *testing.T) {
	comment := dao.Comment{
		UserId:     7,
		VideoId:    3,
		Content:    "test",
		ActionType: 1,
	}
	commentRes, err := GetCommentServiceInstance().CommentAction(comment)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(commentRes)
}
