package test_autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
}

// InitArticleRouter 初始化 文章 路由信息
func (s *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	articleRouter := Router.Group("article").Use(middleware.OperationRecord())
	articleRouterWithoutRecord := Router.Group("article")
	var articleApi = v1.ApiGroupApp.Test_autocodeApiGroup.ArticleApi
	{
		articleRouter.POST("createArticle", articleApi.CreateArticle)             // 新建文章
		articleRouter.DELETE("deleteArticle", articleApi.DeleteArticle)           // 删除文章
		articleRouter.DELETE("deleteArticleByIds", articleApi.DeleteArticleByIds) // 批量删除文章
		articleRouter.PUT("updateArticle", articleApi.UpdateArticle)              // 更新文章
	}
	{
		articleRouterWithoutRecord.GET("findArticle", articleApi.FindArticle)       // 根据ID获取文章
		articleRouterWithoutRecord.GET("getArticleList", articleApi.GetArticleList) // 获取文章列表
	}
}
