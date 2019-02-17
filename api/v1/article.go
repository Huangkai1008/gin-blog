package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取文章列表
func GetArticles(c *gin.Context) {

	name := c.Query("name")

	maps := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	data := models.GetArticles(util.GetPage(c), setting.PageSize, maps)

	c.JSON(http.StatusOK, data)
}
