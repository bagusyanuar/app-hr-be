package http

import (
	"github.com/bagusyanuar/app-hr-be/internal/config"
	"github.com/bagusyanuar/app-hr-be/internal/di"
)

func NewRouter(cfg *config.AppConfig, handler *di.HandlerDI) {
	app := cfg.App

	app.Get("/", handler.Home.Index)
	app.Post("/auth/login", handler.Auth.Login)

	branch := app.Group("/branch")
	branch.Get("/", handler.Branch.FindAll)
	branch.Post("/", handler.Branch.Create)
}
