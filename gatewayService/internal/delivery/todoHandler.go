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

	"github.com/go-chi/chi/v5"
	"github.com/mohammaderm/todoMicroService/todoService/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	TodoHandler struct {
		HandlerHelper
		cfg *config.Service
	}
	TodoHandlerContrac interface {
		Create(w http.ResponseWriter, r *http.Request)
		Delete(w http.ResponseWriter, r *http.Request)
		GetAll(w http.ResponseWriter, r *http.Request)
		UpdateStatus(w http.ResponseWriter, r *http.Request)
		UpdatePriority(w http.ResponseWriter, r *http.Request)
		UpdateDueDate(w http.ResponseWriter, r *http.Request)
	}
)

func NewTodoHandler(logger logger.Logger, cfg *config.Service) TodoHandlerContrac {
	return &TodoHandler{
		HandlerHelper: HandlerHelper{
			logger: logger,
		},
		cfg: cfg,
	}
}

// @summary     UPDATE TODO DUE_DATE
// @description update todo due_date based on todo Id and due_date. (auth required)
// @tags        Todo
// @accept      json
// @Security apiKey
// @param       dueDate   body     types.UpdateDueDate true " "
// @param       id   path     int true " "
// @success     200     {object} delivery.jsonResponse
// @failure     400,500 {object} delivery.jsonResponse "error"
// @router      /todo/due_date/{id} [put]
func (t *TodoHandler) UpdateDueDate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		t.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		t.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}
	var priority types.UpdateDueDate
	t.readJSON(w, r, &priority)
	dueDate, _ := time.Parse("2006-01-02 15:04:05", priority.DueDate)

	accountInfo := r.Context().Value(accountInfoKeyCtx).(types.AccountInfo)
	// grpc request
	conn, err := grpc.DialContext(r.Context(), t.cfg.Todo.Host+":"+t.cfg.Todo.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.errorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	client := proto.NewTodoServiceClient(conn)
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(t.cfg.Todo.ContextDeadline)*time.Second)
	defer cancel()
	respons, err := client.UpdateDueDate(ctx, &proto.UpdateDueDateRequest{
		AccountId: accountInfo.Id,
		Id:        uint64(idInt),
		DueDate:   timestamppb.New(dueDate),
	})
	if err != nil {
		t.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	payload := jsonResponse{
		Error:   respons.Error,
		Message: respons.Message,
		Data:    nil,
	}
	t.writeJSON(w, http.StatusOK, payload)
}

// @summary     UPDATE TODO PRIORITY
// @description update todo priority based on todo Id and priority. (auth required)
// @tags        Todo
// @accept      json
// @Security apiKey
// @param       priority   body     types.UpdatePriority true " "
// @param       id   path     int true " "
// @success     200     {object} delivery.jsonResponse
// @failure     400,500 {object} delivery.jsonResponse "error"
// @router      /todo/priority/{id} [put]
func (t *TodoHandler) UpdatePriority(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		t.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		t.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}
	var priority types.UpdatePriority
	t.readJSON(w, r, &priority)

	accountInfo := r.Context().Value(accountInfoKeyCtx).(types.AccountInfo)
	// grpc request
	conn, err := grpc.DialContext(r.Context(), t.cfg.Todo.Host+":"+t.cfg.Todo.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.errorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	client := proto.NewTodoServiceClient(conn)
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(t.cfg.Todo.ContextDeadline)*time.Second)
	defer cancel()
	respons, err := client.UpdatePriority(ctx, &proto.UpdatePriorityRequest{
		AccountId: accountInfo.Id,
		Id:        uint64(idInt),
		Priority:  uint64(priority.Priority),
	})
	if err != nil {
		t.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	payload := jsonResponse{
		Error:   respons.Error,
		Message: respons.Message,
		Data:    nil,
	}
	t.writeJSON(w, http.StatusOK, payload)
}

// @summary     UPDATE TODO STATUS
// @description update todo Status based on todo Id. (auth required)
// @tags        Todo
// @accept      json
// @Security apiKey
// @param       id   path     int true " "
// @success     200     {object} delivery.jsonResponse
// @failure     400,500 {object} delivery.jsonResponse "error"
// @router      /todo/status/{id} [put]
func (t *TodoHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		t.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		t.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}

	accountInfo := r.Context().Value(accountInfoKeyCtx).(types.AccountInfo)
	// grpc request
	conn, err := grpc.DialContext(r.Context(), t.cfg.Todo.Host+":"+t.cfg.Todo.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.errorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	client := proto.NewTodoServiceClient(conn)
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(t.cfg.Todo.ContextDeadline)*time.Second)
	defer cancel()
	respons, err := client.UpdateStatus(ctx, &proto.UpdateStatusRequest{
		AccountId: accountInfo.Id,
		Id:        uint64(idInt),
	})
	if err != nil {
		t.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	payload := jsonResponse{
		Error:   respons.Error,
		Message: respons.Message,
		Data:    nil,
	}
	t.writeJSON(w, http.StatusOK, payload)
}

// @summary     DELETE TODO
// @description delete todo based on todo Id. (auth required)
// @tags        Todo
// @accept      json
// @Security apiKey
// @param       id   path     int string " "
// @success     200     {object} delivery.jsonResponse
// @failure     400,500 {object} delivery.jsonResponse "error"
// @router      /todo/{id} [delete]

func (t *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		t.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		t.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}

	accountInfo := r.Context().Value(accountInfoKeyCtx).(types.AccountInfo)
	// grpc request
	conn, err := grpc.DialContext(r.Context(), t.cfg.Todo.Host+":"+t.cfg.Todo.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.errorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	client := proto.NewTodoServiceClient(conn)
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(t.cfg.Todo.ContextDeadline)*time.Second)
	defer cancel()
	respons, err := client.Delete(ctx, &proto.DeleteRequest{
		AccountId: accountInfo.Id,
		Id:        uint64(idInt),
	})
	if err != nil {
		t.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	payload := jsonResponse{
		Error:   respons.Error,
		Message: respons.Message,
		Data:    nil,
	}
	t.writeJSON(w, http.StatusOK, payload)
}

// @summary     GETALL TODO
// @description get all todo based on offset for pagination. (auth required)
// @tags        Todo
// @accept      json
// @Security apiKey
// @param       offset   query     int string "minimum number for offset is '0', defualt limit is '5'"
// @success     200     {object} delivery.jsonResponse
// @failure     400,500 {object} delivery.jsonResponse "error"
// @router      /todo/getall [get]
func (t *TodoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	offset := r.URL.Query().Get("offset")
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		t.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}
	accountInfo := r.Context().Value(accountInfoKeyCtx).(types.AccountInfo)
	// grpc request
	conn, err := grpc.DialContext(r.Context(), t.cfg.Todo.Host+":"+t.cfg.Todo.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.errorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	client := proto.NewTodoServiceClient(conn)
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(t.cfg.Todo.ContextDeadline)*time.Second)
	defer cancel()
	respons, err := client.GetAll(ctx, &proto.GetAllRequest{
		AccountId: accountInfo.Id,
		Offset:    uint64(offsetInt),
	})
	if err != nil {
		t.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	payload := jsonResponse{
		Error:   respons.Error,
		Message: respons.Message,
		Data:    respons.Todos,
	}
	t.writeJSON(w, http.StatusOK, payload)
}

// @summary     CREATE TODO
// @description create todo based on params. (auth required)
// @tags        Todo
// @accept      json
// @Security apiKey
// @param       todo   body     types.CreateTodoReq true " "
// @success     200     {object} delivery.jsonResponse
// @failure     400,500 {object} delivery.jsonResponse "error"
// @router      /todo/create [post]
func (t *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req types.CreateTodoReq
	err := t.readJSON(w, r, &req)
	if err != nil {
		t.errorJSON(w, errors.New("can not parse values"), http.StatusBadRequest)
		return
	}
	accountInfo := r.Context().Value(accountInfoKeyCtx).(types.AccountInfo)
	// grpc reqest
	conn, err := grpc.DialContext(r.Context(), t.cfg.Todo.Host+":"+t.cfg.Todo.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.errorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	client := proto.NewTodoServiceClient(conn)
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(t.cfg.Todo.ContextDeadline)*time.Second)
	defer cancel()

	respons, err := client.Create(ctx, &proto.CreateRequest{
		AccountId:   accountInfo.Id,
		CategoryId:  req.CategoryId,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		t.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	payload := jsonResponse{
		Error:   respons.Error,
		Message: respons.Message,
		Data:    nil,
	}
	t.writeJSON(w, http.StatusOK, payload)

}
