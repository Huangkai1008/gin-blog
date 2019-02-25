package v1

import (
	"encoding/json"
	"gin-blog/models"
	"gin-blog/pickle"
	"gin-blog/pkg/e"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 获取文章列表
func GetArticles(c *gin.Context) {

	name := c.Query("name")

	maps := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	articles := models.GetArticles(util.GetPage(c), setting.PageSize, maps)

	var data []pickle.ArticleJson
	var tags []int

	for _, a := range articles {
		err := json.Unmarshal([]byte(a.Tags), &tags)

		if err != nil {
			log.Fatal(err)
		}

		data = append(data, pickle.ArticleJson{
			ID:         a.ID,
			Title:      a.Title,
			Content:    a.Content,
			Category:   a.Category,
			CreateTime: a.CreateTime,
			UpdateTime: a.UpdateTime,
			Tags:       tags})
	}

	c.JSON(http.StatusOK, data)
}

// 获取文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if id == 0 {
		id = 1
	}

	maps := map[string]interface{}{
		"id": id,
	}

	article := models.GetArticle(maps)

	var data pickle.ArticleJson
	var tags []int

	err := json.Unmarshal([]byte(article.Tags), &tags)

	if err != nil {
		log.Fatal(err)
	}
	data = pickle.ArticleJson{
		ID:         article.ID,
		Title:      article.Title,
		Content:    article.Content,
		Category:   article.Category,
		CreateTime: article.CreateTime,
		UpdateTime: article.UpdateTime,
		Tags:       tags}

	c.JSON(http.StatusOK, data)
}

// 新增文章
func AddArticle(c *gin.Context) {
	var articleJson pickle.ArticleJson

	err := c.Bind(&articleJson)
	if err != nil {
		log.Fatal(err)
	}

	valid := validation.Validation{}
	valid.Required(articleJson.Title, "title").Message("文章标题不能为空")
	valid.Required(articleJson.Content, "content").Message("文章内容不能为空")
	valid.MaxSize(articleJson.Title, 100, "title").Message("文章标题不能超过100个字符")

	if !valid.HasErrors() {
		models.AddArticle(
			articleJson.Title,
			articleJson.Content,
			articleJson.Category,
			articleJson.Tags)
	}

	c.JSON(http.StatusOK, articleJson)
}

// 修改文章内容
func UpdateArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	var article models.Article

	err := c.BindJSON(&article)
	if err != nil {
		log.Fatal(err)
	}

	valid := validation.Validation{}

	if !valid.HasErrors() {
		params := map[string]interface{}{
			"id": id,
		}

		if !models.ExistArticle(params) {
			models.UpdateArticle(id, article)
		} else {
			c.JSON(e.ERROR_NOT_EXIST_ARTICLE, article)
		}
	}

	c.JSON(http.StatusOK, article)
}

// 删除文章内容
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id最小不能为0")

	if !valid.HasErrors() {
		params := map[string]interface{}{
			"id": id,
		}

		if !models.ExistArticle(params) {
			models.DeleteArticle(id)
		} else {
			c.JSON(e.ERROR_NOT_EXIST_ARTICLE, id)
		}
	}
	c.JSON(http.StatusOK, id)
}
