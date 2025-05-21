package processor

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"reviewservice/config"
	"reviewservice/internal/db"
	"reviewservice/internal/models"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gorm.io/gorm"
)

func RunIngestion(cfg *config.Config, s3Client *s3.S3, dbConn *gorm.DB) error {
	// List files in the bucket
	listOutput, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: awsString(cfg.S3Bucket),
		Prefix: awsString(cfg.S3Prefix),
	})
	if err != nil {
		return err
	}

	for _, obj := range listOutput.Contents {
		if strings.HasSuffix(*obj.Key, ".jl") {
			log.Printf("Processing file: %s", *obj.Key)

			// Download file
			buf := aws.NewWriteAtBuffer([]byte{})
			downloader := s3manager.NewDownloaderWithClient(s3Client)
			_, err := downloader.Download(buf, &s3.GetObjectInput{
				Bucket: awsString(cfg.S3Bucket),
				Key:    obj.Key,
			})
			if err != nil {
				log.Printf("Failed to download %s: %v", *obj.Key, err)
				continue
			}

			scanner := bufio.NewScanner(bytes.NewReader(buf.Bytes()))
			for scanner.Scan() {
				line := scanner.Text()
				var review models.Review
				if err := json.Unmarshal([]byte(line), &review); err != nil {
					log.Printf("Skipping malformed line: %v", err)
					continue
				}

				if err := db.SaveReview(dbConn, &review); err != nil {
					log.Printf("Failed to save review: %v", err)
				}
			}
		}
	}

	return nil
}

func awsString(s string) *string {
	return &s
}
