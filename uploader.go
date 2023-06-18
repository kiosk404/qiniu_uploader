package main

import "github.com/qiniu/api.v7/storage"

// 上传工具接口
type Uploader interface {
	Upload(filepath string) error
}

type UploaderFactory struct {}

// 创建上传工具
func (f *UploaderFactory) CreateUploader(accessKey, secretKey, bucket string, zone *storage.Zone) Uploader {
	return NewQiniuUploader(accessKey, secretKey, bucket, zone)
}
