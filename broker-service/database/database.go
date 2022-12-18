package database

import (
	"broker/data"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitDb() (*Database, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		getenv("MYSQL_USER", "user_app"),
		getenv("MYSQL_PASSWORD", "password"),
		getenv("MYSQL_HOST", "localhost"),
		getenv("MYSQL_DATABASE", "images_db"),
	)

	fmt.Printf("%s", dsn)
	config := createGornConfig()

	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&data.ImagesDescription{}, &data.Tag{}, &data.User{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	sqlDB.SetConnMaxLifetime(time.Second)
	return &Database{
		Gorm: db,
	}, nil
}

func createGornConfig() *gorm.Config {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	return &gorm.Config{
		PrepareStmt: false,
		Logger:      newLogger,
	}
}

type Database struct {
	Gorm *gorm.DB
}

func getenv(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}

	return v
}
