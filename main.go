package main

import (
	"context"
	"fmt"
	"github.com/fs-go/src/entity"
	"github.com/fs-go/src/service"
	"github.com/fs-go/src/service/file_manager"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	dsn := "host=localhost user=sajjad password=sajjad123 dbname=fs port=5432 sslmode=require TimeZone=UTC"
	container, err := service.New(ctx, dsn)
	if err != nil {
		panic(err)
	}
	err = container.Migrate()
	if err != nil {
		fmt.Println("Migrate failed")
	}
	attachable := entity.Attachable{
		AttachableID:    "5",
		AttachableType:  "profile",
		AttachableField: "cover",
	}
	f, err := os.Open("/tmp/s.txt")
	if err != nil {
		fmt.Printf("read file Error | %s", err)

	}

	file := file_manager.File{
		File:     f,
		FileName: "file e man",
	}
	err = container.FileManager.Upload(ctx, &attachable, file)
	if err != nil {
		fmt.Printf("upload file file Error | %s", err)

	}
	fmt.Println(container.FileManager.GetLink(ctx, attachable))

}
