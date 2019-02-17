package routers

import (
	"gin-blog/api/v1"
	"gin-blog/middleware"
	"gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
	}

	return r
}
