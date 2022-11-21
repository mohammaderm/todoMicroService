package delivery

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/mohammaderm/todoMicroService/gatewayService/config"
	"github.com/mohammaderm/todoMicroService/gatewayService/internal/types"
	"github.com/mohammaderm/todoMicroService/gatewayService/pkg/logger"
	"github.com/mohammaderm/todoMicroService/todoService/proto"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	CategoryHandler struct {
		HandlerHelper
		cfg *config.Service
	}
	CategoryHandlerContrac interface {
		CreateCat(w http.ResponseWriter, r *http.Request)
		DeleteCat(w http.ResponseWriter, r *http.Request)
		GetAllCat(w http.ResponseWriter, r *http.Request)
	}
)

func NewCategoryHandler(logger logger.Logger, cfg *config.Service) CategoryHandlerContrac {
	return &CategoryHandler{
		HandlerHelper: HandlerHelper{
			logger: logger,
		},
		cfg: cfg,
	}
}

// @summary     DELETE CATEGORY
// @description delete category based on category Id. (auth required)
// @tags        Category
// @accept      json
// @Security apiKey
// @param       id   path     int true " "
// @success     200     {object} delivery.jsonResponse
// @failure     400,500 {object} delivery.jsonResponse "error"
// @router      /category/{id} [delete]
func (c *CategoryHandler) DeleteCat(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		c.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}

	accountInfo := r.Context().Value(accountInfoKeyCtx).(types.AccountInfo)
	// grpc request
	conn, err := grpc.DialContext(r.Context(), c.cfg.Todo.Host+":"+c.cfg.Todo.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.errorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	client := proto.NewCategoryServiceClient(conn)
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(c.cfg.Todo.ContextDeadline)*time.Second)
	defer cancel()
	respons, err := client.DeleteCat(ctx, &proto.DeleteCatRequest{
		AccountId: accountInfo.Id,
		Id:        uint64(idInt),
	})
	if err != nil {
		c.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	payload := jsonResponse{
		Error:   respons.Error,
		Message: respons.Message,
		Data:    nil,
	}
	c.writeJSON(w, http.StatusOK, payload)
}

// @summary     GETALL CATEGORY
// @description get all category. (auth required)
// @tags        Category
// @accept      json
// @Security apiKey
// @success     200     {object} delivery.jsonResponse
// @failure     400,500 {object} delivery.jsonResponse "error"
// @router      /category/getall [get]
func (c *CategoryHandler) GetAllCat(w http.ResponseWriter, r *http.Request) {
	accountInfo := r.Context().Value(accountInfoKeyCtx).(types.AccountInfo)
	// grpc request
	conn, err := grpc.DialContext(r.Context(), c.cfg.Todo.Host+":"+c.cfg.Todo.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.errorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	client := proto.NewCategoryServiceClient(conn)
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(c.cfg.Todo.ContextDeadline)*time.Second)
	defer cancel()
	respons, err := client.GetAllCat(ctx, &proto.GetAllCatRequest{
		AccountId: accountInfo.Id,
	})
	if err != nil {
		c.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	payload := jsonResponse{
		Error:   respons.Error,
		Message: respons.Message,
		Data:    respons.Categorys,
	}
	c.writeJSON(w, http.StatusOK, payload)
}

// @summary     CREATE CATEGORY
// @description create category based on params. (auth required)
// @tags        Category
// @accept      json
// @Security apiKey
// @param       category   body  types.CreateCategoryReq  true " "
// @success     200     {object} delivery.jsonResponse
// @failure     400,500 {object} delivery.jsonResponse "error"
// @router      /category/create [post]
func (c *CategoryHandler) CreateCat(w http.ResponseWriter, r *http.Request) {
	var req types.CreateCategoryReq
	err := c.readJSON(w, r, &req)
	if err != nil {
		c.errorJSON(w, errors.New("can not parse values"), http.StatusBadRequest)
		return
	}
	accountInfo := r.Context().Value(accountInfoKeyCtx).(types.AccountInfo)
	// grpc reqest
	conn, err := grpc.DialContext(r.Context(), c.cfg.Todo.Host+":"+c.cfg.Todo.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.errorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	client := proto.NewCategoryServiceClient(conn)
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(c.cfg.Todo.ContextDeadline)*time.Second)
	defer cancel()

	respons, err := client.CreateCat(ctx, &proto.CreateCatRequest{
		Title:     req.Title,
		AccountId: accountInfo.Id,
	})
	if err != nil {
		c.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	payload := jsonResponse{
		Error:   respons.Error,
		Message: respons.Message,
		Data:    respons.Category,
	}
	c.writeJSON(w, http.StatusOK, payload)

}
