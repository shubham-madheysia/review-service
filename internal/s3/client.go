package s3

import (
	"reviewservice/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewClient(cfg *config.Config) (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(cfg.AWSAccessKey, cfg.AWSSecretKey, ""),
	})
	if err != nil {
		return nil, err
	}

	return s3.New(sess), nil
}
