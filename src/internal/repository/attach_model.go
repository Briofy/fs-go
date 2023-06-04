package repository

import (
	"github.com/fs-go/src/entity"
	"gorm.io/gorm"
)

type Attach struct {
	gorm.Model
	Path string
	// todo every thing with need
}

func (e *Attach) ToEntity() *entity.Attach {
	return &entity.Attach{
		ID:   e.ID,
		Path: e.Path,
	}
}

func mapAttachFromEntity(e *entity.Attach) *Attach {
	return &Attach{
		Path: e.Path,
	}
}
