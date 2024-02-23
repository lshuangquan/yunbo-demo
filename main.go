package main

import (
	_ "github.com/lshuangquan/yunbo-demo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/lshuangquan/yunbo-demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
