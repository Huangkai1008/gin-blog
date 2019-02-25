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

	apiV1 := r.Group("/api/v1")
	{
		// 获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		// 新建标签
		apiV1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
		// 获取文章列表
		apiV1.GET("/articles", v1.GetArticles)
		// 获取指定id文章
		apiV1.GET("/articles/:id", v1.GetArticle)
		// 新建文章
		apiV1.POST("/articles", v1.AddArticle)
		// 修改文章
		apiV1.PUT("/articles/:id", v1.UpdateArticle)
		// 删除文章
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
