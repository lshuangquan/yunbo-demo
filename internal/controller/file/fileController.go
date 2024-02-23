package file

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
	v1 "github.com/lishuangquan/yunbo-demo/api/file/v1"
	"github.com/lishuangquan/yunbo-demo/internal/service"
	"github.com/minio/minio-go/v7"
)

func (c *FileController) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, errRes error) {
	defer func() {
		err := recover()
		if err != nil {
			glog.Errorf(ctx, "upload file error fileName:%+v", err)
		}
	}()
	glog.Infof(ctx, "file length:%d", req.File.Size)
	if req.File.Size > 1000000 {

		return nil, gerror.NewCode(gcode.New(1, "upload file length > 1000", "valid error"))
	}

	//todo check fileName is exit;

	//filePath, saveFileError := req.File.Save("/tmp/")
	//if saveFileError != nil {
	//	return nil, gerror.NewCode(gcode.New(1, "upload file length > 1000", " valid error"))
	//}

	fileOpen, _ := req.File.Open()
	defer fileOpen.Close()
	uploadInfo, uploadError := service.MinioClient.PutObject(ctx, service.GetBucketName(), req.File.Filename, fileOpen, req.File.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if uploadError != nil {
		glog.Errorf(ctx, "upload error :%+v", uploadError)
		return nil, gerror.NewCode(gcode.New(1, "upload file error", "system internal error"))
	}
	file := new(v1.UploadRes)
	file.FileUrl = uploadInfo.Location
	return file, nil
}

func (c *FileController) View(ctx context.Context, req *v1.ViewReq) (res *v1.ViewRes, err error) {
	return nil, nil
}

func (c *FileController) Delete(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	return nil, nil
}
