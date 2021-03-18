package main

import (
	"github.com/micro/micro/v3/service"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/pashamad/microapi/dbconn"
	"github.com/pashamad/microapi/org/handler"
	pb "github.com/pashamad/microapi/org/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("org"),
		service.Version("latest"),
	)

	// Get db connection
	db, err := dbconn.GetConn("biz")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Register handler
	if err := pb.RegisterOrgHandler(srv.Server(), &handler.Org{DB: db.Debug()}); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
