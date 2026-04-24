package setup

import (
	"context"
	"log"
	"unc/services/unc-service/config"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func setupObjectStorage(ctx context.Context) *s3.Client {
	cfg, err := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				config.GetConfig().LeapcellAccessID,
				config.GetConfig().LeapcellAccessSecret,
				""),
		),
		awsconfig.WithRegion("us-east-1"),
		awsconfig.WithBaseEndpoint("https://objstorage.leapcell.io"),
	)
	if err != nil {
		log.Fatalf("Unable to connect to Leapcell Object Storage: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	return client
}
