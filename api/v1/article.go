package v1

import (
	"encoding/json"
	"fmt"
	"gin-blog/models"
	"gin-blog/pickle"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"github.com/Unknwon/com"
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

	maps := make(map[string]interface{})

	maps["id"] = id

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

	fmt.Println(data)

	c.JSON(http.StatusOK, data)
}
