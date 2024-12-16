package api

import (
	"errors"
	"github.com/bookstore-rest-api-server/auth"
	endpointshandler "github.com/bookstore-rest-api-server/endpoints-handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func ServeEndpoints(port string) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(auth.Authenticate)

	router.Post("/authorize", endpointshandler.Login)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("listen: %s\n", err)
	}
}
