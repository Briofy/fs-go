package file_manager_test

import (
	"context"
	"github.com/Briofy/fs-go/src/entity"
	AttachableMock "github.com/Briofy/fs-go/src/internal/mock/repository/attachable"
	StorageMock "github.com/Briofy/fs-go/src/internal/mock/repository/storage"
	"github.com/Briofy/fs-go/src/service/file_manager"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

type mock struct {
	mockStorageRepo    *StorageMock.MockRepository
	mockAttachableRepo *AttachableMock.MockRepository
}

func setup(t *testing.T) (file_manager.UseCase, mock, func()) {

	controller := gomock.NewController(t)
	mockSotrageRepo := StorageMock.NewMockRepository(controller)
	mockAttachableRepo := AttachableMock.NewMockRepository(controller)
	fs := file_manager.New(mockAttachableRepo, mockSotrageRepo)
	mocky := mock{
		mockAttachableRepo: mockAttachableRepo,
		mockStorageRepo:    mockSotrageRepo,
	}
	return fs, mocky, func() {
		controller.Finish()
	}
}

func TestGenerateSecret(t *testing.T) {
	fileManager, mocky, teardowm := setup(t)
	defer teardowm()
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	t.Run(
		"upload file",
		func(t *testing.T) {
			mocky.mockAttachableRepo.EXPECT().Create(ctx, gomock.Any()).Return(nil)
			mocky.mockStorageRepo.EXPECT().Upload(ctx, gomock.Any(), gomock.Any()).Return(nil)
			attachable := entity.Attachable{
				AttachableType:  "profile",
				AttachableField: "cover",
				AttachableID:    "1",
			}
			var file io.Reader
			fileMock := file_manager.File{
				File:     file,
				FileName: "somefilename",
			}
			err := fileManager.Upload(ctx, &attachable, fileMock)
			assert.Nil(t, err)
		},
	)

}
