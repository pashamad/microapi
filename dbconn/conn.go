package dbconn

import (
	"fmt"
	"github.com/micro/micro/v3/service/config"
	log "github.com/micro/micro/v3/service/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbAddress = "postgresql://root:test@localhost:5432/%s?sslmode=allow"

var dbMap = map[string]string{
	"app": "onlife",
	"biz": "onlife_business",
}

func GetConn(id string) (db *gorm.DB, err error) {
	// Check the id
	var dbAddrDef string
	if dbName, ok := dbMap[id]; !ok {
		log.Errorf("database conn with id %s is not registered", id)
		dbAddrDef = fmt.Sprintf(dbAddress, dbName)
	}

	// Connect to the database
	cfg, err := config.Get(fmt.Sprintf("database.%s.dsn", id), config.Secret(false))
	if err != nil {
		log.Errorf("Config value for database conn \"%s\" not found: %v", id, err)
		return nil, err
	}

	addr := cfg.String(dbAddrDef)
	log.Debugf("DSN to database: %s", addr)
	// @todo reuse connections
	db, err = gorm.Open(postgres.Open(addr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	// @todo configure migrations
	//if err := db.AutoMigrate(); err != nil {
	//	log.Fatalf("Error migrating database: %v", err)
	//}

	return db, nil
}
