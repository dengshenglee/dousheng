package model

import "gorm.io/gorm"

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
	Likers        []User    `gorm:"many2many:user_video"` // Many-to-many relationship with User for likers
	Comments      []Comment `gorm:"foreignKey:CVideo"`    // One-to-many relationship with Comment
}
