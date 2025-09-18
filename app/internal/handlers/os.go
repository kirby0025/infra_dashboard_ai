package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"infra-dashboard/internal/database"
	"infra-dashboard/internal/models"

	"github.com/gorilla/mux"
)

// OSHandler handles operating system-related HTTP requests
type OSHandler struct {
	repo *database.OSRepository
}

// NewOSHandler creates a new OS handler
func NewOSHandler(repo *database.OSRepository) *OSHandler {
	return &OSHandler{repo: repo}
}

// GetOperatingSystems handles GET /os - retrieves all operating systems
func (h *OSHandler) GetOperatingSystems(w http.ResponseWriter, r *http.Request) {
	oss, err := h.repo.GetAll()
	if err != nil {
		log.Printf("Error getting operating systems: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(oss); err != nil {
		log.Printf("Error encoding operating systems response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// GetOperatingSystem handles GET /os/{id} - retrieves an operating system by ID
func (h *OSHandler) GetOperatingSystem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		http.Error(w, "Operating system ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid operating system ID", http.StatusBadRequest)
		return
	}

	os, err := h.repo.GetByID(id)
	if err != nil {
		log.Printf("Error getting operating system by ID %d: %v", id, err)
		http.Error(w, "Operating system not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(os); err != nil {
		log.Printf("Error encoding operating system response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// CreateOperatingSystem handles POST /os - creates a new operating system
func (h *OSHandler) CreateOperatingSystem(w http.ResponseWriter, r *http.Request) {
	var req models.CreateOSRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.Name == "" || req.Version == "" || req.EndOfSupport == "" {
		http.Error(w, "Name, version, and end of support date are required", http.StatusBadRequest)
		return
	}

	os, err := h.repo.Create(&req)
	if err != nil {
		log.Printf("Error creating operating system: %v", err)
		if err.Error() == "invalid end of support date format" {
			http.Error(w, "Invalid end of support date format. Use YYYY-MM-DD", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to create operating system", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(os); err != nil {
		log.Printf("Error encoding created operating system response: %v", err)
		return
	}
}

// UpdateOperatingSystem handles PUT /os/{id} - updates an existing operating system
func (h *OSHandler) UpdateOperatingSystem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		http.Error(w, "Operating system ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid operating system ID", http.StatusBadRequest)
		return
	}

	var req models.UpdateOSRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	os, err := h.repo.Update(id, &req)
	if err != nil {
		log.Printf("Error updating operating system with ID %d: %v", id, err)
		if err.Error() == "invalid end of support date format" {
			http.Error(w, "Invalid end of support date format. Use YYYY-MM-DD", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to update operating system", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(os); err != nil {
		log.Printf("Error encoding updated operating system response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// DeleteOperatingSystem handles DELETE /os/{id} - deletes an operating system
func (h *OSHandler) DeleteOperatingSystem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		http.Error(w, "Operating system ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid operating system ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		log.Printf("Error deleting operating system with ID %d: %v", id, err)
		if err.Error() == "cannot delete operating system: servers are using it" {
			http.Error(w, "Cannot delete operating system: servers are using it", http.StatusConflict)
		} else {
			http.Error(w, "Failed to delete operating system", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
