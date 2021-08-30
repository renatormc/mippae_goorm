package mippae_goorm

import (
	"fmt"
	"testing"

	"github.com/renatormc/mippae_goorm/models"
)

func TestDatabase(t *testing.T) {
	db, err := ConnectSqlite("C:\\temp\\foo.db")
	if err != nil {
		panic(err)
	}
	var message models.Message
	db.First(&message)
	fmt.Println(message.Timestamp)
}
