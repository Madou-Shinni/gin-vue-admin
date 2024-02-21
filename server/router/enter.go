package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/test_autocode"
)

type RouterGroup struct {
	System        system.RouterGroup
	Example       example.RouterGroup
	Test_autocode test_autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
