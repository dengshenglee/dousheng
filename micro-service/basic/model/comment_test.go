package model

import "testing"

func TestCreateComment(t *testing.T) {
	err := CreateComment(&Comment{
		UserID:  1,
		VideoID: 1,
		Content: "test",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestGetCommentByVideo(t *testing.T) {
	video, err := GetVideoById(1)
	if err != nil {
		t.Fatal("get video err", err)
	}
	comments, err := GetCommentsByVideo(video)
	if err != nil {
		t.Fatal("get comments err", err)
	}
	for _, c := range comments {
		t.Log(c.Content)
	}
}

func TestGetCommentByUser(t *testing.T) {
	user, err := GetUserById(1)
	if err != nil {
		t.Fatal("get video err", err)
	}
	comments, err := GetCommentsByUser(user)
	if err != nil {
		t.Fatal("get comments err", err)
	}
	for _, c := range comments {
		t.Log(c.Content)
	}
}

func TestGetCommentByVideoAndUser(t *testing.T) {
	user, err := GetUserById(1)
	if err != nil {
		t.Fatal("get video err", err)
	}
	video, err := GetVideoById(1)
	if err != nil {
		t.Fatal("get video err", err)
	}
	comments, err := GetCommentByVideoAndUser(video, user)
	if err != nil {
		t.Fatal("get comments err", err)
	}
	if len(comments) == 0 {
		t.Log("No comments")
	} else {
		t.Log(comments[0].Content)
	}
}

func TestDeleteComment(t *testing.T) {
	user, err := GetUserById(1)
	if err != nil {
		t.Fatal("get video err", err)
	}
	video, err := GetVideoById(1)
	if err != nil {
		t.Fatal("get video err", err)
	}
	comment, err := GetCommentByVideoAndUser(video, user)
	if err != nil {
		t.Fatal("get comments err", err)
	}
	err = DeleteComment(&comment[0])
	if err != nil {
		t.Fatal("delete comment err", err)
	}
	t.Log("delete comment success")
}
