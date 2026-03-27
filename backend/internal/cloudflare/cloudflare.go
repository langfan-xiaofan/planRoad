package cloudflare

import (
	"backend/internal/config"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var r2Client *s3.Client

func CloudFlareInit() error {
	endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", config.Conf.Cloudflare.AccountId)
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion("auto"),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			config.Conf.Cloudflare.ApiKey,
			config.Conf.Cloudflare.ApiSecret,
			"",
		)),
	)
	if err != nil {
		return err
	}
	r2Client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})
	return nil
}

func UploadFile(ctx context.Context, ObjecktKey string, file io.Reader, fileType string) (string, error) {
	_, err := r2Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(config.Conf.Cloudflare.BucketName),
		Key:         aws.String(ObjecktKey),
		Body:        file,
		ContentType: aws.String(fileType),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", config.Conf.Cloudflare.PublicKey, ObjecktKey), nil
}
