package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	PostID uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Konten string    `gorm:"not null" json:"konten"`
}

func (Post) TableName() string {
	return "post"
}

func (u *Post) BeforeCreate(tx *gorm.DB) (err error) {
	u.PostID = uuid.New()
	return
}
