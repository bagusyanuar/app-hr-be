package di

import (
	"github.com/bagusyanuar/app-hr-be/internal/config"
	"github.com/bagusyanuar/app-hr-be/internal/http/handler"
)

// dependency injection for handler package
type HandlerDI struct {
	Home *handler.HomeHandler
}

func InitializeDIHandler(cfg *config.AppConfig) *HandlerDI {
	return &HandlerDI{
		Home: handler.NewHomeHandler(cfg),
	}
}
