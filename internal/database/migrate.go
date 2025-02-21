package database

import (
	"isolusi/internal/config"
	"isolusi/internal/model/entity"
)

func Migrate() {
	db := config.Connection()

	status := db.Migrator().HasTable(&entity.Account{})
	if status == false {
		db.AutoMigrate(&entity.Account{})
	}

	status = db.Migrator().HasTable(&entity.Balance{})
	if status == false {
		db.AutoMigrate(&entity.Balance{})
	}

	status = db.Migrator().HasTable(&entity.BalanceTransaction{})
	if status == false {
		db.AutoMigrate(&entity.BalanceTransaction{})
	}
}
