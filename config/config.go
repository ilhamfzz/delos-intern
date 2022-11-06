package config

import (
	"delos-intern/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
	)
	var err error
	model.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return model.DB
}

func InitDatabaseTest() *gorm.DB {
	dbHost_test := os.Getenv("DB_HOST_TEST")
	dbUser_test := os.Getenv("DB_USER_TEST")
	dbPass_test := os.Getenv("DB_PASS_TEST")
	dbName_test := os.Getenv("DB_NAME_TEST")
	dbPort_test := os.Getenv("DB_PORT_TEST")
	dsn_test := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost_test,
		dbUser_test,
		dbPass_test,
		dbName_test,
		dbPort_test,
	)
	var err error
	model.DB, err = gorm.Open(postgres.Open(dsn_test), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	model.DB.Migrator().DropTable(&model.Farm{}, &model.Pond{}, &model.Telemetry{})
	model.DB.AutoMigrate(&model.Farm{}, &model.Pond{}, &model.Telemetry{})
	return model.DB
}
