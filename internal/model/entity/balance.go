package entity

type Balance struct {
	ID         uint
	NoRekening string `gorm:"type:varchar(19)"`
	Saldo      int    `gorm:"type:int"`
}
