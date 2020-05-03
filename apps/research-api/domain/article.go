package domain

import "time"

type ArticleId string

type ArticleOverview struct {
	Title        string
	Editor       string
	LastModified time.Time
	Thumbnail    string
	Description  string
	//Tag Tag
}

type Article struct {
	ArticleId       ArticleId
	ArticleOverview ArticleOverview
	Content         map[string]interface{}
}

type Articles []Article
