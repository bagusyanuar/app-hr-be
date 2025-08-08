package di

import (
	"github.com/bagusyanuar/app-hr-be/internal/config"
	"github.com/bagusyanuar/app-hr-be/internal/http/handler"
)

// dependency injection for handler package
type HandlerDI struct {
	Home   *handler.HomeHandler
	Auth   *handler.AuthHandler
	Branch *handler.BranchHandler
}

func InitializeDIHandler(cfg *config.AppConfig, diService *ServiceDI) *HandlerDI {
	return &HandlerDI{
		Home:   handler.NewHomeHandler(cfg),
		Auth:   handler.NewAuthHandler(diService.Auth, cfg),
		Branch: handler.NewBranchHandler(diService.Branch, cfg),
	}
}
