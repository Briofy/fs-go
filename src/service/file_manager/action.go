package file_manager

import (
	"context"
	"github.com/Briofy/fs-go/src/entity"
	"github.com/Briofy/fs-go/src/pkg/utils"
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

func (f FileManager) GetLink(ctx context.Context, attachable entity.Attachable) (string, error) {
	attaches, err := f.attachableRepo.GetAttaches(ctx, attachable)
	if err != nil {
		return "", err
	}
	link, err := f.storageRepo.Get(ctx, attaches[len(attaches)-1].Path)
	if err != nil {
		return "", err
	}
	return link, err
}

func (f FileManager) GetLinks(ctx context.Context, attachable entity.Attachable) ([]string, error) {
	attaches, err := f.attachableRepo.GetAttaches(ctx, attachable)
	if err != nil {
		return nil, err
	}
	filePaths := utils.MapSlice[*entity.Attach, string](attaches, func(attach *entity.Attach) string { return attach.Path })
	links, err := f.storageRepo.GetBatch(ctx, filePaths)
	if err != nil {
		return nil, err
	}
	return links, err
}
