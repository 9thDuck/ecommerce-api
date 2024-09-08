package db

import (
	"github.com/9thDuck/ecommerce-api.git/common"
	"github.com/9thDuck/ecommerce-api.git/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormConfig = gorm.Config{
	PrepareStmt: true,
}

func connect() (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(common.APP_CONFIG.GetPostgresDbDsn()), &gormConfig)

	if err != nil {
		return nil, err
	}

	return db, nil
}

var Instance *gorm.DB

func SetupDbInstance(tables []interface{}) {
	db, err := connect()
	utils.LogFatalCustomError("failed to connect to db", err)

	migrate(db, tables)

	Instance = db
}

func migrate(db *gorm.DB, tables []interface{}) {
	for _, val := range tables {
		err := db.AutoMigrate(val)
		utils.LogFatalCustomError("failed to migrate users table", err)
	}
}
