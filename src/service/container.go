package service

import (
	"context"
	"github.com/fs-go/src/internal/contract/attachable"
	"github.com/fs-go/src/internal/contract/storage"
	"github.com/fs-go/src/internal/repository"
	"github.com/fs-go/src/service/file_manager"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type Container struct {
	db             *gorm.DB
	attachableRepo attachable.Repository
	storageRepo    storage.Repository
	FileManager    file_manager.UseCase
}

func New(ctx context.Context, dsn string) (*Container, error) {
	database, err := newPostgresConn(dsn)
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
		FileManager:    fileManager,
	}, nil
}

func newPostgresConn(dsn string) (*gorm.DB, error) {
	_ = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
	dbConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	if err != nil {
		log.Println("DB Connection error : ", err.Error())
		return nil, err
	}
	dbSql := stdlib.OpenDB(*dbConfig)
	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbSql,
	}), &gorm.Config{})
	gdb.Debug()
	if err != nil {
		return nil, err
	}

	return gdb, nil
}

func (c Container) Migrate() error {
	return repository.Migrate(c.db)
}
