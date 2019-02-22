package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Article struct {
	Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category int    `json:"category"`
	Tags     string `json:"tags"`
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return

}

func GetArticle(maps interface{}) (article Article) {
	db.Where(maps).First(&article)
	return
}

func AddArticle(title string, content string, category int, tags []int) bool {

	tagStr, err := json.Marshal(tags)
	if err != nil {
		log.Fatal(err)
	}

	db.Create(&Article{
		Title:    title,
		Content:  content,
		Category: category,
		Tags:     string(tagStr),
	})

	return true
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	// 文章创建
	_ = scope.SetColumn("CreateTime", time.Now())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	// 文章更新
	_ = scope.SetColumn("UpdateTime", time.Now())

	return nil
}
