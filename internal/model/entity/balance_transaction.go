package entity

import "gorm.io/gorm"

type BalanceTransaction struct {
	ID             uint
	NoRekening     string `gorm:"type:varchar(19)"`
	JenisTransaksi string `gorm:"type:varchar(10)"`
	Nominal        int    `gorm:"type:int"`
	gorm.Model
}
