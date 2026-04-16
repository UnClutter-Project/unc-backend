package setup

import (
	"context"
	"fmt"
	"log"
	"unc/services/unc-service/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func setupR2(ctx context.Context) *s3.Client {
	cfg, err := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				config.GetConfig().CloudflareAccessID,
				config.GetConfig().CloudflareAccessSecret,
				""),
		),
		awsconfig.WithRegion("auto"),
	)
	if err != nil {
		log.Fatalf("Unable to connect to R2: %v", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", config.GetConfig().CloudflareID))
	})

	return client
}
