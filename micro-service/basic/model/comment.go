package model

import "gorm.io/gorm"

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
	CContent   string `gorm:"not null" `
	CUser      int    `gorm:"not null"`
	CVideo     int    `gorm:"not null"`
	User       User   `gorm:"-"` // Belongs-to relationship with User for commenter
	Video      Video  `gorm:"-"` // Belongs-to relationship with Video for commented video
}
