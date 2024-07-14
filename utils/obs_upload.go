package utils

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"io"
	"mime/multipart"
	"path"
	"time"
)

// ObsBucket 定义OBS桶类
type ObsBucket struct {
	endpoint   string
	bucketName string
	domain     string
	ak         string // 敏感
	sk         string // 敏感
}

// ObsCommonBucket 定义第一个公共桶对象
var ObsCommonBucket ObsBucket

// InitObsCommonBucket 第一个公共桶对象的初始化函数
func InitObsCommonBucket() {
	ObsCommonBucket = ObsBucket{
		endpoint:   conf.ObsEndPoint,
		bucketName: conf.ObsBucketName,
		domain:     conf.ObsDomain,
		ak:         conf.ObsAk,
		sk:         conf.ObsSk,
	}
}

// UploadFile 给OBS桶类定义上传文件的对象函数
func (bucket ObsBucket) UploadFile(uploadPath string, file io.Reader, fileSuffix string) (string, string, string, error) {
	var obsClient, _ = obs.New(bucket.ak, bucket.sk, bucket.endpoint)

	uuid := GetUUID()
	fileName := time.Now().Format("20060102") + uuid[:20] + fileSuffix

	input := &obs.PutObjectInput{}
	input.Bucket = bucket.bucketName
	input.Key = uploadPath + fileName
	input.Body = file

	output, err := obsClient.PutObject(input)
	// noinspection all
	if err == nil {
		logger.Infof("RequestId %v", output.RequestId)
		logger.Infof("ETag %v", output.ETag)
		logger.Infof("上传成功 %v", uploadPath+fileName)
		return bucket.domain, uploadPath, fileName, nil
	} else if obsError, ok := err.(obs.ObsError); ok {
		logger.Errorf("Code %v", obsError.Code)
		logger.Errorf("Message %v", obsError.Message)
	}
	return "", "", "", errors.New("上传OBS错误")
}

func (bucket ObsBucket) Upload(uploadPath string, fileHeader *multipart.FileHeader) (string, string, string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		logger.Errorf("文件打开错误 %v", err)
		return "", "", "", errors.New("文件打开错误")
	}
	return bucket.UploadFile(uploadPath, file, path.Ext(fileHeader.Filename))
}
