package s3

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"time"
)

func (s3Storage AWS) Connect() (*session.Session, error) {
	key := s3Storage.SpaceKey
	secret := s3Storage.SpaceSecret

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String(s3Storage.SpaceEndpoint),
		Region:      aws.String(s3Storage.SpaceRegion),
	}

	sess, err := session.NewSession(s3Config)
	if err != nil {
		return nil, err
	}
	return sess, err
}

func (s3Storage AWS) Upload(ctx context.Context, sess *session.Session, file io.Reader, fileName string) (*s3manager.UploadOutput, error) {
	bucket := s3Storage.SpaceBucket

	uploader := s3manager.NewUploader(sess)
	return uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})
}

func (s3Storage AWS) Delete(sess *session.Session, filePath string) error {
	bucket := s3Storage.SpaceBucket
	svc := s3.New(sess)

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filePath),
	})
	if err != nil {
		return err
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filePath),
	})
	if err != nil {
		return err
	}

	return nil

}

func (s3Storage AWS) Get(ctx context.Context, sess *session.Session, filePath string, signedPeriod time.Duration) (string, error) {

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s3Storage.SpaceBucket),
		Key:    aws.String(filePath),
	})
	urlStr, err := req.Presign(signedPeriod)

	if err != nil {
		return "", err
	}

	return urlStr, err
}
