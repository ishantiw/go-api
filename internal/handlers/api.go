package handlers

import (
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/ishantiw/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	// global middleware
	r.Use(chimiddleware.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		// Middleware for /account route
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
