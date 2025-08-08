package dto

import "github.com/bagusyanuar/app-hr-be/internal/domain/entity"

type BranchDTO struct {
	ID       string             `json:"id"`
	Name     string             `json:"name"`
	Address  *BranchAddressDTO  `json:"address"`
	Contacts []BranchContactDTO `json:"contacts"`
}

type BranchAddressDTO struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

type BranchContactDTO struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func ToBranch(data *entity.Branch) *BranchDTO {
	var address *BranchAddressDTO
	contacts := make([]BranchContactDTO, 0)

	if data.Address != nil {
		address = &BranchAddressDTO{
			ID:      data.Address.ID.String(),
			Address: data.Address.Address,
		}
	}

	for _, contact := range data.Contacts {
		c := &BranchContactDTO{
			ID:    contact.ID.String(),
			Type:  contact.Type,
			Value: contact.Value,
		}
		contacts = append(contacts, *c)
	}

	return &BranchDTO{
		ID:       data.ID.String(),
		Name:     data.Name,
		Address:  address,
		Contacts: contacts,
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
