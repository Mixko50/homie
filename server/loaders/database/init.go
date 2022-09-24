package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"server/repository"
	"server/utils/config"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(mysql.Open(config.C.MySqlConfig), &gorm.Config{
		Logger: logger.New(
			log.New(
				os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel:                  logger.Info,
				Colorful:                  true,
				IgnoreRecordNotFoundError: false,
			}),
	})
	if err != nil {
		panic(err)
	}

	// * Initialize model migrations
	if config.C.AutoMigrate {
		if err := migrate(db); err != nil {
			panic("UNABLE TO MIGRATE GORM MODEL")
		}
	}

	DB = db
}

func migrate(db *gorm.DB) error {
	// * Migrate model
	if err := db.AutoMigrate(
		new(repository.Group),
	); err != nil {
		return err
	}

	return nil
}
