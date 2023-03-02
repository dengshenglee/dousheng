package model

import "gorm.io/gorm"

//	type User struct {
//		gorm.Model
//		Name            string `gorm:"type:varchar(20);not null; unique"`
//		Password        string `gorm:"type:varchar(20);not null"`
//		FollowCount     int    `gorm:"type:int;not null;default:0"`
//		FollowerCount   int64  `gorm:"type:int;not null;default:0"` // 粉丝总数
//		IsFollow        bool   `gorm:"-"`                           // true-已关注，false-未关注
//		Avatar          string `gorm:"type:varchar(32);not null"`   //用户头像
//		BackgroundImage string `gorm:"type:varchar(32);not null"`   //用户个人页顶部大图
//		Signature       string `gorm:"type:varchar(256);not null"`  //个人简介
//		TotalFavorited  int64  `gorm:"type:int;not null;default:0"` //获赞数量
//		WorkCount       int64  `gorm:"type:int;not null;default:0"` ///作品数量
//		FavoriteCount   int64  `gorm:"type:int;not null;default:0"` ///点赞数量
//	}
//
// Define User struct
type User struct {
	gorm.Model                // Include fields ID, CreatedAt, UpdatedAt, DeletedAt
	UserName        string    `gorm:"not null;unique"`
	UserPassword    string    `gorm:"not null" `
	FollowCount     int       `gorm:"type:int;not null;default:0"`
	FollowerCount   int64     `gorm:"type:int;not null;default:0"`     //粉丝总数
	IsFollow        bool      `gorm:"-"`                               //true-已关注，false-未关注
	Avatar          string    `gorm:"type:varchar(32);not null"`       //用户头像
	BackgroundImage string    `gorm:"type:varchar(32);not null"`       //用户个人页顶部大图
	Signature       string    `gorm:"type:varchar(256);not null"`      //个人简介
	Videos          []Video   `gorm:"foreignKey:AuthorID"  `           // One-to-many relationship with Video
	FavoriteVideos  []Video   `gorm:"many2many:user_favorite_video"  ` // Many-to-many relationship with Video for likes
	Comments        []Comment `gorm:"foreignKey:UserID" `              // One-to-many relationship with Comment
	Fllow_Users     []User    `gorm:"many2many:follow_follows;"`
	// Friends         []User    `gorm:"many2many:user_friends;"`
}

func GetUserByName(name string) (*User, error) {
	var user = new(User)
	var err error

	db, _ := GetDB()
	err = db.Where("user_name = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, err
}

func CreateUser(user *User) error {
	db, _ := GetDB()
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserById(id int) (*User, error) {
	var user = new(User)
	var err error

	db, _ := GetDB()
	err = db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, err
}
