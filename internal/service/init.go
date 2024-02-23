package service

import (
	"fmt"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var MinioClient *minio.Client
var bucketName string
var minioConfig *MinioConfig

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	BucketName      string
}

func init() {
	minio_setup()
}

func minio_setup() {
	minioConfig = new(MinioConfig)

	err := g.Cfg().MustGet(gctx.New(), "oss").Struct(minioConfig)
	if err != nil {
		glog.Fatal(gctx.New(), "init minio config error")
		return
	}
	glog.Infof(gctx.New(), "init minioConfig %+v", minioConfig)
	MinioClient, err = minio.New(minioConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioConfig.AccessKeyID, minioConfig.SecretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
	}

}

func GetBucketName() string {
	fmt.Println("bucketNmae --- " + minioConfig.BucketName)
	return minioConfig.BucketName
}
