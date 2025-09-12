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

// ServerHandler handles server-related HTTP requests
type ServerHandler struct {
	repo *database.ServerRepository
}

// NewServerHandler creates a new server handler
func NewServerHandler(repo *database.ServerRepository) *ServerHandler {
	return &ServerHandler{repo: repo}
}

// GetServers handles GET /servers - retrieves all servers
func (h *ServerHandler) GetServers(w http.ResponseWriter, r *http.Request) {
	servers, err := h.repo.GetAll()
	if err != nil {
		log.Printf("Error getting servers: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(servers); err != nil {
		log.Printf("Error encoding servers response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// GetServer handles GET /servers/{id} - retrieves a server by ID
func (h *ServerHandler) GetServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		http.Error(w, "Server ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid server ID", http.StatusBadRequest)
		return
	}

	server, err := h.repo.GetByID(id)
	if err != nil {
		log.Printf("Error getting server by ID %d: %v", id, err)
		http.Error(w, "Server not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(server); err != nil {
		log.Printf("Error encoding server response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// CreateServer handles POST /servers - creates a new server
func (h *ServerHandler) CreateServer(w http.ResponseWriter, r *http.Request) {
	var req models.CreateServerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.Name == "" || req.OS == "" || req.OSVersion == "" {
		http.Error(w, "Name, OS, and OS version are required", http.StatusBadRequest)
		return
	}

	server, err := h.repo.Create(&req)
	if err != nil {
		log.Printf("Error creating server: %v", err)
		http.Error(w, "Failed to create server", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(server); err != nil {
		log.Printf("Error encoding created server response: %v", err)
		return
	}
}

// UpdateServer handles PUT /servers/{id} - updates an existing server
func (h *ServerHandler) UpdateServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		http.Error(w, "Server ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid server ID", http.StatusBadRequest)
		return
	}

	var req models.UpdateServerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	server, err := h.repo.Update(id, &req)
	if err != nil {
		log.Printf("Error updating server with ID %d: %v", id, err)
		http.Error(w, "Failed to update server", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(server); err != nil {
		log.Printf("Error encoding updated server response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// DeleteServer handles DELETE /servers/{id} - deletes a server
func (h *ServerHandler) DeleteServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		http.Error(w, "Server ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid server ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		log.Printf("Error deleting server with ID %d: %v", id, err)
		http.Error(w, "Failed to delete server", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// HealthCheck handles GET /health - simple health check endpoint
func (h *ServerHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "healthy",
		"service": "infra-dashboard",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
