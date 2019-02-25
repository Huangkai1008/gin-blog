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

func ExistArticle(params map[string]interface{}) bool {
	// 是否存在文章
	var article Article
	name, existName := params["name"]
	if existName {
		db.Select("name").Where("name = ?", name).First(&article)
	}

	if article.ID > 0 {
		return true
	}

	id, existId := params["id"]
	if existId {
		db.Select("id").Where("id = ?", id).First(&article)
	}

	if article.ID > 0 {
		return true
	}

	return false

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

func UpdateArticle(id int, data interface{}) bool {
	// 修改文章
	db.Model(&Tag{}).Where("id = ?", id).Update(data)
	return true
}

func DeleteArticle(id int) bool {
	// 删除文章
	db.Where("id = ?", id).Delete(&Tag{})
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
