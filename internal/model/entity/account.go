package entity

import "gorm.io/gorm"

type Account struct {
	ID         uint
	Nama       string `gorm:"type:varchar(50)"`
	NIK        string `gorm:"type:varchar(16)"`
	NoHp       string `gorm:"type:varchar(13)"`
	NoRekening string `gorm:"type:varchar(19)"`
	gorm.Model
}
