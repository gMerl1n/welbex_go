package server

import (
	"net/http"

	"github.com/gMerl1n/welbex_go/internal/configs"
	"github.com/gMerl1n/welbex_go/internal/handlers"
	"github.com/go-chi/chi"
)

func NewServer(cfg *configs.Config, handlers *handlers.Handlers) *http.Server {

	router := chi.NewRouter()

	router.Get("/", handlers.TestOk)

	return &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}
}
