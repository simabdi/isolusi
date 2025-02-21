package handler

import (
	"github.com/gofiber/fiber/v2"
	"isolusi/internal/helper"
	"isolusi/internal/model/request"
	"isolusi/internal/model/resource"
	"isolusi/internal/service"
)

type balanceHandler struct {
	balanceService service.BalanceService
}

func NewBalanceHandler(balanceService service.BalanceService) *balanceHandler {
	return &balanceHandler{balanceService}
}

func (h *balanceHandler) Tabung(ctx *fiber.Ctx) error {
	var input request.TransactionRequest
	err := ctx.BodyParser(&input)
	if err != nil {
		json := helper.JsonResponse(fiber.StatusUnprocessableEntity, "", err.Error(), nil)
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(json)
	}

	result, err := h.balanceService.Kredit(input)
	if err != nil {
		json := helper.JsonResponse(fiber.StatusBadRequest, "", err.Error(), nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(json)
	}

	json := helper.JsonResponse(fiber.StatusOK, "Anda berhasil menabung", "", resource.BalanceResource(result))
	return ctx.Status(fiber.StatusOK).JSON(json)
}

func (h *balanceHandler) Tarik(ctx *fiber.Ctx) error {
	var input request.TransactionRequest
	err := ctx.BodyParser(&input)
	if err != nil {
		json := helper.JsonResponse(fiber.StatusUnprocessableEntity, "", err.Error(), nil)
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(json)
	}

	result, err := h.balanceService.Debit(input)
	if err != nil {
		json := helper.JsonResponse(fiber.StatusBadRequest, "", err.Error(), nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(json)
	}

	json := helper.JsonResponse(fiber.StatusOK, "Anda berhasil menarik uang", "", resource.BalanceResource(result))
	return ctx.Status(fiber.StatusOK).JSON(json)
}

func (h *balanceHandler) CekSaldo(ctx *fiber.Ctx) error {
	result, err := h.balanceService.GetByNoRekening(ctx.Params("no_rekening"))
	if err != nil {
		json := helper.JsonResponse(fiber.StatusBadRequest, "", err.Error(), nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(json)
	}

	if result.ID == 0 {
		json := helper.JsonResponse(fiber.StatusBadRequest, "No rekening yang anda masukan tidak sah", "", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(json)
	}

	json := helper.JsonResponse(fiber.StatusOK, "", "", resource.BalanceResource(result))
	return ctx.Status(fiber.StatusOK).JSON(json)
}
