package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/smquartz/redeye/handler"
	"github.com/smquartz/redeye/subscriber"

	example "github.com/smquartz/redeye/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("video.quartz.famethyst.srv.redeye"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("topic.video.quartz.famethyst.srv.redeye", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("topic.video.quartz.famethyst.srv.redeye", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
