package repository

import (
	"context"
	"fmt"

	"github.com/bagusyanuar/app-hr-be/internal/domain/entity"
	"github.com/bagusyanuar/app-hr-be/internal/schema"
	"github.com/bagusyanuar/app-hr-be/pkg/pagination"
	"gorm.io/gorm"
)

type (
	BranchRepository interface {
		FindAll(ctx context.Context, queryParams *schema.BranchQuery) ([]entity.Branch, *pagination.PaginationMeta, error)
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

	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&branch).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return branch, nil
}

// FindAll implements BranchRepository.
func (b *branchRepositoryImpl) FindAll(ctx context.Context, queryParams *schema.BranchQuery) ([]entity.Branch, *pagination.PaginationMeta, error) {
	tx := b.DB.WithContext(ctx)
	baseQuery := b.baseQuery(tx, queryParams)

	var totalRows int64
	if err := baseQuery.
		Model(&entity.Branch{}).
		Count(&totalRows).Error; err != nil {
		return []entity.Branch{}, nil, err
	}

	var data []entity.Branch

	sortFieldMap := map[string]string{
		"name": "name",
	}
	sort := pagination.GetSortField(queryParams.Sort, "name", sortFieldMap)
	order := pagination.GetOrder(queryParams.Order)
	if err := baseQuery.
		Scopes(
			pagination.SortScope(sort, order),
			pagination.Paginate(tx, queryParams.Page, queryParams.PageSize),
		).
		Find(&data).Error; err != nil {
		return []entity.Branch{}, nil, err
	}

	pagination := pagination.MakePagination(queryParams.Page, queryParams.PageSize, totalRows)
	return data, &pagination, nil
}

// FindByID implements BranchRepository.
func (b *branchRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Branch, error) {
	panic("unimplemented")
}

func (b *branchRepositoryImpl) baseQuery(tx *gorm.DB, queryParams *schema.BranchQuery) *gorm.DB {
	param := fmt.Sprintf("%%%s%%", queryParams.Param)

	return tx.
		Preload("Address").
		Preload("Contacts").
		Scopes(
			b.filterByParam(param),
		)
}

func (b *branchRepositoryImpl) filterByParam(param string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if param == "" {
			return tx
		}
		return tx.
			Where("name ILIKE ?", param)
	}
}
