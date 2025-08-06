package di

import (
	"github.com/bagusyanuar/app-hr-be/internal/config"
	"github.com/bagusyanuar/app-hr-be/internal/service"
)

type ServiceDI struct {
	Auth service.AuthService
}

func InitializeDIService(cfg *config.AppConfig, diRepository *RepositoryDI) *ServiceDI {
	return &ServiceDI{
		Auth: service.NewAuthService(diRepository.User, cfg),
	}
}
