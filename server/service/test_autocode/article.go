package test_autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/test_autocode"
	test_autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/test_autocode/request"
)

type ArticleService struct {
}

// CreateArticle 创建文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) CreateArticle(article *test_autocode.Article) (err error) {
	err = global.GVA_DB.Create(article).Error
	return err
}

// DeleteArticle 删除文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) DeleteArticle(ID string) (err error) {
	err = global.GVA_DB.Delete(&test_autocode.Article{}, "id = ?", ID).Error
	return err
}

// DeleteArticleByIds 批量删除文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) DeleteArticleByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]test_autocode.Article{}, "id in ?", IDs).Error
	return err
}

// UpdateArticle 更新文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) UpdateArticle(article test_autocode.Article) (err error) {
	err = global.GVA_DB.Save(&article).Error
	return err
}

// GetArticle 根据ID获取文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) GetArticle(ID string) (article test_autocode.Article, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&article).Error
	return
}

// GetArticleInfoList 分页获取文章记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) GetArticleInfoList(info test_autocodeReq.ArticleSearch) (list []test_autocode.Article, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&test_autocode.Article{})
	var articles []test_autocode.Article
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

	err = db.Find(&articles).Error
	return articles, total, err
}
