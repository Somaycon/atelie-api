package config

import (
	"os"

	"github.com/Somaycon/atelie-api/schemas"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/atelie.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("Database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.Errorf("Error creating database directory: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil{
		logger.Errorf("Sql initialization error: %v", err)
		return nil, err
	}
	// Ensure all schema models are migrated so their tables are created
	err = db.AutoMigrate(&schemas.Products{}, &schemas.Materials{}, &schemas.MaterialsList{})
	if err != nil{
		logger.Errorf("Sqlite AutoMigrate error: %v", err)
		return nil, err
	}
	return db, nil
}