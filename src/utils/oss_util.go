package utils

import (
	"bytes"
	"dou-xiao-yin/src/config"
	"dou-xiao-yin/src/model"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

func UploadVideo(file *multipart.FileHeader, video *model.Video) (error, string) {
	// 得到连接
	bucket, err := GetDefaultBucket()
	if err != nil {
		return err, ""
	}

	// 打开文件流
	src, err := file.Open()
	if err != nil {
		return err, ""
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(src)

	// TODO: 获取文件封面

	// 将文件流上传至视频目录下
	path := "videos/" + strconv.Itoa(video.AuthorId) + "/" + strconv.Itoa(video.Id) + ".mp4"
	err = bucket.PutObject(path, src)
	if err != nil {
		fmt.Println("Error:", err, path)
		return err, ""
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
	return nil, symObjectKey
}

// UploadCover : 从视频中截取封面并上传至oss
func UploadCover(videoPath string) {
	// 通过video路径获取authorId和videoId，从而获得封面本地保存路径
	arr := strings.Split(videoPath, "/")
	authorId := arr[len(arr)-2]
	videoId := arr[len(arr)-1]
	imgFile := fmt.Sprintf("./images/%s_%s.jpg", authorId, videoId)

	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "1", "-f", "image2", imgFile)
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	if cmd.Run() != nil {
		fmt.Println("could not generate frame")
	}
	defer os.Remove(imgFile)
	// -----------------------------------------

	bucket, err := GetDefaultBucket()
	if err != nil {
		fmt.Println(err)
	}

	// 打开文件流
	img, err := os.Open(imgFile)
	if err != nil {
		fmt.Println(err)
	}
	defer img.Close()

	// 将文件流上传至视频目录下
	path := "cover/" + authorId + "/" + videoId + ".jpg"
	err = bucket.PutObject(path, img)
	if err != nil {
		fmt.Println("Error:", err, path)
	}

	// 为文件创建路径软连接
	symObjectKey := "douyin/resources/cover/" + authorId + "/" + videoId
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
}
