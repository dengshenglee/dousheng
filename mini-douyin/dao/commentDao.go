package dao

import (
	"errors"
	"log"
	"time"
)

type Comment struct {
	Id         int64
	UserId     int64
	VideoId    int64
	Content    string
	ActionType int64 // 发布评论为1，取消评论为2
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Comment) TableName() string {
	return "comment"
}

func InsertComment(comment Comment) (Comment, error) {
	if err := Db.Model(Comment{}).Create(&comment).Error; err != nil {
		log.Println(err.Error())
		return Comment{}, err
	}
	return comment, nil
}

func DeleteComment(commentId int64) error {
	var comment Comment
	//check if comment exist
	result := Db.Where("id = ?", commentId).First(&comment)
	if result.Error != nil {
		return errors.New("del comment is not exist")
	}
	//delete comment, set action_type to 2
	result = Db.Model(Comment{}).Where("id=?", commentId).Update("action_type", 2)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCommentList(videoId int64) ([]Comment, error) {
	var commentList []Comment
	//use map[string]interface{} to pass multiple parameters
	result := Db.Model(Comment{}).Where(map[string]interface{}{"video_id": videoId, "action_type": 1}).Order("created_at desc").Find(&commentList)
	if result.Error != nil {
		log.Println(result.Error)
		return commentList, errors.New("get comment list failed")
	}
	return commentList, nil
}

func GetCommentListByUserId(userId int64) ([]Comment, error) {
	var commentList []Comment
	result := Db.Model(Comment{}).Where(map[string]interface{}{"user_id": userId, "action_type": 1}).Order("created_at desc").Find(&commentList)
	if result.Error != nil {
		log.Println(result.Error)
		return commentList, errors.New("get comment list failed")
	}
	return commentList, nil
}

func GetCommentCnt(videoId int64) (int64, error) {
	var cnt int64
	result := Db.Model(Comment{}).Where(map[string]interface{}{"video_id": videoId, "action_type": 1}).Count(&cnt)
	if result.Error != nil {
		log.Println(result.Error)
		return cnt, errors.New("get comment cnt failed")
	}
	return cnt, nil
}
