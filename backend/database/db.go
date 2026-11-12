package database

import (
	"log/slog"
	"pkl/backend/configs"
	"pkl/backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB(cfg configs.AppConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DatabaseFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// migrasi database
	if cfg.AutoMigrate {
		slog.Info("automigrate database")
		err = db.AutoMigrate(
			&models.User{},
		)

		if err != nil {
			panic(err)
		}
	}

	return db, nil
}
