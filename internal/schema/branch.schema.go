package schema

import "github.com/bagusyanuar/app-hr-be/pkg/pagination"

type BranchSchema struct {
	Name string `json:"name" validate:"required"`
}

type BranchQuery struct {
	Param string `json:"param" query:"param"`
	pagination.QueryPagination
	pagination.QuerySort
}
