package config

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/resource"
)

func GetResources(svc resource.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resources, err := svc.GetResources()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Message: "Unable to get resources"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resources)
	}
}

func GetResource(svc resource.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		res, err := svc.GetResourceByID(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.APIError{Message: "Resource not found"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func CreateResource(svc resource.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res models.Resource
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Message: "Invalid request body"})
			return
		}
		if err := json.Unmarshal(body, &res); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Message: "Invalid JSON format"})
			return
		}

		if res.FilePath == "" {
			res.FilePath = helpers.DefaultICSURL
		}

		if err := svc.CreateResource(&res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Message: "Unable to create resource"})
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}

func UpdateResource(svc resource.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var res models.Resource
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Message: "Invalid request body"})
			return
		}
		if err := json.Unmarshal(body, &res); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Message: "Invalid JSON format"})
			return
		}
		res.ID = id
		if err := svc.UpdateResource(&res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Message: "Unable to update resource"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func DeleteResource(svc resource.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if err := svc.DeleteResource(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Message: "Unable to delete resource"})
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
