package di

import (
	"github.com/bagusyanuar/app-hr-be/internal/config"
	"github.com/bagusyanuar/app-hr-be/internal/repository"
)

type RepositoryDI struct {
	User   repository.UserRepository
	Branch repository.BranchRepository
}

func InitializeDIRepository(cfg *config.AppConfig) *RepositoryDI {
	return &RepositoryDI{
		User:   repository.NewUserRepository(cfg.DB),
		Branch: repository.NewBranchRepository(cfg.DB),
	}
}
