package db

import (
	"go-rest-api/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(cfg *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	
	
	
	return &Db{db}
}
