package mippae_goorm

import (
	"fmt"
	"testing"

	"github.com/renatormc/mippae_goorm/models"
)

func TestDatabase(t *testing.T) {
	db, err := ConnectSqlite("G:\\laudos\\trabalhando\\B\\midia\\.report\\db.db")
	if err != nil {
		panic(err)
	}
	var chats []models.Chat
	// db.Where("id = ?", 17).First(&sms)
	db.Where("n_messages < ?", 10).Find(&chats)
	for _, chat := range chats {
		fmt.Println(chat.NMessages)
	}
	// spew.Dump(contact.Entries)
	// var parts []models.SmsPart
	// db.Model(&sms).Association("Parts").Find(&parts)
	// spew.Dump(parts)
}
