package main

import (
	"bussinessenv/api/file/handler"
	"github.com/micro/go-micro"
	"log"

	proto "bussinessenv/api/file/proto/file"
	file "bussinessenv/srv/file/proto/file"
)

func main() {
	service := micro.NewService(
		micro.Name("com.lzqysoft.bussinessenv.api.file"),
		micro.Version("0.1"),
	)

	// Init to parse flags
	service.Init()

	// Register Handlers
	if err := proto.RegisterFileHandler(service.Server(), &handler.File{
		// Create Service Client
		Client: file.NewFileService("com.lzqysoft.bussinessenv.srv.file", service.Client()),
	}); err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
