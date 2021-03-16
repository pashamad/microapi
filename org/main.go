package main

import (
	"github.com/pashamad/microapi/org/handler"
	pb "github.com/pashamad/microapi/org/proto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/logger"
)

var dbAddress = "postgresql://root:test¡@localhost:5432/onlife_business?sslmode=allow"

func main() {
	// Create service
	srv := service.New(
		service.Name("org"),
		service.Version("latest"),
	)

	// Connect to the database
	cfg, err := config.Get("database.biz.dsn", config.Secret(false))
	if err != nil {
		logger.Fatalf("Error loading config: %v", err)
	}
	addr := cfg.String(dbAddress)
	db, err := gorm.Open(postgres.Open(addr), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	if err := pb.RegisterOrgHandler(srv.Server(), &handler.Org{DB: db.Debug()}); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
