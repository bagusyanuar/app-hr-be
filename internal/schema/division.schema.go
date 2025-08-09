package schema

import "github.com/bagusyanuar/app-hr-be/pkg/pagination"

type DivisionSchema struct {
	Name string `json:"name" validate:"required"`
}

type DivisionQuery struct {
	Param string `json:"param" query:"param"`
	pagination.QueryPagination
	pagination.QuerySort
}
