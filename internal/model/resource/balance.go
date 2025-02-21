package resource

import (
	"isolusi/internal/model/entity"
	"isolusi/internal/model/formatter"
)

func BalanceResource(balance entity.Balance) formatter.BalanceFormatter {
	resource := formatter.BalanceFormatter{
		Saldo: balance.Saldo,
	}

	return resource
}
