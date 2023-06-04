package repository

import (
	"context"
	"fmt"
	"github.com/fs-go/src/entity"
	"github.com/fs-go/src/internal/contract/attachable"
	"gorm.io/gorm"
)

type AttachableRepo struct {
	db *gorm.DB
}

func NewAttachableRepo(db *gorm.DB) attachable.Repository {
	return &AttachableRepo{
		db: db,
	}
}

func (a AttachableRepo) GetAttaches(ctx context.Context, attachable entity.Attachable) ([]*entity.Attach, error) {
	type Result struct {
		AttachId uint
	}
	var result []Result
	err := a.db.WithContext(ctx).Model(&Attachable{}).Select("attach_id").Where(fmt.Sprintf("attachable_type = '%s' AND attachable_field = '%s' AND attachable_id = '%s' ",
		attachable.AttachableType, attachable.AttachableField, attachable.AttachableID)).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	ids := mapSlice[Result, uint](result, func(result Result) uint { return result.AttachId })
	var attaches []*Attach
	err = a.db.WithContext(ctx).Where(" id in ?", ids).Find(&attaches).Error
	if err != nil {
		return nil, err
	}
	var attachEnts []*entity.Attach
	for _, attach := range attaches {
		attachEnts = append(attachEnts, attach.ToEntity())
	}
	return attachEnts, nil
}

func (a AttachableRepo) Create(ctx context.Context, attachable *entity.Attachable) error {
	err := a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		attachableModel := mapAttachableFromEntity(attachable)
		if err := tx.Create(attachableModel).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (a AttachableRepo) Delete(ctx context.Context, attachable entity.Attachable) error {
	//TODO implement me
	panic("implement me")
}

func mapSlice[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}
