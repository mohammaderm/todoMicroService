package grpc

import (
	"context"

	"github.com/mohammaderm/todoMicroService/authService/internal/dto"
	"github.com/mohammaderm/todoMicroService/authService/internal/usecase"
	"github.com/mohammaderm/todoMicroService/authService/pkg/logger"
	"github.com/mohammaderm/todoMicroService/authService/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthServer struct {
	logger      logger.Logger
	authUsecase usecase.AuthUscase
	proto.UnimplementedAuthServiceServer
}

func New(logger logger.Logger, authUseCase usecase.AuthUscase) proto.AuthServiceServer {
	return &AuthServer{
		logger:      logger,
		authUsecase: authUseCase,
	}
}

func (a *AuthServer) RefreshToken(ctx context.Context, req *proto.RefreshTokenRequest) (*proto.RefreshTokenRespons, error) {
	refresh := dto.RefreshReq{
		RefreshToken: req.RefreshToken,
	}
	pairToken, err := a.authUsecase.RefreshToken(ctx, refresh)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	res := &proto.RefreshTokenRespons{
		PairToken: &proto.PairToken{
			AccessToken:  pairToken.AccessToken,
			RefreshToken: pairToken.RefreshToken,
		},
	}
	return res, nil
}

func (a *AuthServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginRespons, error) {
	user := dto.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	}
	loginRes, err := a.authUsecase.Login(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	// respons
	res := &proto.LoginRespons{
		User: &proto.User{
			Id:        loginRes.User.Id,
			Email:     loginRes.User.Email,
			Usernae:   loginRes.User.Username,
			Password:  "*",
			CreatedAt: timestamppb.New(loginRes.User.CreatedAt),
		},
		PairToken: &proto.PairToken{AccessToken: loginRes.AccessToken, RefreshToken: loginRes.RefreshToken},
	}
	return res, nil

}

func (a *AuthServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterRespons, error) {
	user := dto.RegisterReq{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	result, err := a.authUsecase.Register(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// respons
	res := &proto.RegisterRespons{
		Message: result.Message,
		Error:   result.Error,
	}
	return res, nil
}
