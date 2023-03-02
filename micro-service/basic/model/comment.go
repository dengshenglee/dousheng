package model

import (
	"gorm.io/gorm"
)

// type Comment struct {
// 	gorm.Model
// 	Content string `gorm:"type:text;not null"`
// 	User    User   `gorm:"-"`
// 	UserID  uint   `gorm:"not null"`
// 	V       Video  `gorm:"-"`
// 	VideoID uint   `gorm:"not null"`
// }

// Define Comment struct
type Comment struct {
	gorm.Model        // Include fields ID, CreatedAt, UpdatedAt, DeletedAt
	Content    string `gorm:"not null" `
	UserID     int    `gorm:"not null"`
	VideoID    int    `gorm:"not null"`
	User       User   `gorm:"-"` // Belongs-to relationship with User for commenter
	Video      Video  `gorm:"-"` // Belongs-to relationship with Video for commented video
}

func CreateComment(comment *Comment) error {
	db, _ := GetDB()
	err := db.Create(comment).Error
	if err != nil {
		return err
	}
	video, err := GetVideoById(uint(comment.VideoID))
	if err != nil {
		return err
	}
	err = UpdateVideoCommentCount(video)
	if err != nil {
		return err
	}

	return err
}

func GetCommentById(id int64) (*Comment, error) {
	var comment Comment
	var err error

	db, _ := GetDB()
	err = db.First(&comment, id).Error
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func DeleteComment(comment *Comment) error {
	db, _ := GetDB()
	err := db.Delete(comment).Error
	if err != nil {
		return err
	}
	video, err := GetVideoById(uint(comment.VideoID))
	if err != nil {
		return err
	}
	err = UpdateVideoCommentCount(video)
	if err != nil {
		return err
	}

	return nil
}

func GetCommentsByVideo(video *Video) ([]Comment, error) {
	db, _ := GetDB()
	var comments []Comment
	err := db.Where("video_id = ?", video.ID).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func GetCommentsByUser(user *User) ([]Comment, error) {
	db, _ := GetDB()
	var comments []Comment
	err := db.Where("user_id = ?", user.ID).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func GetCommentByVideoAndUser(video *Video, user *User) ([]Comment, error) {
	db, _ := GetDB()
	var comments []Comment
	err := db.Where("video_id = ? && user_id = ?", video.ID, user.ID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
