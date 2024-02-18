package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/adapters/web/handlers"
	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/application"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer(pService application.ProductServiceInterface) *WebServer {
	return &WebServer{
		Service: pService,
	}
}

func (w *WebServer) Serve() {

	mx := chi.NewRouter()

	mx.Use(middleware.Logger)

	handlers.MakeProductHandlers(mx, w.Service)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           mx,
		ErrorLog:          log.New(os.Stderr, "log:", log.Lshortfile),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
