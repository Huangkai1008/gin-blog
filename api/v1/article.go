package v1

import (
	"encoding/json"
	"gin-blog/models"
	"gin-blog/pickle"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
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
