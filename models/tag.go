package models

type Tag struct {
	// 文章分类
	Model
	Name string `json:"name"`
}

func GetTags(page int, size int, maps interface{}) (tags []Tag) {
	// 获取标签
	db.Where(maps).Offset(page).Limit(size).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	// 获取标签总数
	db.Model(Tag{}).Where(maps).Count(&count)
	return
}

func GetTag(maps interface{}) (tag Tag) {
	db.Where(maps).First(&tag)
	return
}
