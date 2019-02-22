package pickle

import "time"

type ArticleJson struct {
	// 文章json
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Category   int       `json:"category"`
	Tags       []int     `json:"tags"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
