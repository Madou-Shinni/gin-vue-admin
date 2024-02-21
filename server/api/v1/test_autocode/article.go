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

type ArticleApi struct {
}

var articleService = service.ServiceGroupApp.Test_autocodeServiceGroup.ArticleService

// CreateArticle 创建文章
// @Tags Article
// @Summary 创建文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test_autocode.Article true "创建文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /article/createArticle [post]
func (articleApi *ArticleApi) CreateArticle(c *gin.Context) {
	var article test_autocode.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := articleService.CreateArticle(&article); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteArticle 删除文章
// @Tags Article
// @Summary 删除文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test_autocode.Article true "删除文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /article/deleteArticle [delete]
func (articleApi *ArticleApi) DeleteArticle(c *gin.Context) {
	ID := c.Query("ID")
	if err := articleService.DeleteArticle(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteArticleByIds 批量删除文章
// @Tags Article
// @Summary 批量删除文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /article/deleteArticleByIds [delete]
func (articleApi *ArticleApi) DeleteArticleByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := articleService.DeleteArticleByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateArticle 更新文章
// @Tags Article
// @Summary 更新文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test_autocode.Article true "更新文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /article/updateArticle [put]
func (articleApi *ArticleApi) UpdateArticle(c *gin.Context) {
	var article test_autocode.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := articleService.UpdateArticle(article); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindArticle 用id查询文章
// @Tags Article
// @Summary 用id查询文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query test_autocode.Article true "用id查询文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /article/findArticle [get]
func (articleApi *ArticleApi) FindArticle(c *gin.Context) {
	ID := c.Query("ID")
	if rearticle, err := articleService.GetArticle(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rearticle": rearticle}, c)
	}
}

// GetArticleList 分页获取文章列表
// @Tags Article
// @Summary 分页获取文章列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query test_autocodeReq.ArticleSearch true "分页获取文章列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /article/getArticleList [get]
func (articleApi *ArticleApi) GetArticleList(c *gin.Context) {
	var pageInfo test_autocodeReq.ArticleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := articleService.GetArticleInfoList(pageInfo); err != nil {
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
