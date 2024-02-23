package page

import (
	"context"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	v1 "github.com/lshuangquan/yunbo-demo/api/page/v1"
)

func (p *PageController) ViewAdd(ctx context.Context, req *v1.ViewAddReq) (res *v1.ViewAddRes, err error) {
	//TODO implement me
	defer func() {
		exception := recover()
		if exception != nil {
			glog.Errorf(ctx, "viewAdd  error:%+v", err)
		}
	}()
	//todo . check the pagePath is right exist
	viewCount, redisErr := g.Redis().IncrBy(ctx, req.PagePath, 1)
	if redisErr != nil {
		glog.Errorf(ctx, "redis query  error:%+v", redisErr)
		res = nil
		err = gerror.NewCode(gcode.New(500, "internal error ", "system internal error"))
		panic(err)
	}
	res = new(v1.ViewAddRes)
	res.Count = viewCount
	return res, nil
}

func (p *PageController) ViewCount(ctx context.Context, req *v1.ViewCountReq) (res *v1.ViewCountRes, err error) {

	defer func() {
		exception := recover()
		if exception != nil {
			glog.Errorf(ctx, "page viewCount  error:%+v", err)
		}
	}()

	viewCount, redisErr := g.Redis().Get(ctx, req.PagePath)
	if redisErr != nil {
		glog.Errorf(ctx, "redis query  error:%+v", redisErr)
		res = nil
		err = gerror.NewCode(gcode.New(500, "internal error ", "system internal error"))
		panic(err)
	}
	res = new(v1.ViewCountRes)
	res.Count = viewCount.Int64()
	return res, nil
}
