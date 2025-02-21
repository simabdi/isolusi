package route

import (
	"isolusi/internal/config"
	"isolusi/internal/handler"
	"isolusi/internal/repository"
	"isolusi/internal/service"
)

func Apis() {
	db := config.Connection()

	balanceRepository := repository.NewBalanceRepository(db)
	balanceService := service.NewBalanceRepository(balanceRepository)
	balanceHandler := handler.NewBalanceHandler(balanceService)

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository, balanceRepository)
	accountHandler := handler.NewAccountHandler(accountService)

	app := SetupApi()
	app.Post("/daftar", accountHandler.Daftar)
	app.Post("/tabung", balanceHandler.Tabung)
	app.Post("/tarik", balanceHandler.Tarik)
	app.Get("/saldo/:no_rekening", balanceHandler.CekSaldo)

	app.Listen(":1453")
}
