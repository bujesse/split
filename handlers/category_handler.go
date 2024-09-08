package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"split/config/logger"
	"split/models"
	"split/repositories"
	"split/views/components"
	"split/views/partials"
	"strconv"
)

type CategoryHandler struct {
	repo repositories.CategoryRepository
}

func NewCategoryHandler(repo repositories.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{repo}
}

func (h *CategoryHandler) CreateCategory(response http.ResponseWriter, request *http.Request) {
	logger.Debug.Println("Creating category")

	category := models.Category{
		Name:        request.FormValue("name"),
		Description: request.FormValue("description"),
		Type:        request.FormValue("type"),
	}

	if err := h.repo.Create(&category); err != nil {
		http.Error(response, "Failed to save category", http.StatusInternalServerError)
		return
	}

	logger.Debug.Println("Created Category with ID: ", category.ID)

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("HX-Trigger", "reloadCategories")
	json.NewEncoder(response).Encode(category)
	response.WriteHeader(http.StatusCreated)
}

func (h *CategoryHandler) GetAllCategories(response http.ResponseWriter, request *http.Request) {
	categories, err := h.repo.GetAll()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "text/html")
	partials.CategoriesTable(categories).Render(context.Background(), response)
}

func (h *CategoryHandler) DeleteCategory(response http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.repo.Delete(uint(id)); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusNoContent)
}

func (h *CategoryHandler) EditCategoryByID(w http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	category, err := h.repo.GetByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	components.CategoriesForm(category).Render(context.Background(), w)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	categoryType := r.FormValue("type")

	category, err := h.repo.GetByID(uint(id))

	category.Name = name
	category.Description = description
	category.Type = categoryType

	if err := h.repo.Update(category); err != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "reloadCategories")
	w.WriteHeader(http.StatusOK)
}
