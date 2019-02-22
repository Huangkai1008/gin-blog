package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
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

//新增文章标签
func AddTag(c *gin.Context) {
	var tag models.Tag
	err := c.Bind(&tag)
	if err != nil {
		log.Fatal(err)
	}

	valid := validation.Validation{}
	valid.Required(tag.Name, "name").Message("标签名称不能为空")
	valid.MaxSize(tag.Name, 100, "name").Message("标签长度不能超过100")

	if !valid.HasErrors() {
		models.AddTag(tag.Name)
	}

	c.JSON(http.StatusOK, tag)

}

//修改文章标签
func EditTag(c *gin.Context) {

}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
