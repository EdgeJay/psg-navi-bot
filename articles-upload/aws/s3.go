package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Path struct {
	Bucket *string
	Key    *string
}

func GetS3Client() *s3.S3 {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)
	return svc
}

func CreateGetObjectInput(path S3Path) *s3.GetObjectInput {
	return &s3.GetObjectInput{
		Bucket: path.Bucket,
		Key:    path.Key,
	}
}
