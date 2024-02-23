package file

import (
	"github.com/lishuangquan/yunbo-demo/api/file"
)

type FileController struct {
}

func New() file.IFileV1 {
	return &FileController{}
}
