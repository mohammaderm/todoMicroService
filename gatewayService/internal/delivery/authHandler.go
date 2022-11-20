package delivery

import (
	"context"
	"errors"

	"github.com/mohammaderm/todoMicroService/authService/proto"
	"github.com/mohammaderm/todoMicroService/gatewayService/config"
	"github.com/mohammaderm/todoMicroService/gatewayService/pkg/logger"

	"net/http"
	"time"

	"github.com/mohammaderm/todoMicroService/gatewayService/internal/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthHandler struct {
	HandlerHelper
	cfg *config.Service
}

type AuthHandlerContract interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler(logger logger.Logger, cfg *config.Service) AuthHandlerContract {
	return &AuthHandler{
		HandlerHelper: HandlerHelper{
			logger: logger,
		},
		cfg: cfg,
	}
}

// @summary     LOGIN USER
// @description login User with "email" and "password" to get token for authentication user to use other endpoints
// @tags        Auth
// @accept      json
// @param       login   body     types.LoginReq true " "
// @success     200     {object} types.PairToken
// @failure     400,500 {object} delivery.jsonResponse "error"
// @router      /auth/login [post]
func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq types.LoginReq
	err := a.readJSON(w, r, &loginReq)
	if err != nil {
		a.errorJSON(w, errors.New("can not parse values"), http.StatusBadRequest)
		return
	}
	// grpc reqest
	conn, err := grpc.DialContext(r.Context(), a.cfg.Auth.Host+":"+a.cfg.Auth.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		a.errorJSON(w, err)
		return
	}
	defer conn.Close()
	client := proto.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(a.cfg.Auth.ContextDeadline)*time.Second)
	defer cancel()
	result, err := client.Login(ctx, &proto.LoginRequest{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	})

	if err != nil {
		a.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// respons
	payload := jsonResponse{
		Error:   false,
		Message: "user logined succesfully",
		Data: types.LoginRes{
			Account: types.Account{
				Id:        result.User.Id,
				Username:  result.User.Usernae,
				Email:     result.User.Email,
				Password:  result.User.Password,
				CreatedAt: result.User.CreatedAt.AsTime(),
			},
			AccessToken:  result.PairToken.AccessToken,
			RefreshToken: result.PairToken.RefreshToken,
		},
	}
	a.writeJSON(w, http.StatusOK, payload)

}

// @summary     REGISTER USER
// @description register User for use api
// @tags        Auth
// @accept      json
// @param       register   body     types.RegisterReq true " "
// @success     200     {object} types.PairToken
// @failure     400,500 {object} delivery.jsonResponse "error"
// @router      /auth/register [post]
func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerReq types.RegisterReq
	err := a.readJSON(w, r, &registerReq)
	if err != nil {
		a.errorJSON(w, errors.New("can not parse values"), http.StatusBadRequest)
		return
	}
	// grpc reqest
	conn, err := grpc.DialContext(r.Context(), a.cfg.Auth.Host+":"+a.cfg.Auth.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		a.errorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	client := proto.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(a.cfg.Auth.ContextDeadline)*time.Second)
	defer cancel()

	result, err := client.Register(ctx, &proto.RegisterRequest{
		Email:    registerReq.Email,
		Username: registerReq.UserName,
		Password: registerReq.Password,
	})
	defer conn.Close()

	if err != nil {
		a.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// respons
	payload := jsonResponse{
		Error:   result.Error,
		Message: result.Message,
		Data:    nil,
	}
	a.writeJSON(w, http.StatusOK, payload)

}
