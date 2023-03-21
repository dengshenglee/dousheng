package service

import (
	"errors"
	"log"
	"mini-douyin/config"
	"mini-douyin/dao"
	"sync"
)

type CommentServiceImpl struct {
	UserService
}

var (
	commentServiceImpl *CommentServiceImpl
	commentOnce        sync.Once
)

var _ CommentService = (*CommentServiceImpl)(nil)

func GetCommentServiceInstance() *CommentServiceImpl {
	commentOnce.Do(func() {
		commentServiceImpl = &CommentServiceImpl{
			UserService: GetUserServiceInstance(),
		}
	})
	return commentServiceImpl
}

// CommentAction implements CommentService
func (*CommentServiceImpl) CommentAction(comment dao.Comment) (Comment, error) {
	csi := GetCommentServiceInstance()
	commentRes, err := dao.InsertComment(comment)
	if err != nil {
		return Comment{}, err
	}
	user, err := csi.GetUserLoginInfoById(comment.UserId)
	if err != nil {
		log.Println("Err:", err.Error())
	}
	commentData := Comment{
		Id:         user.Id,
		User:       user,
		Content:    commentRes.Content,
		CreateAt:   commentRes.CreatedAt.Format(config.GO_STARTER_TIME),
		LikeCount:  0,
		TeaseCount: 0,
	}
	return commentData, nil
}

var err_not_found = errors.New("del comment is not exist")

// DeleteComment implements CommentService
func (*CommentServiceImpl) DeleteComment(commentId int64) error {
	err := dao.DeleteComment(commentId)
	if err != nil && !errors.Is(err, err_not_found) {
		log.Println("Err:", err.Error())
		return err
	}
	return nil
}

// GetCommentCnt implements CommentService
func (*CommentServiceImpl) GetCommentCnt(videoId int64) (int64, error) {
	cnt, err := dao.GetCommentCnt(videoId)
	if err != nil {
		log.Println("Err:", err.Error())
		return -1, err
	}
	return cnt, nil
}

// GetCommentList implements CommentService
func (*CommentServiceImpl) GetCommentList(videoId int64, userId int64) ([]Comment, error) {
	plainCommentList, err := dao.GetCommentList(videoId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	n := len(plainCommentList)
	if n == 0 {
		//no commemt in database
		return nil, nil
	}
	commentInfoList := make([]Comment, n, n)
	var wg sync.WaitGroup
	wg.Add(n)
	for i, comment := range plainCommentList {
		//TODO:why is this commentData not conflict with other goroutine
		var commentData Comment
		go func(index int, comment dao.Comment) {
			combineComment(&commentData, &comment)
			commentInfoList[index] = commentData
			wg.Done()
		}(i, comment)
	}
	wg.Wait()
	return commentInfoList, nil

}

func combineComment(comment *Comment, plainComment *dao.Comment) error {
	commentService := GetCommentServiceInstance()
	user, err := commentService.GetUserLoginInfoById(plainComment.UserId)
	if err != nil {
		log.Println("user database error", err)
		return err
	}
	comment.User = user
	comment.Id = plainComment.Id
	comment.Content = plainComment.Content
	comment.CreateAt = plainComment.CreatedAt.Format(config.GO_STARTER_TIME)
	comment.LikeCount = 0
	comment.TeaseCount = 0
	return nil
}
