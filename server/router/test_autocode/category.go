package test_autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
}

// InitCategoryRouter 初始化 分类 路由信息
func (s *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	categoryRouter := Router.Group("category").Use(middleware.OperationRecord())
	categoryRouterWithoutRecord := Router.Group("category")
	var categoryApi = v1.ApiGroupApp.Test_autocodeApiGroup.CategoryApi
	{
		categoryRouter.POST("createCategory", categoryApi.CreateCategory)             // 新建分类
		categoryRouter.DELETE("deleteCategory", categoryApi.DeleteCategory)           // 删除分类
		categoryRouter.DELETE("deleteCategoryByIds", categoryApi.DeleteCategoryByIds) // 批量删除分类
		categoryRouter.PUT("updateCategory", categoryApi.UpdateCategory)              // 更新分类
	}
	{
		categoryRouterWithoutRecord.GET("findCategory", categoryApi.FindCategory)       // 根据ID获取分类
		categoryRouterWithoutRecord.GET("getCategoryList", categoryApi.GetCategoryList) // 获取分类列表
	}
}
