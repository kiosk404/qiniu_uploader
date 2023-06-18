package main

import (
	"fmt"
	"github.com/qiniu/api.v7/storage"
	"github.com/urfave/cli"
	"os"
)

// 定义命令行选项
var (
	accessKey 	string
	secretKey 	string
	bucket    	string
	zone      	string
	baseDir   	string
	domainHost 	string
	newName     string
)

func main() {
	app := cli.NewApp()
	app.Name = "qiniu-uploader"
	app.Usage = `Upload file to Qiniu Cloud Storage
	example: ./qiniu-uploader --newName www.png /home/work/Pictures/abcd.png
	`

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "access-key",
			Usage:       "Access key of Cloud Storage",
			Destination: &accessKey,
			EnvVar:      "UPLOADER_ACCESS_KEY",
			Required:    true,
		},
		cli.StringFlag{
			Name:        "secret-key",
			Usage:       "Secret key of Cloud Storage",
			Destination: &secretKey,
			EnvVar:      "UPLOADER_SECRET_KEY",
			Required:    true,
		},
		cli.StringFlag{
			Name:        "bucket",
			Usage:       "Bucket name of Cloud Storage",
			Destination: &bucket,
			EnvVar:      "UPLOADER_BUCKET",
			Required:    true,
		},
		cli.StringFlag{
			Name:        "zone",
			Usage:       "Zone of Cloud Storage (huadong, huabei, huanan, beimei)",
			Value:       "huanan",
			Destination: &zone,
			EnvVar:      "UPLOADER_ZONE",
			Required:    true,
		},
		cli.StringFlag{
			Name:        "baseDir",
			Usage:       "the basic path for uploading files, default value \"static\"",
			EnvVar:      "UPLOADER_BASE_DIR",
			Required:    false,
			Value:       "",
			Destination: &baseDir,
		},
		cli.StringFlag{
			Name:        "domainHost",
			Usage:       "domain url",
			EnvVar:      "UPLOADER_DOMAIN",
			Required:    false,
			Value:       "",
			Destination: &domainHost,
		},
		cli.StringFlag{
			Name:        "newName",
			Usage:       "used to replace the previous file name for uploading",
			Required:    false,
			Destination: &newName,
		},
	}
	app.Action = func(c *cli.Context) error {
		// 获取要上传的文件地址
		args := c.Args()
		if len(args) < 1 {
			return cli.NewExitError("Please specify the file path", 1)
		}
		localFilePath := args[0]

		_, err := os.Stat(localFilePath)
		if os.IsNotExist(err) {
			return cli.NewExitError(fmt.Sprintf("file %s does not exist", localFilePath),1)
		}

		var z *storage.Zone
		switch zone {
		case "huadong":
			z = &storage.ZoneHuadong
		case "huabei":
			z = &storage.ZoneHuabei
		case "huanan":
			z = &storage.ZoneHuanan
		case "beimei":
			z = &storage.ZoneBeimei
		default:
			return cli.NewExitError("Invalid zone", 1)
		}

		uploader := NewQiniuUploader(accessKey, secretKey, bucket, z)
		// 上传文件到七牛云存储
		err = uploader.Upload(localFilePath)
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		fmt.Println("upload success")
		fmt.Println(fmt.Sprintf("http://%s/%s",domainHost, newPathFormat(localFilePath)))
		return nil
	}
	// 运行命令行应用
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}



