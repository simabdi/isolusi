package service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"isolusi/internal/helper"
	"isolusi/internal/model/entity"
	"isolusi/internal/model/request"
	"isolusi/internal/repository"
	"strconv"
)

type AccountService interface {
	Store(input request.RegisterRequest) (entity.Account, error)
}

type accountService struct {
	accountRepository repository.AccountRepository
	balanceRepository repository.BalanceRepository
}

func NewAccountService(accountRepository repository.AccountRepository, balanceRepository repository.BalanceRepository) *accountService {
	return &accountService{accountRepository, balanceRepository}
}

func (s *accountService) Store(input request.RegisterRequest) (entity.Account, error) {
	check, err := s.accountRepository.CheckAccount(input.Nik, input.NoHp)
	if err != nil {
		return entity.Account{}, err
	}

	if check.ID != 0 {
		return entity.Account{}, errors.New("NIK atau No Hp sudah terdaftar")
	}

	payload := entity.Account{
		NIK:        input.Nik,
		Nama:       input.Nama,
		NoHp:       input.NoHp,
		NoRekening: strconv.Itoa(helper.GenerateRand(6)),
	}

	log.WithFields(log.Fields{
		"Payload Daftar": payload,
	}).Info("Payload daftar")

	save, err := s.accountRepository.Store(payload)
	if err != nil {
		return entity.Account{}, err
	}

	balance := entity.Balance{
		NoRekening: payload.NoRekening,
		Saldo:      0,
	}

	log.WithFields(log.Fields{
		"Payload Balance": balance,
	}).Info("Payload balance")

	_, err = s.balanceRepository.Store(balance)
	if err != nil {
		return entity.Account{}, err
	}

	return save, nil
}
