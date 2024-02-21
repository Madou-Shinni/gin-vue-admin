import service from '@/utils/request'

// @Tags Article
// @Summary 创建文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Article true "创建文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /article/createArticle [post]
export const createArticle = (data) => {
  return service({
    url: '/article/createArticle',
    method: 'post',
    data
  })
}

// @Tags Article
// @Summary 删除文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Article true "删除文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /article/deleteArticle [delete]
export const deleteArticle = (params) => {
  return service({
    url: '/article/deleteArticle',
    method: 'delete',
    params
  })
}

// @Tags Article
// @Summary 批量删除文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /article/deleteArticle [delete]
export const deleteArticleByIds = (params) => {
  return service({
    url: '/article/deleteArticleByIds',
    method: 'delete',
    params
  })
}

// @Tags Article
// @Summary 更新文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Article true "更新文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /article/updateArticle [put]
export const updateArticle = (data) => {
  return service({
    url: '/article/updateArticle',
    method: 'put',
    data
  })
}

// @Tags Article
// @Summary 用id查询文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Article true "用id查询文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /article/findArticle [get]
export const findArticle = (params) => {
  return service({
    url: '/article/findArticle',
    method: 'get',
    params
  })
}

// @Tags Article
// @Summary 分页获取文章列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取文章列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /article/getArticleList [get]
export const getArticleList = (params) => {
  return service({
    url: '/article/getArticleList',
    method: 'get',
    params
  })
}
