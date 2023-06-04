package file_manager

import (
	"context"
	"fmt"
	"github.com/fs-go/src/entity"
)

func (f FileManager) Upload(ctx context.Context, attachable *entity.Attachable, file File) error {
	// todo get file checksum

	err := f.storageRepo.Upload(ctx, file.File, file.FileName)
	if err != nil {
		return err
	}
	attach := entity.Attach{
		Path: file.FileName,
	}
	attachable.Attach = &attach
	err = f.attachableRepo.Create(ctx, attachable)
	if err != nil {
		return err
	}
	return nil
}

func (f FileManager) GetLink(ctx context.Context, attachable entity.Attachable) string {
	attaches, err := f.attachableRepo.GetAttaches(ctx, attachable)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", attaches)
}
