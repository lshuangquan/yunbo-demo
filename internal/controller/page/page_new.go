package page

import (
	"github.com/lshuangquan/yunbo-demo/api/page"
)

type PageController struct {
}

func New() page.IPageV1 {
	return new(PageController)
}
