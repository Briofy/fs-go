package storage

import (
	"context"
	"io"
)

type Reader interface {
	Get(ctx context.Context, id uint) (string, error)
}

type Writer interface {
	Upload(ctx context.Context, reader io.Reader, fileName string) error
}

//go:generate mockgen -destination=../../mock/repository/storage/storage_mock.go -package=StorageMock -self_package=github.com/fs-go/src/internal/contract/storage github.com/fs-go/src/internal/contract/storage Repository

type Repository interface {
	Reader
	Writer
}
