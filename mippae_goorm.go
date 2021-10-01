package mippae_goorm

import (
	"fmt"
	"time"

	"github.com/renatormc/mippae_goorm/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectPostgres(host string, port string, user string, password string, name string) (*gorm.DB, error) {
	var db *gorm.DB
	str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, name, password)

	db, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	return db, nil
}

func ConnectSqlite(file string) (*gorm.DB, error) {
	var db *gorm.DB

	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseConn(db *gorm.DB) error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Tag{},
		&models.Device{},
		&models.DataSource{},
		&models.Call{},
		&models.CallPart{},
		&models.Chat{},
		&models.ChatParticipant{},
		&models.Contact{},
		&models.ContactEntry{},
		&models.Message{},
		&models.File{},
		&models.Sms{},
		&models.SmsPart{},
		&models.UserAccount{},
	)
}
