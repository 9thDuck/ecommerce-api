package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(DSN string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(DSN), &gormConfig)

	if err != nil {
		return nil, err
	}

	return db, nil
}

var gormConfig = gorm.Config{
	PrepareStmt: true,
}
