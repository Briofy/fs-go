package attachable

import (
	"context"
	"github.com/fs-go/src/entity"
)

type Reader interface {
	GetAttaches(ctx context.Context, attachable entity.Attachable) ([]*entity.Attach, error)
}

type Writer interface {
	Create(ctx context.Context, e *entity.Attachable) error
	Delete(ctx context.Context, attachable entity.Attachable) error
}

//go:generate mockgen -destination=../../mock/repository/attachable/attachable_mock.go -package=AttachableMock -self_package=github.com/fs-go/src/internal/contract/attachable github.com/fs-go/src/internal/contract/attachable Repository

type Repository interface {
	Reader
	Writer
}
