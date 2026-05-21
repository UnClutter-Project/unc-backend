package client

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageClient interface {
	UploadFile(ctx context.Context, fileKey string, fileData []byte) (string, error)
	GetPresignedLink(ctx context.Context, fileKey string) (string, error)
}

type StorageClientImpl struct {
	s3Client        *s3.Client
	s3PresignClient *s3.PresignClient
	bucketName      string
	presignDuration time.Duration
}

func NewStorageClient(s3Client *s3.Client, bucketName string, presignDuration time.Duration) *StorageClientImpl {
	return &StorageClientImpl{
		s3Client:        s3Client,
		s3PresignClient: s3.NewPresignClient(s3Client),
		bucketName:      bucketName,
		presignDuration: presignDuration,
	}
}

func (c *StorageClientImpl) UploadFile(ctx context.Context, fileKey string, fileData []byte) (string, error) {
	_, err := c.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(c.bucketName),
		Key:         aws.String(fileKey),
		Body:        bytes.NewReader(fileData),
		ContentType: aws.String(http.DetectContentType(fileData[:512])),
	})
	if err != nil {
		return "", err
	}

	return fileKey, nil
}

func (c *StorageClientImpl) GetPresignedLink(ctx context.Context, fileKey string) (string, error) {
	file, err := c.s3PresignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(fileKey),
	}, s3.WithPresignExpires(c.presignDuration))
	if err != nil {
		return "", err
	}

	return file.URL, nil
}
