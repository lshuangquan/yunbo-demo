package page

import (
	"context"
	v1 "github.com/lshuangquan/yunbo-demo/api/page/v1"
)

type IPageV1 interface {
	ViewAdd(ctx context.Context, req *v1.ViewAddReq) (res *v1.ViewAddRes, err error)
	ViewCount(ctx context.Context, req *v1.ViewCountReq) (res *v1.ViewCountRes, err error)
}
