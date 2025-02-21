package resource

import (
	"isolusi/internal/model/entity"
	"isolusi/internal/model/formatter"
)

func AccountResource(account entity.Account) formatter.AccountFormatter {
	resource := formatter.AccountFormatter{
		NoRekening: account.NoRekening,
	}

	return resource
}
