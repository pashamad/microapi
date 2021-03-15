package main

import (
	"github.com/pashamad/microapi/orders/handler"
	pb "github.com/pashamad/microapi/orders/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("orders"),
		service.Version("latest"),
	)

	// Register handler
	if err := pb.RegisterOrdersHandler(srv.Server(), new(handler.Orders)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
