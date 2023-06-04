package repository

import (
	"github.com/fs-go/src/entity"
	"gorm.io/gorm"
)

type Attachable struct {
	gorm.Model
	Attach          *Attach `json:"attach"`
	AttachId        uint    `json:"attachId"`
	AttachableType  string  `json:"attachable_type"`
	AttachableField string  `json:"attachable_field"`
	AttachableID    string  `json:"attachable_id"`
}

func (e *Attachable) ToEntity() *entity.Attachable {
	var attach *entity.Attach
	if e.Attach != nil {
		attach = e.Attach.ToEntity()
	}
	return &entity.Attachable{
		ID:              e.ID,
		Attach:          attach,
		AttachId:        e.AttachId,
		AttachableType:  e.AttachableType,
		AttachableField: e.AttachableField,
		AttachableID:    e.AttachableID,
	}
}

func mapAttachableFromEntity(e *entity.Attachable) *Attachable {
	var attach *Attach
	if e.Attach != nil {
		attach = mapAttachFromEntity(e.Attach)
	}
	return &Attachable{
		Attach:          attach,
		AttachId:        e.AttachId,
		AttachableType:  e.AttachableType,
		AttachableField: e.AttachableField,
		AttachableID:    e.AttachableID,
	}
}
