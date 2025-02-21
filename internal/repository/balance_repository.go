package repository

import (
	"gorm.io/gorm"
	"isolusi/internal/model/entity"
)

type BalanceRepository interface {
	Store(data entity.Balance) (entity.Balance, error)
	Update(data entity.Balance) (entity.Balance, error)
	Transaction(data entity.BalanceTransaction) error
	GetByNoRekening(noRekening string) (entity.Balance, error)
}

type balanceRepository struct {
	db *gorm.DB
}

func NewBalanceRepository(db *gorm.DB) *balanceRepository {
	return &balanceRepository{db}
}

func (r *balanceRepository) Store(data entity.Balance) (entity.Balance, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.Balance{}, err
	}

	return data, nil
}

func (r *balanceRepository) Update(data entity.Balance) (entity.Balance, error) {
	err := r.db.Save(&data).Error
	if err != nil {
		return entity.Balance{}, err
	}

	return data, nil
}

func (r *balanceRepository) GetByNoRekening(noRekening string) (entity.Balance, error) {
	var balance entity.Balance
	err := r.db.Where("no_rekening = ?", noRekening).Find(&balance).Error
	if err != nil {
		return entity.Balance{}, err
	}

	return balance, nil
}

func (r *balanceRepository) Transaction(data entity.BalanceTransaction) error {
	err := r.db.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}
