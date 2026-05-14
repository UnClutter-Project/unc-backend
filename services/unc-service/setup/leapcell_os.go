package setup

import (
	"context"
	"unc/services/unc-service/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func setupObjectStorage(ctx context.Context) *s3.Client {
	//cfg, err := awsconfig.LoadDefaultConfig(ctx,
	//    awsconfig.WithCredentialsProvider(
	//        credentials.NewStaticCredentialsProvider(
	//            config.GetConfig().ObjectStorageAccessID,
	//            config.GetConfig().ObjectStorageAccessSecret,
	//            ""),
	//    ),
	//    awsconfig.WithRegion("us-east-1"),
	//    awsconfig.WithBaseEndpoint("https://objstorage.leapcell.io"),
	//)
	cfg := aws.Config{
		Region:       "us-east-1",
		Credentials:  credentials.NewStaticCredentialsProvider(config.GetConfig().ObjectStorageAccessID, config.GetConfig().ObjectStorageAccessSecret, ""),
		BaseEndpoint: aws.String("https://objstorage.leapcell.io"),
	}
	//if err != nil {
	//    log.Fatalf("Unable to connect to Leapcell Object Storage: %v", err)
	//}

	client := s3.NewFromConfig(cfg)

	return client
}
