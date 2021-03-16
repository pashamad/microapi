package config

import (
	"github.com/micro/micro/v3/service/config"
	log "github.com/micro/micro/v3/service/logger"
)

type Conf struct {
	Database database `json:"database"`
}

type database struct {
	DB  string `json:"db"`
	DSN string `json:"dsn"`
}

func GetConfig() Conf {

	v, _ := config.Get("database")
	d := &database{}
	err := v.Scan(d)
	if err != nil {
		log.Error(err.Error())
	}

	c := &Conf{
		Database: *d,
	}

	return *c
}
