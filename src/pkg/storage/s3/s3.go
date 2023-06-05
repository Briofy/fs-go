package s3

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"time"
)

type Storage interface {
	Connect() (*session.Session, error)
	Upload(ctx context.Context, sess *session.Session, file io.Reader, fileName string) (*s3manager.UploadOutput, error)
	Delete(sess *session.Session, filePath string) error
	Get(ctx context.Context, sess *session.Session, filePath string, signedPeriod time.Duration) (string, error)
}
