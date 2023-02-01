package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/apply"
	"github.com/flipped-aurora/gin-vue-admin/server/router/community"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/general"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	Community community.RouterGroup
	Apply     apply.RouterGroup
	General   general.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
