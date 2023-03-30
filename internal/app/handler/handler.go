package handler

import (
	"github.com/gorilla/mux"
	"technodom/internal/app/config"
	"technodom/internal/repository"
	"technodom/internal/service"
	"technodom/internal/util/logger"
)

// Handler for http requests
type Handler struct {
	mux *mux.Router

	logger  logger.Logger
	config  *config.TomlConfig
	service *service.Service
	cache   repository.Cache
}

// New http handler
func New(
	mux *mux.Router,
	logger logger.Logger,
	config *config.TomlConfig,
	service *service.Service,
	cache repository.Cache,
) *Handler {
	handler := Handler{mux, logger, config, service, cache}
	handler.registerRoutes()

	return &handler
}

// Register all routes
func (h *Handler) registerRoutes() {
	h.mux.HandleFunc("/admin/redirects/", h.getList).Methods("GET")
	h.mux.HandleFunc("/admin/redirects/{id:.*\\/.*}", h.getByID).Methods("GET")
	h.mux.HandleFunc("/admin/redirects", h.post).Methods("POST")
	h.mux.HandleFunc("/admin/redirects/{id:.*\\/.*}", h.patch).Methods("PATCH")
	h.mux.HandleFunc("/admin/redirects/{id:.*\\/.*}", h.delete).Methods("DELETE")
	h.mux.HandleFunc("/redirects/{id:.*\\/.*}", h.get).Methods("GET")
}
