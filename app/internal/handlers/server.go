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
	repo   *database.ServerRepository
	osRepo *database.OSRepository
}

// NewServerHandler creates a new server handler
func NewServerHandler(repo *database.ServerRepository, osRepo *database.OSRepository) *ServerHandler {
	return &ServerHandler{repo: repo, osRepo: osRepo}
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
	if req.Name == "" || req.OSID == 0 {
		http.Error(w, "Name and OS ID are required", http.StatusBadRequest)
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

// GetComplianceReport handles GET /servers/compliance - generates compliance report
func (h *ServerHandler) GetComplianceReport(w http.ResponseWriter, r *http.Request) {
	// Get all servers with OS information
	servers, err := h.repo.GetAll()
	if err != nil {
		log.Printf("Error getting servers for compliance report: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Generate compliance report using utility functions
	complianceUtils := models.NewComplianceUtils()
	report := complianceUtils.GenerateComplianceReport(servers)

	// Get all OS data for recommendations
	allOS, err := h.osRepo.GetAll()
	if err != nil {
		log.Printf("Error getting OS data for recommendations: %v", err)
		// Continue without recommendations rather than failing
		allOS = []models.OS{}
	}

	// Add recommendations to the report
	recommendations := complianceUtils.GetRecommendations(servers, allOS)

	// Create extended report with compliance score and recommendations
	extendedReport := struct {
		models.ComplianceReport
		ComplianceScore  float64  `json:"compliance_score"`
		Recommendations  []string `json:"recommendations"`
		ScoreDescription string   `json:"score_description"`
	}{
		ComplianceReport: report,
		ComplianceScore:  complianceUtils.GetComplianceScore(servers),
		Recommendations:  recommendations,
		ScoreDescription: getScoreDescription(complianceUtils.GetComplianceScore(servers)),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(extendedReport); err != nil {
		log.Printf("Error encoding compliance report response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// getScoreDescription provides human-readable description for compliance scores
func getScoreDescription(score float64) string {
	switch {
	case score >= 90:
		return "Excellent - Infrastructure is well maintained and compliant"
	case score >= 75:
		return "Good - Minor compliance issues that should be addressed"
	case score >= 50:
		return "Fair - Several compliance issues requiring attention"
	case score >= 25:
		return "Poor - Significant compliance issues need immediate action"
	default:
		return "Critical - Infrastructure has serious compliance problems"
	}
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
