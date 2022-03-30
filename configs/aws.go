package configs

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	AWS_S3_REGION = "" // Region
	AWS_S3_BUCKET = "" // Bucket
)

// We will be using this client everywhere in our code
var awsS3Client *s3.Client

// configS3 creates the S3 client
func configS3() {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(AWS_S3_REGION))
	if err != nil {
		log.Fatal(err)
	}

	awsS3Client = s3.NewFromConfig(cfg)
}
