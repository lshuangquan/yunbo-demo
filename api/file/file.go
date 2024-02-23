package file

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	v1 "github.com/lshuangquan/yunbo-demo/api/file/v1"
)

type IFileV1 interface {
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
	View(req *ghttp.Request)
	Delete(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error)
}
