package handler

import (
	"github.com/gofiber/fiber/v2"
	"isolusi/internal/helper"
	"isolusi/internal/model/request"
	"isolusi/internal/model/resource"
	"isolusi/internal/service"
	"isolusi/validation"
)

type accountHandler struct {
	accountService service.AccountService
}

func NewAccountHandler(accountService service.AccountService) *accountHandler {
	return &accountHandler{accountService}
}

func (h *accountHandler) Daftar(ctx *fiber.Ctx) error {
	var input request.RegisterRequest

	err := ctx.BodyParser(&input)
	if err != nil {
		json := helper.JsonResponse(fiber.StatusUnprocessableEntity, "", validation.Validate(input), nil)
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(json)
	}

	save, err := h.accountService.Store(input)
	if err != nil {
		json := helper.JsonResponse(fiber.StatusBadRequest, "", err.Error(), nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(json)
	}

	json := helper.JsonResponse(fiber.StatusOK, "Pendafaran akun berhasil", "", resource.AccountResource(save))
	return ctx.Status(fiber.StatusOK).JSON(json)
}
