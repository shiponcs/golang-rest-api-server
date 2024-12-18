package api

import (
	"errors"
	"github.com/bookstore-rest-api-server/auth"
	"github.com/bookstore-rest-api-server/db"
	endpointshandler "github.com/bookstore-rest-api-server/endpoint_handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func ServeEndpoints(port string) {
	db.Init()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(auth.Authenticate)

	router.Post("/authorize", endpointshandler.Login)
	router.Get("/users", endpointshandler.GetUsers)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("listen: %s\n", err)
	}
}
