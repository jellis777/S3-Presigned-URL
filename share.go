package s3share

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Share(client *s3.Client, key *string, bucket *string) (string, error) {
	var url string

	lifetimeSecs := int64(3600)
	s3PresignClient := s3.NewPresignClient(client)
	req, err := s3PresignClient.PresignGetObject(
		context.TODO(),
		&s3.GetObjectInput{
			Bucket: bucket,
			Key:    key,
		},
		func(opts *s3.PresignOptions) {
			opts.Expires = time.Duration(lifetimeSecs * int64(time.Second))
		})
	if err != nil {
		log.Printf("Couldn't get a presigned request to get %v:%v. Here's why: %v\n",
			*bucket, *key, err)
		return "", err
	}

	url = string(req.URL)

	return url, nil
}
