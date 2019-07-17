package main

import (
	"bussinessenv/srv/file/handler"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	file "bussinessenv/srv/file/proto/file"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.lzqysoft.bussinessenv.srv.file"),
		micro.Version("0.1"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	//file.RegisterExampleHandler(service.Server(), new(handler.Example))
	if err := file.RegisterFileHandler(service.Server(), new(handler.File)); err != nil {
		log.Fatal(err)
	}
	// Register Struct as Subscriber
	//micro.RegisterSubscriber("com.lzqysoft.bussinessenv.srv.file", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("com.lzqysoft.bussinessenv.srv.file", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
