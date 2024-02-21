package test_autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/test_autocode"
	test_autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/test_autocode/request"
)

type CategoryService struct {
}

// CreateCategory 创建分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) CreateCategory(category *test_autocode.Category) (err error) {
	err = global.GVA_DB.Create(category).Error
	return err
}

// DeleteCategory 删除分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) DeleteCategory(ID string) (err error) {
	err = global.GVA_DB.Delete(&test_autocode.Category{}, "id = ?", ID).Error
	return err
}

// DeleteCategoryByIds 批量删除分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) DeleteCategoryByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]test_autocode.Category{}, "id in ?", IDs).Error
	return err
}

// UpdateCategory 更新分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) UpdateCategory(category test_autocode.Category) (err error) {
	err = global.GVA_DB.Save(&category).Error
	return err
}

// GetCategory 根据ID获取分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) GetCategory(ID string) (category test_autocode.Category, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&category).Error
	return
}

// GetCategoryInfoList 分页获取分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) GetCategoryInfoList(info test_autocodeReq.CategorySearch) (list []test_autocode.Category, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&test_autocode.Category{})
	var categorys []test_autocode.Category
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&categorys).Error
	return categorys, total, err
}
