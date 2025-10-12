package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	loggger *Logger
)

func Init() error{
	var err error

	db, err = InitializeSQLite()

	if err != nil{
		return fmt.Errorf("error initializing sqlite: %v", err)
	}
	return nil
}

func GetDb() *gorm.DB{
	return db
}

func GetLogger(p string) *Logger{
	logger := NewLogger(p)
	return logger
}