package file

import (
	"context"
	v1 "github.com/lishuangquan/yunbo-demo/api/file/v1"
)

type IFileV1 interface {
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
	View(ctx context.Context, req *v1.ViewReq) (res *v1.ViewRes, err error)
	Delete(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error)
}
