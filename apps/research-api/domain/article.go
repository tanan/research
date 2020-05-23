package domain

import "time"

type ArticleId string

type Editor struct {
	Id   int
	Name string
	Icon string
}

type ThumbnailUrl string
type ThumbnailId string

type ArticleOverview struct {
	Title        string
	Editor       Editor
	LastModified time.Time
	Thumbnail    string
	Description  string
	//Tag Tag
}

type Article struct {
	ArticleId       ArticleId
	ArticleOverview ArticleOverview
	Content         map[string]interface{}
	Includes        map[string]interface{}
}

type Articles []Article
