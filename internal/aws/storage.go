package aws

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var BUCKET_NAME = os.Getenv("AWS_S3_BUCKET")
const (
	REGION = "us-east-2"
	PROGRESS_ENTRY_CONTAINER = "progress-entry"
	PROGRESS_VIDEO_CONTAINER = "progress-video"
)


func NewAWSCredentials() (cfg aws.Config, err error){
	return config.LoadDefaultConfig(context.TODO(), config.WithRegion(REGION))
}

func NewS3Client() (*s3.Client, error) {
	cfg, err := NewAWSCredentials()

	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}

func CreatePresignedUrl(key string, option string, expires time.Time) (*v4.PresignedHTTPRequest, error) {
	s3Client, err := NewS3Client()

	if err != nil {
		return nil, err
	}

	presignerClient := s3.NewPresignClient(s3Client)

	if option == "READ" {
		return presignerClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
			Bucket: aws.String(BUCKET_NAME),
			Key: aws.String(key),
			ResponseExpires: &expires,
		})
	}else if option == "WRITE" {
		return presignerClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(BUCKET_NAME),
			Key: aws.String(key),
			Expires: &expires,
		})
	}

	return nil, errors.New("invalid option")
}