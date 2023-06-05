package repository

import (
	"context"
	"fmt"
	"github.com/Briofy/fs-go/src/internal/contract/storage"
	"io"
	"os"
)

type StorageRepo struct {
	DirName string
}

func NewStorageRepo() storage.Repository {
	dirName := "storage"
	_ = os.MkdirAll(dirName, 0777)
	return &StorageRepo{
		DirName: dirName,
	}
}

func (s StorageRepo) Get(ctx context.Context, filePath string) (string, error) {
	return filePath, nil
}

func (s StorageRepo) GetBatch(ctx context.Context, filePath []string) ([]string, error) {
	return filePath, nil
}

func (s StorageRepo) Upload(ctx context.Context, reader io.Reader, fileName string) error {

	f, err := os.Create(fmt.Sprintf("%s/%s", s.DirName, fileName))
	if err != nil {
		return err
	}

	_, err = io.Copy(f, reader)
	if err != nil {
		return err
	}

	return nil
}
