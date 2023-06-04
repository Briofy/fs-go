package file_manager

import (
	"context"
	"github.com/fs-go/src/entity"
)

type UseCase interface {
	Upload(ctx context.Context, attachable *entity.Attachable, file File) error
	GetLink(ctx context.Context, attachable entity.Attachable) string
}
