package config

import (
	"errors"
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
	if err = db.AutoMigrate(&entity.User{}, &entity.Answer{}, &entity.Question{}); err == nil && db.Migrator().HasTable(&entity.User{}) {
		if err := db.First(&entity.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			//Insert seed data
			seedData(db)
		}
	}

	return db
}

func seedData(db *gorm.DB) {
	db.Where("1 = 1").Delete(&entity.User{})
	db.Where("1 = 1").Delete(&entity.Answer{})
	db.Where("1 = 1").Delete(&entity.Question{})

	db.Create(&entity.User{Name: "Admin", Email: "admin@gmail.com", Password: "$2b$10$vqAX8n/aE7HTOnVD9r3ogeu4nT8eOf94jEZLBmNUc74wlrJhId7CW"})
	db.Create(&entity.User{Name: "Admin Dua", Email: "admin2@gmail.com", Password: "$2b$10$vqAX8n/aE7HTOnVD9r3ogeu4nT8eOf94jEZLBmNUc74wlrJhId7CW"})
	db.Create(&entity.Question{Question: "Kamu Nanya ?", UserID: 1})
	db.Create(&entity.Answer{Answer: "Nih Ya Aku kasih Tau", UserID: 1, QuestionID: 1})
}
