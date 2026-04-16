package client

import (
	"unc/services/unc-service/config"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Clients struct {
	EmailClient   EmailClient
	StorageClient StorageClient
}

func SetupClients(s3Client *s3.Client) *Clients {
	return &Clients{
		EmailClient:   NewEmailClient(),
		StorageClient: NewStorageClient(s3Client, config.GetConfig().BucketName, config.GetConfig().PresignDuration),
	}
}
