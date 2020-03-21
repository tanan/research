package model

import "time"

type Article struct {
	ArticleId    string    `gorm:"column:id"`
	Title        string    `gorm:"column:title"`
	Description  string    `gorm:"column:description"`
	Editor       string    `gorm:"column:editor"`
	Thumbnail    string    `gorm:"column:thumbnail"`
	LastModified time.Time `gorm:"column:last_modified"`
}

func (a Article) TableName() string {
	return "research.article"
}
