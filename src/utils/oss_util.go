package utils

import (
	"dou-xiao-yin/src/config"
	"dou-xiao-yin/src/model"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"strconv"
)

var bucket *oss.Bucket

func InitOssConnect() error {
	// 配置 oss 的相关参数
	conf := config.GetConf()
	Endpoint := conf.Oss.Endpoint
	AccessKeyId := conf.Oss.AccessKeyId
	AccessKeySecret := conf.Oss.AccessKeySecret
	BucketName := conf.Oss.BucketName
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		return err
	}
	Bucket, err := client.Bucket(BucketName)
	if err != nil {
		return err
	}
	if bucket == nil {
		bucket = Bucket
	}
	return nil
}

func GetDefaultBucket() (*oss.Bucket, error) {
	if bucket == nil {
		err := InitOssConnect()
		if err != nil {
			return nil, err
		}
	}
	return bucket, nil
}

func UploadVideo(file *multipart.FileHeader, video *model.Video) error {
	// 得到连接
	bucket, err := GetDefaultBucket()
	if err != nil {
		return err
	}

	// 打开文件流
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(src)

	// 将文件流上传至视频目录下
	path := "videos/" + strconv.Itoa(video.AuthorId) + "/" + strconv.Itoa(video.Id) + ".mp4"
	err = bucket.PutObject(path, src)
	if err != nil {
		fmt.Println("Error:", err, path)
		return err
	}

	// 为文件创建路径软连接
	symObjectKey := "douyin/resources/video/" + strconv.Itoa(video.AuthorId) + "/" + strconv.Itoa(video.Id)
	objectName := path

	option := []oss.Option{
		// 指定创建软链接时是否覆盖同名Object。
		oss.ForbidOverWrite(true),
		// 指定Object的访问权限。此处指定为PublicRead，表示公共读。
		oss.ObjectACL(oss.ACLPublicRead),
		// 指定Object的存储类型。此处指定为Standard，表示标准存储类型。
		oss.StorageClass(oss.StorageStandard),
	}
	err = bucket.PutSymlink(symObjectKey, objectName, option...)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return nil
}
