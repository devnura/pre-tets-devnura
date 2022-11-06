package config

import (
	"fmt"
	"os"
	"time"

	"github.com/devnura/pre-tets-devnura/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB() *gorm.DB {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// dns config
	dns := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=UTF8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	// mysql config
	mysqlConfig := mysql.Config{
		DSN:                       dns,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}

	// gorm config
	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
		DryRun:                 false,
		PrepareStmt:            true,
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig)
	if err != nil {
		panic(err)
	}

	// connection pool config
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	// migrate table
	db.AutoMigrate(
		&entity.User{},
		&entity.Answer{},
		&entity.Question{},
	)

	seedData(db)

	return db
}

func seedData(db *gorm.DB) {

	// db.Where("1 = 1").Delete(&repository.Role{})
	// db.Where("1 = 1").Delete(&repository.User{})

	// db.Create(&repository.User{Username: "tirmizee", Password: "123", Email: "tirmizee@hotmail.com", FirstName: "pratya", LastName: "yeekhaday"})
	// db.Create(&repository.User{Username: "kiskdifw", Password: "123", Email: "kiskdifw@hotmail.com", FirstName: "poikue", LastName: "poiloipuy"})
	// db.Create(&repository.Role{Code: "R001", Name: "admin", Desc: "admin"})
	// db.Create(&repository.Role{Code: "R002", Name: "user", Desc: "user"})
}
