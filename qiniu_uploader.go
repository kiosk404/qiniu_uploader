package main

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// 七牛云存储上传工具
type QiniuUploader struct {
	mac *qbox.Mac
	uploader *storage.FormUploader
	bucket string
}

// 创建七牛云存储上传工具
func NewQiniuUploader(accessKey, secretKey, bucket string, zone *storage.Zone) *QiniuUploader {
	// 构建鉴权对象
	mac := qbox.NewMac(accessKey, secretKey)
	// 构建配置对象
	cfg := storage.Config{
		Zone: zone,
	}
	uploader := storage.NewFormUploader(&cfg)
	return &QiniuUploader{
		mac:      mac,
		uploader: uploader,
		bucket:   bucket,
	}
}

// 上传到七牛云存储
func (qu *QiniuUploader) Upload(localFilePath string) error {
	putPolicy := storage.PutPolicy{
		Scope: qu.bucket,
	}
	upToken := putPolicy.UploadToken(qu.mac)

	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "gallery",
		},
	}

	fileKey := newPathFormat(localFilePath)
	err := qu.uploader.PutFile(context.Background(), &ret, upToken, fileKey, localFilePath, &putExtra)

	if err == nil {
		fmt.Println("URL", ret.Key)
		fmt.Println("Hash", ret.Hash)
	}

	return err
}