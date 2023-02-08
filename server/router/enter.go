package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/app"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	//Community community.RouterGroup
	//Apply     apply.RouterGroup
	//General   general.RouterGroup
	App app.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
