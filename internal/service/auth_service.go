package service

import (
	"context"
	"time"

	"github.com/bagusyanuar/app-hr-be/internal/config"
	"github.com/bagusyanuar/app-hr-be/internal/domain/entity"
	"github.com/bagusyanuar/app-hr-be/internal/domain/schema"
	"github.com/bagusyanuar/app-hr-be/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

type (
	AuthService interface {
		Login(ctx context.Context, schema schema.LoginSchema) (accessToken string, refreshToken string, err error)
	}

	authServiceImpl struct {
		UserRepository repository.UserRepository
		Config         *config.AppConfig
	}
)

func NewAuthService(
	userRepository repository.UserRepository,
	config *config.AppConfig,
) AuthService {
	return &authServiceImpl{
		UserRepository: userRepository,
		Config:         config,
	}
}

// Login implements AuthService.
func (a *authServiceImpl) Login(ctx context.Context, schema schema.LoginSchema) (accessToken string, refreshToken string, err error) {
	email := schema.Email

	user, err := a.UserRepository.FindByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	accessToken, err = a.createAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = a.createRefreshToken(user)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, err
}

func (a *authServiceImpl) createAccessToken(user *entity.User) (string, error) {
	JWTSignInMethod := jwt.SigningMethodHS256
	exp := time.Now().Add(time.Minute * time.Duration(a.Config.JWT.Expiration))
	claims := jwt.RegisteredClaims{
		Issuer:    a.Config.JWT.Issuer,
		ExpiresAt: jwt.NewNumericDate(exp),
		Subject:   user.ID.String(),
	}
	accessToken := jwt.NewWithClaims(JWTSignInMethod, claims)
	return accessToken.SignedString([]byte(a.Config.JWT.Secret))
}

func (a *authServiceImpl) createRefreshToken(user *entity.User) (string, error) {
	JWTSignInMethod := jwt.SigningMethodHS256
	exp := time.Now().Add(time.Hour * 24 * time.Duration(a.Config.JWT.ExpirationRefreh))
	claims := jwt.RegisteredClaims{
		Issuer:    a.Config.JWT.Issuer,
		ExpiresAt: jwt.NewNumericDate(exp),
		Subject:   user.ID.String(),
	}
	refreshToken := jwt.NewWithClaims(JWTSignInMethod, claims)
	return refreshToken.SignedString([]byte(a.Config.JWT.SecretRefresh))
}
