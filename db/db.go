package db

import (
	"github.com/9thDuck/ecommerce-api.git/users"
	"github.com/9thDuck/ecommerce-api.git/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormConfig = gorm.Config{
	PrepareStmt: true,
}

func connect() (db *gorm.DB, err error) {
	dsn := utils.GetEnvOrThrow("POSTGRES_DB_DSN")
	db, err = gorm.Open(postgres.Open(dsn), &gormConfig)

	if err != nil {
		return nil, err
	}

	return db, nil
}

type DB struct {
	instance *gorm.DB
}

func Setup() *DB {
	dbInstance, err := connect()
	utils.LogFatalCustomError("failed to connect to db", err)

	db := DB{dbInstance}
	db.Migrate()

	return &db
}

func (db *DB) Migrate() {
	err := db.instance.AutoMigrate(&users.User{})
	utils.LogFatalCustomError("failed to migrate users table", err)
}
