package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/adapters/dto"
	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/application"
	"github.com/go-chi/chi/v5"
)

func MakeProductHandlers(mx *chi.Mux, service application.ProductServiceInterface) {

	mx.Get("/products/{id}", getProduct(service))
	mx.Post("/products", createProduct(service))
	mx.Put("/products/{id}/enable", enableProduct(service))
	mx.Put("/products/{id}/disable", disableProduct(service))
}

func getProduct(service application.ProductServiceInterface) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		p, err := service.Get(id)

		if err != nil {
			err := jsonError(err.Error())
			w.WriteHeader(http.StatusNotFound)
			w.Write(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(p)

		if err != nil {
			log.Println("Error encoding product", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func createProduct(service application.ProductServiceInterface) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var p dto.ProductDTO

		err := json.NewDecoder(r.Body).Decode(&p)

		if err != nil {
			log.Println("Error decoding product", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		product, err := p.Bind(&application.Product{})

		if err != nil {
			log.Println("Error binding product", err.Error())
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		pc, err := service.Create(product.GetName(), product.GetPrice())

		if err != nil {
			log.Println("Error creating product", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(pc)
	}
}

func enableProduct(service application.ProductServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := chi.URLParam(r, "id")

		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		p, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			err := jsonError(err.Error())
			w.Write(err)
			return
		}

		err = p.Enable()

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err := jsonError(err.Error())
			w.Write(err)
			return
		}

		json.NewEncoder(w).Encode(p)
	}
}

func disableProduct(service application.ProductServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := chi.URLParam(r, "id")

		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		p, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			err := jsonError(err.Error())
			w.Write(err)
			return
		}

		err = p.Disable()

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err := jsonError(err.Error())
			w.Write(err)
			return
		}

		json.NewEncoder(w).Encode(p)
	}
}
