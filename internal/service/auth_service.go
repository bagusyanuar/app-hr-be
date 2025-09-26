package service

import (
	"context"

	"github.com/bagusyanuar/app-hr-be/internal/domain/schema"
)

type (
	AuthService interface {
		Login(ctx context.Context, schema schema.LoginSchema) error
	}
)
