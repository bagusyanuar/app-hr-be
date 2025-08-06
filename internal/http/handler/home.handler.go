package handler

import (
	"github.com/bagusyanuar/app-hr-be/internal/config"
	"github.com/bagusyanuar/app-hr-be/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	Config *config.AppConfig
}

func NewHomeHandler(cfg *config.AppConfig) *HomeHandler {
	return &HomeHandler{
		Config: cfg,
	}
}

func (c *HomeHandler) Index(ctx *fiber.Ctx) error {
	return response.MakeAPIResponse(ctx, response.APIResponse[any]{
		Code:    fiber.StatusOK,
		Message: "welcome to app-hr-be",
		Data: &fiber.Map{
			"app_name":    c.Config.Viper.GetString("APP_NAME"),
			"app_version": c.Config.Viper.GetString("APP_VERSION"),
		},
	})
}
