package setup

import (
	"unc/services/unc-service/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func setupLeapcellObjectStorage() *s3.Client {
	cfg := aws.Config{
		Region:       "us-east-1",
		Credentials:  credentials.NewStaticCredentialsProvider(config.GetConfig().ObjectStorageAccessID, config.GetConfig().ObjectStorageAccessSecret, ""),
		BaseEndpoint: aws.String("https://objstorage.leapcell.io"),
	}

	client := s3.NewFromConfig(cfg)

	return client
}
