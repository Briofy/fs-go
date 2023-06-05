package repository

import (
	"context"
	"github.com/Briofy/fs-go/src/config"
	"github.com/Briofy/fs-go/src/internal/contract/storage"
	s3Pkg "github.com/Briofy/fs-go/src/pkg/storage/s3"
	"io"
	"time"
)

type S3StorageRepo struct {
	s3           s3Pkg.Storage
	signedPeriod time.Duration
}

func (s S3StorageRepo) GetBatch(ctx context.Context, filePath []string) ([]string, error) {
	sess, err := s.s3.Connect()
	if err != nil {
		return nil, err
	}
	var links []string
	for _, path := range filePath {
		link, err := s.s3.Get(ctx, sess, path, s.signedPeriod)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}
	return links, nil
}

func NewS3StorageRepo(s3 s3Pkg.Storage, cfg config.Config) storage.Repository {
	return &S3StorageRepo{
		s3:           s3,
		signedPeriod: cfg.GeSignedPeriod(),
	}
}

func (s S3StorageRepo) Get(ctx context.Context, filePath string) (string, error) {
	sess, err := s.s3.Connect()
	if err != nil {
		return "", err
	}
	return s.s3.Get(ctx, sess, filePath, s.signedPeriod)
}

func (s S3StorageRepo) Upload(ctx context.Context, reader io.Reader, fileName string) error {
	sess, err := s.s3.Connect()
	if err != nil {
		return err
	}
	_, err = s.s3.Upload(ctx, sess, reader, fileName)
	return err
}

func (s S3StorageRepo) Delete(filePath string) error {
	sess, err := s.s3.Connect()
	if err != nil {
		return err
	}
	return s.s3.Delete(sess, filePath)
}
