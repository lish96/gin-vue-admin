package v1

import (
	"github.com/lish96/gin-vue-admin/server/api/v1/example"
	"github.com/lish96/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
