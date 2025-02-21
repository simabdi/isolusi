package request

type RegisterRequest struct {
	Nama string `json:"nama" validate:"required"`
	Nik  string `json:"nik" validate:"required,min=16,max=16"`
	NoHp string `json:"no_hp" validate:"required,min=11,max=13"`
}
