import service from '@/utils/request'

// @Tags Category
// @Summary 创建分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category true "创建分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /category/createCategory [post]
export const createCategory = (data) => {
  return service({
    url: '/category/createCategory',
    method: 'post',
    data
  })
}

// @Tags Category
// @Summary 删除分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category true "删除分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /category/deleteCategory [delete]
export const deleteCategory = (params) => {
  return service({
    url: '/category/deleteCategory',
    method: 'delete',
    params
  })
}

// @Tags Category
// @Summary 批量删除分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /category/deleteCategory [delete]
export const deleteCategoryByIds = (params) => {
  return service({
    url: '/category/deleteCategoryByIds',
    method: 'delete',
    params
  })
}

// @Tags Category
// @Summary 更新分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Category true "更新分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /category/updateCategory [put]
export const updateCategory = (data) => {
  return service({
    url: '/category/updateCategory',
    method: 'put',
    data
  })
}

// @Tags Category
// @Summary 用id查询分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Category true "用id查询分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /category/findCategory [get]
export const findCategory = (params) => {
  return service({
    url: '/category/findCategory',
    method: 'get',
    params
  })
}

// @Tags Category
// @Summary 分页获取分类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取分类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category/getCategoryList [get]
export const getCategoryList = (params) => {
  return service({
    url: '/category/getCategoryList',
    method: 'get',
    params
  })
}
