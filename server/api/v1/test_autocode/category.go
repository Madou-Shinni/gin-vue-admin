package test_autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/test_autocode"
	test_autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/test_autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryApi struct {
}

var categoryService = service.ServiceGroupApp.Test_autocodeServiceGroup.CategoryService

// CreateCategory 创建分类
// @Tags Category
// @Summary 创建分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test_autocode.Category true "创建分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /category/createCategory [post]
func (categoryApi *CategoryApi) CreateCategory(c *gin.Context) {
	var category test_autocode.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := categoryService.CreateCategory(&category); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCategory 删除分类
// @Tags Category
// @Summary 删除分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test_autocode.Category true "删除分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /category/deleteCategory [delete]
func (categoryApi *CategoryApi) DeleteCategory(c *gin.Context) {
	ID := c.Query("ID")
	if err := categoryService.DeleteCategory(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCategoryByIds 批量删除分类
// @Tags Category
// @Summary 批量删除分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /category/deleteCategoryByIds [delete]
func (categoryApi *CategoryApi) DeleteCategoryByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := categoryService.DeleteCategoryByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCategory 更新分类
// @Tags Category
// @Summary 更新分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test_autocode.Category true "更新分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /category/updateCategory [put]
func (categoryApi *CategoryApi) UpdateCategory(c *gin.Context) {
	var category test_autocode.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := categoryService.UpdateCategory(category); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCategory 用id查询分类
// @Tags Category
// @Summary 用id查询分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query test_autocode.Category true "用id查询分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /category/findCategory [get]
func (categoryApi *CategoryApi) FindCategory(c *gin.Context) {
	ID := c.Query("ID")
	if recategory, err := categoryService.GetCategory(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recategory": recategory}, c)
	}
}

// GetCategoryList 分页获取分类列表
// @Tags Category
// @Summary 分页获取分类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query test_autocodeReq.CategorySearch true "分页获取分类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category/getCategoryList [get]
func (categoryApi *CategoryApi) GetCategoryList(c *gin.Context) {
	var pageInfo test_autocodeReq.CategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := categoryService.GetCategoryInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
