package service

import (
	"context"

	"github.com/bagusyanuar/app-hr-be/internal/config"
	"github.com/bagusyanuar/app-hr-be/internal/domain/dto"
	"github.com/bagusyanuar/app-hr-be/internal/domain/entity"
	"github.com/bagusyanuar/app-hr-be/internal/repository"
	"github.com/bagusyanuar/app-hr-be/internal/schema"
	"github.com/bagusyanuar/app-hr-be/pkg/pagination"
)

type (
	BranchService interface {
		FindAll(ctx context.Context, queryParams *schema.BranchQuery) (*[]dto.BranchDTO, *pagination.PaginationMeta, error)
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

	address := schema.Address
	contacts := make([]entity.BranchContact, 0)
	for _, contact := range schema.Contacts {
		c := entity.BranchContact{
			Type:  contact.Type,
			Name:  contact.Name,
			Value: contact.Value,
		}
		contacts = append(contacts, c)
	}

	data := entity.Branch{
		Name: schema.Name,
		Address: &entity.BranchAddress{
			Address: address,
		},
		Contacts: contacts,
	}

	branch, err := b.BranchRepository.Create(ctx, &data)
	if err != nil {
		return nil, err
	}

	branchDTO := dto.ToBranch(branch)
	return branchDTO, nil
}

// FindAll implements BranchService.
func (b *branchServiceImpl) FindAll(ctx context.Context, queryParams *schema.BranchQuery) (*[]dto.BranchDTO, *pagination.PaginationMeta, error) {
	branches, pagination, err := b.BranchRepository.FindAll(ctx, queryParams)
	if err != nil {
		return &[]dto.BranchDTO{}, pagination, err
	}

	data := dto.ToBranches(branches)
	return &data, pagination, nil
}

// FindByID implements BranchService.
func (b *branchServiceImpl) FindByID(ctx context.Context, id string) (*dto.BranchDTO, error) {
	panic("unimplemented")
}
