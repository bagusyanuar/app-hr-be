package service

import (
	"context"

	"github.com/bagusyanuar/app-hr-be/internal/config"
	"github.com/bagusyanuar/app-hr-be/internal/domain/dto"
	"github.com/bagusyanuar/app-hr-be/internal/domain/entity"
	"github.com/bagusyanuar/app-hr-be/internal/repository"
	"github.com/bagusyanuar/app-hr-be/internal/schema"
)

type (
	BranchService interface {
		FindAll(ctx context.Context, queryParams *schema.BranchQuery) (*[]dto.BranchDTO, error)
		FindByID(ctx context.Context, id string) (*dto.BranchDTO, error)
		Create(ctx context.Context, schema *schema.BranchSchema) (*dto.BranchDTO, error)
	}

	branchServiceImpl struct {
		BranchRepository repository.BranchRepository
		Config           *config.AppConfig
	}
)

func NewBranchService(
	branchRepository repository.BranchRepository,
	cfg *config.AppConfig,
) BranchService {
	return &branchServiceImpl{
		BranchRepository: branchRepository,
		Config:           cfg,
	}
}

// Create implements BranchService.
func (b *branchServiceImpl) Create(ctx context.Context, schema *schema.BranchSchema) (*dto.BranchDTO, error) {
	data := &entity.Branch{
		Name: schema.Name,
	}

	branch, err := b.BranchRepository.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	branchDTO := dto.ToBranch(branch)
	return branchDTO, nil
}

// FindAll implements BranchService.
func (b *branchServiceImpl) FindAll(ctx context.Context, queryParams *schema.BranchQuery) (*[]dto.BranchDTO, error) {
	panic("unimplemented")
}

// FindByID implements BranchService.
func (b *branchServiceImpl) FindByID(ctx context.Context, id string) (*dto.BranchDTO, error) {
	panic("unimplemented")
}
