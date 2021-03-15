package main

import (
	"receipt/handler"
	pb "receipt/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("receipt"),
		service.Version("latest"),
	)

	// Register handler
	if err := pb.RegisterReceiptHandler(srv.Server(), new(handler.Receipt)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
