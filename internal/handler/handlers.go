package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "mygo/internal/model"
)

type Repository interface {
	GetBannersHandler() ([]models.Banner, error)
	CreateBannerHandler(banner models.Banner) error
	UpdateBannerHandler(banner models.Banner) error
	DeleteBannerHandler(id int) error
}

type Handler struct {
	repo Repository
}

func NewHander(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetBanners(w http.ResponseWriter, r *http.Request) {
	banners, err := h.repo.GetBannersHandler()
	if err != nil {
		http.Error(w, err.Error(), 404)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(banners)
}

func (h *Handler) CreateBanner(w http.ResponseWriter, r *http.Request) {
	var banner models.Banner
	err := json.NewDecoder(r.Body).Decode(&banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.repo.CreateBannerHandler(banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Banner created successfully"})
}

func (h *Handler) UpdateBanner(w http.ResponseWriter, r *http.Request) {
	var banner models.Banner
	err := json.NewDecoder(r.Body).Decode(&banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.repo.UpdateBannerHandler(banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Banner updated successfully"})
}

func (h *Handler) DeleteBanner(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	err = h.repo.DeleteBannerHandler(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Banner deleted successfully"})
}
