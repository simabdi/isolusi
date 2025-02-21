package service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"isolusi/internal/model/entity"
	"isolusi/internal/model/request"
	"isolusi/internal/repository"
)

type BalanceService interface {
	GetByNoRekening(noRekening string) (entity.Balance, error)
	Debit(input request.TransactionRequest) (entity.Balance, error)
	Kredit(input request.TransactionRequest) (entity.Balance, error)
}

type balanceService struct {
	balanceRepository repository.BalanceRepository
}

func NewBalanceRepository(balanceRepository repository.BalanceRepository) *balanceService {
	return &balanceService{balanceRepository}
}

func (s *balanceService) GetByNoRekening(noRekening string) (entity.Balance, error) {
	result, err := s.balanceRepository.GetByNoRekening(noRekening)
	if err != nil {
		return entity.Balance{}, err
	}

	return result, nil
}

func (s *balanceService) Debit(input request.TransactionRequest) (entity.Balance, error) {
	balance, err := s.balanceRepository.GetByNoRekening(input.NoRekening)
	if err != nil {
		return entity.Balance{}, err
	}

	if balance.ID == 0 {
		return entity.Balance{}, errors.New("No rekening tidak ditemukan")
	}

	balance.Saldo -= input.Nominal

	save, err := s.balanceRepository.Update(balance)
	if err != nil {
		return entity.Balance{}, err
	}

	log.WithFields(log.Fields{
		"Balance Input": input,
		"Balance Debit": balance,
	}).Info("Update balance tarik saldo")

	transaction := entity.BalanceTransaction{
		NoRekening:     balance.NoRekening,
		Nominal:        input.Nominal,
		JenisTransaksi: "Debit",
	}

	err = s.balanceRepository.Transaction(transaction)
	if err != nil {
		return entity.Balance{}, err
	}

	return save, nil
}

func (s *balanceService) Kredit(input request.TransactionRequest) (entity.Balance, error) {
	balance, err := s.balanceRepository.GetByNoRekening(input.NoRekening)
	if err != nil {
		return entity.Balance{}, err
	}

	if balance.ID == 0 {
		return entity.Balance{}, errors.New("No rekening tidak ditemukan")
	}

	balance.Saldo += input.Nominal
	save, err := s.balanceRepository.Update(balance)
	if err != nil {
		return entity.Balance{}, err
	}

	log.WithFields(log.Fields{
		"Balance Input":  input,
		"Balance Kredit": balance,
	}).Info("Update balance tabung saldo")

	transaction := entity.BalanceTransaction{
		NoRekening:     balance.NoRekening,
		Nominal:        input.Nominal,
		JenisTransaksi: "Kredit",
	}

	err = s.balanceRepository.Transaction(transaction)
	if err != nil {
		return entity.Balance{}, err
	}

	return save, nil
}
