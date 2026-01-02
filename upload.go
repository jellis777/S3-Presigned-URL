package s3share

import (
	"bytes"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Upload(client *s3.Client, filename *string, bucket *string) error {
	content, err := os.ReadFile(*filename)
	if err != nil {
		return err
	}

	key := filename
	_, err = client.PutObject(
		context.Background(),
		&s3.PutObjectInput{
			Bucket: bucket,
			Key:    key,
			Body:   bytes.NewReader(content),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
