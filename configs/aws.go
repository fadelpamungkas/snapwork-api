package configs

import (
	"context"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	// S3BucketName is the name of the S3 bucket
	S3BucketName = "snapworkupload"
	// S3Region is the region of the S3 bucket
	S3Region = "ap-southeast-1"
)

func UploadFileInS3Bucket(file io.Reader, fileName string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(S3Region),
	)
	client := s3.NewFromConfig(cfg)

	log.Println("Uploading file to S3 bucket...")
	uploadResp, err := manager.NewUploader(client).Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(S3BucketName),
		Key:         aws.String(fileName),
		Body:        file,
		ContentType: aws.String("image"),
	})
	if err != nil {
		log.Println("err: ", err)
		return "", err
	}
	return uploadResp.Location, nil
}
