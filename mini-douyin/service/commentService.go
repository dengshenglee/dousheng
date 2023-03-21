package service

import "mini-douyin/dao"

type CommentService interface {
	GetCommentCnt(videoId int64) (int64, error)
	CommentAction(comment dao.Comment) (Comment, error)
	DeleteComment(commentId int64) error
	GetCommentList(videoId int64, userId int64) ([]Comment, error)
}

type Comment struct {
	Id         int64  `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateAt   string `json:"create_date"`
	LikeCount  int64  `json:"like_count"`
	TeaseCount int64  `json:"tease_count"`
}
