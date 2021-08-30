package mippae_goorm

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/renatormc/mippae_goorm/models"
)

func TestDatabase(t *testing.T) {
	db, err := ConnectSqlite("C:\\temp\\db.db")
	if err != nil {
		panic(err)
	}
	var contact models.Contact
	// db.Where("id = ?", 17).First(&sms)
	db.Preload("Entries").First(&contact, 1)
	spew.Dump(contact.Entries)
	// var parts []models.SmsPart
	// db.Model(&sms).Association("Parts").Find(&parts)
	// spew.Dump(parts)
}
