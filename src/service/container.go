package service

import (
	"context"
	"github.com/Briofy/fs-go/src/config"
	"github.com/Briofy/fs-go/src/internal/contract/attachable"
	"github.com/Briofy/fs-go/src/internal/contract/storage"
	"github.com/Briofy/fs-go/src/internal/repository"
	"github.com/Briofy/fs-go/src/pkg/db"
	databaseDriver "github.com/Briofy/fs-go/src/pkg/db/database"
	"github.com/Briofy/fs-go/src/pkg/storage/s3"
	"github.com/Briofy/fs-go/src/service/file_manager"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Container struct {
	db             *gorm.DB
	attachableRepo attachable.Repository
	storageRepo    storage.Repository
	file_manager.UseCase
}

func NewFile(ctx context.Context, config config.Config) (*Container, error) {
	var database *gorm.DB
	initDb, err := db.GetInitDb(databaseDriver.Driver(config.GetDatabaseDriver()))
	if err != nil {
		return nil, err
	}
	database, err = initDb(config)
	if err != nil {
		return nil, err
	}
	attachableRepo := repository.NewAttachableRepo(database)
	storageRepo := repository.NewStorageRepo()
	fileManager := file_manager.New(attachableRepo, storageRepo)
	return &Container{
		db:             database,
		attachableRepo: attachableRepo,
		storageRepo:    storageRepo,
		UseCase:        fileManager,
	}, nil
}

func NewS3(ctx context.Context, config config.Config) (*Container, error) {
	var database *gorm.DB
	initDb, err := db.GetInitDb(databaseDriver.Driver(config.GetDatabaseDriver()))
	if err != nil {
		return nil, err
	}
	database, err = initDb(config)
	if err != nil {
		return nil, err
	}
	attachableRepo := repository.NewAttachableRepo(database)
	s3StorageConnector := s3.NewS3(config)
	storageRepo := repository.NewS3StorageRepo(s3StorageConnector, config)
	fileManager := file_manager.New(attachableRepo, storageRepo)
	return &Container{
		db:             database,
		attachableRepo: attachableRepo,
		storageRepo:    storageRepo,
		UseCase:        fileManager,
	}, nil
}

func (c Container) Migrate() error {
	return repository.Migrate(c.db)
}
