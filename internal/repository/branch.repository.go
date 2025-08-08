package repository

import (
	"context"

	"github.com/bagusyanuar/app-hr-be/internal/domain/entity"
	"github.com/bagusyanuar/app-hr-be/internal/schema"
	"gorm.io/gorm"
)

type (
	BranchRepository interface {
		FindAll(ctx context.Context, queryParams *schema.BranchQuery) ([]entity.Branch, error)
		FindByID(ctx context.Context, id string) (*entity.Branch, error)
		Create(ctx context.Context, branch *entity.Branch) (*entity.Branch, error)
	}

	branchRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewBranchRepository(db *gorm.DB) BranchRepository {
	return &branchRepositoryImpl{
		DB: db,
	}
}

// Create implements BranchRepository.
func (b *branchRepositoryImpl) Create(ctx context.Context, branch *entity.Branch) (*entity.Branch, error) {
	tx := b.DB.WithContext(ctx)
	if err := tx.Create(&branch).Error; err != nil {
		return nil, err
	}
	return branch, nil
}

// FindAll implements BranchRepository.
func (b *branchRepositoryImpl) FindAll(ctx context.Context, queryParams *schema.BranchQuery) ([]entity.Branch, error) {
	panic("unimplemented")
}

// FindByID implements BranchRepository.
func (b *branchRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Branch, error) {
	panic("unimplemented")
}
