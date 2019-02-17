package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	ID         int    `gorm:"primary_key" json:"id"`
	Title      int    `json:"title"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Category   string `json:"category"`
	Tag        string `json:"tag"`
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	// 文章创建
	_ = scope.SetColumn("CreateTime", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	// 文章更新
	_ = scope.SetColumn("UpdateTime", time.Now().Unix())

	return nil
}
