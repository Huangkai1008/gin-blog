package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
标签tag api部分
*/
// 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	data["tags"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, data)

}

// 获得单个文章标签
func GetTag(c *gin.Context) {
	id := c.Query("id")
	maps := make(map[string]interface{})

	maps["id"] = id

}

//新增文章标签
func AddTag(c *gin.Context) {

}

//修改文章标签
func EditTag(c *gin.Context) {

}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
