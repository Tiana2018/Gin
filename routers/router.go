package routers

import (
	"Gin/pkg/setting"
	v1 "Gin/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("api/v1")
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tags", v1.AddTag)
		// 更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取单个文章
		apiv1.GET("/articles/:id",v1.GetArticle)
		//获取多个文章
		apiv1.GET("/articles",v1.GetArticles)
		//新增文章
		apiv1.POST("/articles",v1.AddArticle)
		//修改文章
		apiv1.PUT("/articles/:id",v1.EditArticle)
		//删除文章
		apiv1.DELETE("/articles/:id",v1.DeleteArticle)
	}

	return r
}
