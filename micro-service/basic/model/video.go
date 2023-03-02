package model

import (
	"time"

	"gorm.io/gorm"
)

// type Video struct {
// 	gorm.Model
// 	AuthorID      uint   `gorm:"index"`
// 	Author        User   `gorm:"-"`
// 	PlayUrl       string `gorm:"type:varchar(255);not null"`
// 	CoverUrl      string `gorm:"type:varchar(255);not null"`
// 	FavoriteCount int    `gorm:"type:int(11);not null;default:0"`
// 	CommentCount  int    `gorm:"type:int(11);not null;default:0"`
// 	IsFavorite    bool   `gorm:"-"`
// 	Title         string `gorm:"type:varchar(255);not null"`
// }

// Define Video struct type Video struct { gorm.Model // Include fields ID, CreatedAt, UpdatedAt, DeletedAt VID int gorm:"primaryKey" VTitle string gorm:"not null" VDescription string VUrl string gorm:"not null" VPublisher int Publisher User gorm:"references:UserID" // Belongs-to relationship with User for publisher Likers []User gorm:"many2many:user_video" // Many-to-many relationship with User for likers Comments []Comment gorm:"foreignKey:CVideo" // One-to-many relationship with Comment }

// Define Comment struct type Comment struct { gorm.Model // Include fields ID, CreatedAt, UpdatedAt, DeletedAt CID int gorm:"primaryKey" CContent string gorm:"not null" CUser int CVideo int User User gorm:"references:UserID" // Belongs-to relationship with User for commenter Video Video gorm:"references:VID" // Belongs-to relationship with Video for commented video }

// Define Video struct
type Video struct {
	gorm.Model              // Include fields ID, CreatedAt, UpdatedAt, DeletedAt
	VTitle        string    `gorm:"not null"`
	VUrl          string    `gorm:"not null" `
	CoverUrl      string    `gorm:"type:varchar(255);not null"`
	FavoriteCount int       `gorm:"type:int(11);not null;default:0"`
	CommentCount  int       `gorm:"type:int(11);not null;default:0"`
	IsFavorite    bool      `gorm:"-"`
	Author        User      `gorm:"-"`
	AuthorID      uint      `gorm:"index"`
	Likers        []User    `gorm:"many2many:user_favorite_video"` // Many-to-many relationship with User for likers
	Comments      []Comment `gorm:"foreignKey:VideoID"`            // One-to-many relationship with Comment
}

func CreateVideo(video *Video) error {
	var err error
	db, _ := GetDB()
	err = db.Create(video).Error
	if err != nil {
		return err
	}
	return nil

}

func GetVideoOrderByTime(time time.Time) ([]Video, error) {
	db, _ := GetDB()
	var videos []Video
	// get all videos
	err := db.Where("created_at < ?", time).Order("created_at desc").Limit(30).Find(&videos).Error
	return videos, err
}

func GetVideoAfterTime(time time.Time) ([]Video, error) {
	db, _ := GetDB()
	var videos []Video
	// get all videos
	err := db.Where("created_at > ?", time).Order("created_at desc").Limit(30).Find(&videos).Error
	return videos, err

}

func GetVideoById(id uint) (*Video, error) {
	var video Video
	db, _ := GetDB()
	// get video
	err := db.Where("id = ?", id).First(&video).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}

func GetVideosByUser(user *User) ([]Video, error) {
	var videos []Video
	db, _ := GetDB()
	err := db.Model(user).Association("Videos").Find(&videos)
	return videos, err
}

func AddFavoriteVideo(user *User, video *Video) error {
	var err error
	db, _ := GetDB()
	err = db.Model(user).Association("FavoriteVideos").Append(video)
	if err != nil {
		return err
	}
	err = UpdateVideoFavoriteCount(video)
	return err
}

func UnFavoriteVideo(user *User, video *Video) error {
	var err error
	db, _ := GetDB()
	err = db.Model(user).Association("FavoriteVideos").Delete(video)
	if err != nil {
		return err
	}
	err = UpdateVideoFavoriteCount(video)
	return err
}

func GetUserFavoriteVideos(user *User) ([]Video, error) {
	var videos []Video
	db, _ := GetDB()
	err := db.Model(user).Association("FavoriteVideos").Find(&videos)
	return videos, err
}

func UpdateVideoFavoriteCount(video *Video) error {
	db, _ := GetDB()
	num_favorite := db.Model(video).Association("Likers").Count()
	return db.Model(video).Update("FavoriteCount", num_favorite).Error
}

func IsFavoriteVideo(user *User, video *Video) bool {
	db, _ := GetDB()
	var u User
	db.Model(video).Association("Likers").Find(&u, user.ID)
	return u.ID != 0
}

func UpdateVideoCommentCount(video *Video) error {
	db, _ := GetDB()
	num_comment := db.Model(video).Association("Comments").Count()
	return db.Model(video).Update("comment_count", num_comment).Error
}
