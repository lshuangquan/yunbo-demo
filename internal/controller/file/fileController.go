package file

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	v1 "github.com/lshuangquan/yunbo-demo/api/file/v1"
	"github.com/lshuangquan/yunbo-demo/internal/service"
	"github.com/minio/minio-go/v7"
	"io"
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

func (c *FileController) View(req *ghttp.Request) {
	//todo 1. check fileName is valid
	fileName := req.GetQuery("fileName").String()
	if fileName == "" {
		req.Response.WriteJson(gcode.New(1, "param invalid", "filename is empty"))
		return
	}
	//todo 2, check the file is exist;
	//todo 3. user has the privileges to view the file
	fileObj, fileErr := service.GetObject(req.GetCtx(), fileName)
	defer fileObj.Close()
	if fileErr != nil {
		glog.Errorf(req.GetCtx(), "view error :%+v", fileErr)
		req.Response.WriteJson(gcode.New(500, "upload file error", "system internal error"))
		return
	}
	writer := req.Response.RawWriter()
	fileBytes, err := io.ReadAll(fileObj)
	if err != nil {
		glog.Errorf(req.GetCtx(), "read file from oss error :%+v", fileErr)
		req.Response.WriteJson(gcode.New(501, "read file from oss error", "system internal error"))
		return
	}
	req.Response.Header().Add("Content-type", "application/octet-stream")
	req.Response.Header().Add("Content-Disposition", fmt.Sprintf("%s; filename=%s", gmd5.MustEncryptString(fileName), fileName))
	writer.WriteHeader(200)
	_, err = writer.Write(fileBytes)
	if err != nil {
		glog.Errorf(req.GetCtx(), "resp write   error :%+v", err)
		//req.Response.WriteJson(gcode.New(502, "resp write error", "system internal error"))
		return
	}

	return
}

func (c *FileController) Delete(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {

	//todo 1. check fileName is valid
	//todo 2, check the file is exist;
	//todo 3. user has the privileges to del the file
	err = service.DelObject(ctx, req.FileName)
	if err != nil {
		return nil, gerror.NewCode(gcode.New(500, "upload file error", "system internal error"))
	}
	return res, nil
}
