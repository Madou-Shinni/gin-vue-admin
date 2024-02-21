// 自动生成模板Article
package test_autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 文章 结构体  Article
type Article struct {
	global.GVA_MODEL
	CategoryId *int   `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:;"` //分类
	Name       string `json:"name" form:"name" gorm:"column:name;comment:;"`                    //文章名称
}

// TableName 文章 Article自定义表名 article
func (Article) TableName() string {
	return "article"
}
