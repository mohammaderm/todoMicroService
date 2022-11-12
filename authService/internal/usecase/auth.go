package usecase

import (
	"context"
	"errors"

	"github.com/mohammaderm/authService/internal/dto"
	"github.com/mohammaderm/authService/internal/models"
	"github.com/mohammaderm/authService/internal/repository"
	"github.com/mohammaderm/authService/pkg/jwt"
	"github.com/mohammaderm/authService/pkg/logger"
	"github.com/mohammaderm/authService/pkg/validator"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrValidationFailed = errors.New("input is not valid")
	ErrWrongCredentials = errors.New("wrong password")
)

type AuthUscase interface {
	Register(ctx context.Context, req dto.RegisterReq) (dto.RegisterRes, error)
	Login(ctx context.Context, req dto.LoginReq) (dto.PairToken, error)
	// refreshToken()
}

type UseCase struct {
	userRepo repository.UserRepository
	logger   logger.Logger
	jwt      jwt.JwtInterface
}

func New(userRepo repository.UserRepository, logger logger.Logger, jwt jwt.JwtInterface) AuthUscase {
	return &UseCase{
		userRepo: userRepo,
		logger:   logger,
		jwt:      jwt,
	}
}

func (u *UseCase) Login(ctx context.Context, req dto.LoginReq) (dto.PairToken, error) {
	err := validator.AuthRequest(ctx, req)
	if err != nil {
		return dto.PairToken{}, ErrValidationFailed
	}
	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.PairToken{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.PairToken{}, ErrWrongCredentials
	}
	pairToken, err := u.jwt.GeneratePairToken(user.Id, user.Email)
	if err != nil {
		return dto.PairToken{}, err
	}
	return dto.PairToken{
		AccessToken:  pairToken["accessToken"],
		RefreshToken: pairToken["refreshToken"],
	}, nil
}

func (u *UseCase) Register(ctx context.Context, req dto.RegisterReq) (dto.RegisterRes, error) {
	err := validator.AuthRequest(ctx, req)
	if err != nil {
		return dto.RegisterRes{
			Message: "user request is not valid",
			Error:   true,
		}, ErrValidationFailed
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterRes{
			Message: "internal error",
			Error:   true,
		}, err
	}
	var user models.User
	user.Email = req.Email
	user.Username = req.Username
	user.Password = string(hashPassword)
	err = u.userRepo.Create(ctx, &user)
	if err != nil {
		return dto.RegisterRes{
			Message: "failed to create user",
			Error:   true,
		}, err
	}
	return dto.RegisterRes{
		Message: "user created succesfully",
		Error:   false,
	}, nil

}
