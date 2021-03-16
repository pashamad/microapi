package main

import (
	"github.com/pashamad/microapi/receipt/handler"
	pb "github.com/pashamad/microapi/receipt/proto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/config"
	log "github.com/micro/micro/v3/service/logger"
)

var dbAddress = "postgresql://root:test@localhost:5432/onlife_business?sslmode=allow"

func main() {
	// Create service
	srv := service.New(
		service.Name("receipt"),
		service.Version("latest"),
	)

	// Connect to the database
	cfg, err := config.Get("database.biz.dsn", config.Secret(false))
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	addr := cfg.String(dbAddress)
	log.Debugf("DSN to database: %s", addr)
	db, err := gorm.Open(postgres.Open(addr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	if err := pb.RegisterReceiptHandler(srv.Server(), &handler.Receipt{DB: db.Debug()}); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
