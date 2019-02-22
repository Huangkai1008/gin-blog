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

func ExistTag(params map[string]interface{}) bool {
	// 是否存在Tag
	var tag Tag
	name, existName := params["name"]

	if existName {
		db.Select("id").Where("name = ?", name).First(&tag)
	}
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string) bool {
	// 创建标签
	db.Create(&Tag{
		Name: name,
	})

	return true
}
