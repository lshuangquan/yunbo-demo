package main

import (
	_ "github.com/lishuangquan/yunbo-demo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/lishuangquan/yunbo-demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
