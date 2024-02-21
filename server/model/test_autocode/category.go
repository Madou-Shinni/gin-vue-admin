// 自动生成模板Category
package test_autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 分类 结构体  Category
type Category struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:;"` //分类名称
}

// TableName 分类 Category自定义表名 category
func (Category) TableName() string {
	return "category"
}
