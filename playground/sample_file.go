package main

import (
	"context"
	"fmt"
	"github.com/Briofy/fs-go/src/config"
	"github.com/Briofy/fs-go/src/entity"
	"github.com/Briofy/fs-go/src/service"
	"github.com/Briofy/fs-go/src/service/file_manager"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	var cfg config.SampleConfig
	container, err := service.NewFile(ctx, cfg)
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
	err = container.Upload(ctx, &attachable, file)
	if err != nil {
		fmt.Printf("upload file file Error | %s", err)

	}
	fmt.Println(container.GetLink(ctx, attachable))

}
