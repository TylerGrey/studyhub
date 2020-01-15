package s3

import (
	"bytes"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
)

// S3 ...
type S3 struct {
	s3 *awsS3.S3
}

// NewS3 ...
func NewS3(accessKey string, secretKey string, region string) *S3 {
	cred := credentials.NewStaticCredentials(accessKey, secretKey, "")
	config := aws.NewConfig().WithRegion(region).WithCredentials(cred)

	s3 := &S3{
		s3: awsS3.New(session.New(config)),
	}
	return s3
}

// UploadS3ByRaw ...
func (s3Obj *S3) UploadS3ByRaw(key string, bucket string, raw []byte, contentType string) error {
	log.Println()
	params := &awsS3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(raw),
		ContentType: aws.String(contentType),
		ACL:         aws.String("public-read"),
		Metadata: map[string]*string{
			"Content-Type": aws.String(contentType), // Required
		},
	}
	if _, err := s3Obj.s3.PutObject(params); err != nil {
		return err
	}

	return nil
}
