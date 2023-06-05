package file_manager

import (
	"github.com/Briofy/fs-go/src/internal/contract/attachable"
	"github.com/Briofy/fs-go/src/internal/contract/storage"
	"io"
)

type FileManager struct {
	attachableRepo attachable.Repository
	storageRepo    storage.Repository
}

type File struct {
	File     io.Reader
	FileName string
}

func New(attachableRepo attachable.Repository, storageRepo storage.Repository) UseCase {
	return &FileManager{
		attachableRepo: attachableRepo,
		storageRepo:    storageRepo,
	}
}
