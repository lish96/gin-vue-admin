package request

import (
	"github.com/lish96/gin-vue-admin/server/model/common/request"
	"github.com/lish96/gin-vue-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
