package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"isolusi/internal/model/entity"
)

type AccountRepository interface {
	Store(data entity.Account) (entity.Account, error)
	CheckAccount(nik, noHp string) (entity.Account, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *accountRepository {
	return &accountRepository{db}
}

func (r *accountRepository) Store(data entity.Account) (entity.Account, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.Account{}, err
	}

	r.db.Logger = logger.Default.LogMode(logger.Info)
	return data, nil
}

func (r *accountRepository) CheckAccount(nik string, noHp string) (entity.Account, error) {
	var account entity.Account
	err := r.db.Where("nik = ?", nik).Or("no_hp = ?", noHp).Find(&account).Error
	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}
