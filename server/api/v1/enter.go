package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/apply"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/community"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/general"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	CommunityApiGroup community.ApiGroup
	ApplyApiGroup     apply.ApiGroup
	GeneralApiGroup   general.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
