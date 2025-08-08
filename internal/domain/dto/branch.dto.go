package dto

import "github.com/bagusyanuar/app-hr-be/internal/domain/entity"

type BranchDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToBranch(data *entity.Branch) *BranchDTO {
	return &BranchDTO{
		ID:   data.ID.String(),
		Name: data.Name,
	}
}

func ToBranchs(data []entity.Branch) []BranchDTO {
	branchs := make([]BranchDTO, 0)
	for _, datum := range data {
		branch := *ToBranch(&datum)
		branchs = append(branchs, branch)
	}
	return branchs
}
