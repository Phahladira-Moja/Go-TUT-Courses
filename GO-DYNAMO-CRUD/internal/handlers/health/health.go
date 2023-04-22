package health

import (
	"errors"
	"githhub.com/phahladira-moja/dynamodb-bulletproof-crud-api/internal/handlers"
	"githhub.com/phahladira-moja/dynamodb-bulletproof-crud-api/internal/repository/adapter"
	HttpStatus "githhub.com/phahladira-moja/dynamodb-bulletproof-crud-api/utils/http"
	"net/http"
)

type Handler struct {
	handlers.Interface
	Repository adapter.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Repository: repository,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	if !h.Repository.Health() {
		HttpStatus.StatusInternalServerError(w, r, errors.New("relational database not live"))
		return
	}

	HttpStatus.StatusOK(w, r, "Service OK")
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r, nil)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r, nil)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r, nil)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}
