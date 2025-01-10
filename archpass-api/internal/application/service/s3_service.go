package service

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

type S3Service struct {
	svc    *s3.S3
	bucket string
}

func NewS3Service(bucket string) (*S3Service, error) {
	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		return nil, fmt.Errorf("AWS_REGION environment variable is not set")
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}

	return &S3Service{
		svc:    s3.New(sess),
		bucket: bucket,
	}, nil
}

func (s *S3Service) UploadFile(key string, body *bytes.Buffer, contentType string) (string, error) {
	input := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(body.Bytes()),
		ContentType: aws.String(contentType),
	}

	_, err := s.svc.PutObject(input)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.bucket, key), nil
}
