package entity

type Post struct {
	PostID int8   `gorm:"primary_key;auto_increment" json:"id"`
	Konten string `gorm:"not null" json:"konten"`
}

func (Post) TableName() string {
	return "post"
}
