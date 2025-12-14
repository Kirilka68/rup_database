package http

import (
	"encoding/json"
	"net/http"
	"rup_database/internal/models"
	"rup_database/internal/repository"
)

type Handler struct {
	repo *repository.ObjectRepo
}

func NewHandler(repo *repository.ObjectRepo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) CreateObject(w http.ResponseWriter, r *http.Request) {
	var dto models.CreateObjectDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	obj, err := h.repo.Create(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(obj)
}

func (h *Handler) GetObject(w http.ResponseWriter, r *http.Request) {
	var dto models.GetObjectDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	obj, err := h.repo.GetByID(r.Context(), dto.ID)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(obj)
}

func (h *Handler) UpdateObject(w http.ResponseWriter, r *http.Request) {
	var dto models.UpdateObjectDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	obj, err := h.repo.Update(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(obj)
}

func (h *Handler) DeleteObject(w http.ResponseWriter, r *http.Request) {
	var dto models.DeleteObjectDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err := h.repo.Delete(r.Context(), dto.ID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(204)
}
