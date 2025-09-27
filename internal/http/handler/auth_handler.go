package handler

import (
	"github.com/bagusyanuar/app-hr-be/internal/config"
	"github.com/bagusyanuar/app-hr-be/internal/domain/schema"
	"github.com/bagusyanuar/app-hr-be/internal/service"
	"github.com/bagusyanuar/app-hr-be/pkg/util"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService service.AuthService
	Config      *config.AppConfig
}

func NewAuthHandler(
	authService service.AuthService,
	config *config.AppConfig,
) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		Config:      config,
	}
}

func (c *AuthHandler) Login(ctx *fiber.Ctx) error {
	request := new(schema.LoginSchema)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "invalid body json",
		})
	}

	messages, err := util.Validate(c.Config.Validator, request)

	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"code":    fiber.StatusUnprocessableEntity,
			"message": fiber.ErrUnprocessableEntity,
			"errors":  messages,
		})
	}

	schema := *request

	accessToken, refreshToken, err := c.AuthService.Login(ctx.UserContext(), schema)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully login",
		"data": map[string]string{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}
