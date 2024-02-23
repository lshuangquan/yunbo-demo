package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ViewAddReq struct {
	g.Meta   `path:"/viewadd" tags:"upload" mime:"application/json" method:"put" summary:"添加统计信息"`
	PagePath string `json:"pagePath" type:"string" dc:"pagePath"`
}
type ViewAddRes struct {
	g.Meta `mime:"application/json" example:"{}"`
	Count  int64 `json:"count" type:"int" dc:"文件路径"`
}

type ViewCountReq struct {
	g.Meta   `path:"/viewcount" tags:"view" method:"get" summary:"count page view "`
	PagePath string `json:"pagePath" type:"string" dc:"count page view"`
}

type ViewCountRes struct {
	g.Meta `mime:"application/json" example:"{}"`
	Count  int64 `json:"count" type:"int" dc:"page view"`
}
