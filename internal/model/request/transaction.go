package request

type TransactionRequest struct {
	NoRekening string `json:"no_rekening" validate:"required"`
	Nominal    int    `json:"nominal" validate:"required"`
}
