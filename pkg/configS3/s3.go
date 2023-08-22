package configs3

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func S3Config() *s3.Client {
	credential := credentials.NewStaticCredentialsProvider(os.Getenv("S3_ACESS_KEY_ID"), os.Getenv("S3_SECRET_ACCESS_KEY"), "")
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(credential), config.WithRegion(os.Getenv("S3_REGION")))
	if err != nil {
		log.Fatal(err)
	}
	return s3.NewFromConfig(config)
}
