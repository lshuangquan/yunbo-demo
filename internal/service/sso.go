package service

import (
	"context"
	"github.com/minio/minio-go/v7"
)

func GetObject(ctx context.Context, fileName string) (object *minio.Object, err error) {
	return MinioClient.GetObject(ctx, GetBucketName(), fileName, minio.GetObjectOptions{})
}

func DelObject(ctx context.Context, fileName string) (err error) {
	return MinioClient.RemoveObject(ctx, GetBucketName(), fileName, minio.RemoveObjectOptions{})
}
