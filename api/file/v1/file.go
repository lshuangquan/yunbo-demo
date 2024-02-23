package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadReq struct {
	g.Meta `path:"/file/upload" tags:"upload" mime:"multipart/form-data" method:"put" summary:"文件上传"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}
type UploadRes struct {
	g.Meta  `mime:"text/html" example:"string"`
	FileUrl string `json:"fileUrl" type:"string" dc:"文件路径"`
}

type ViewReq struct {
	g.Meta `path:"/file/view" tags:"view" method:"get" summary:"文件查看"`
}
type ViewRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type DelReq struct {
	g.Meta `path:"/file/delete" tags:"delete" method:"post" summary:"文件删除"`
}
type DelRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

//文件上传、文件查看、文件删除
