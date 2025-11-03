package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"infra-dashboard/internal/database"
	"infra-dashboard/internal/models"

	"github.com/gorilla/mux"
)

// ChangeHistoryHandler handles HTTP requests for server change history
type ChangeHistoryHandler struct {
	repo *database.ChangeHistoryRepository
}

// NewChangeHistoryHandler creates a new change history handler
func NewChangeHistoryHandler(repo *database.ChangeHistoryRepository) *ChangeHistoryHandler {
	return &ChangeHistoryHandler{repo: repo}
}

// GetChangeHistory retrieves all change history with optional filters
func (h *ChangeHistoryHandler) GetChangeHistory(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters for filters
	filter := &models.ChangeHistoryFilter{}

	// Parse server_id filter
	if serverIDStr := r.URL.Query().Get("server_id"); serverIDStr != "" {
		serverID, err := strconv.Atoi(serverIDStr)
		if err != nil {
			http.Error(w, "Invalid server_id parameter", http.StatusBadRequest)
			return
		}
		filter.ServerID = &serverID
	}

	// Parse change_type filter
	if changeType := r.URL.Query().Get("change_type"); changeType != "" {
		// Validate change_type
		if changeType != "created" && changeType != "os_changed" && changeType != "deleted" {
			http.Error(w, "Invalid change_type. Must be: created, os_changed, or deleted", http.StatusBadRequest)
			return
		}
		filter.ChangeType = &changeType
	}

	// Parse start_date filter
	if startDateStr := r.URL.Query().Get("start_date"); startDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			http.Error(w, "Invalid start_date format. Use YYYY-MM-DD", http.StatusBadRequest)
			return
		}
		filter.StartDate = &startDate
	}

	// Parse end_date filter
	if endDateStr := r.URL.Query().Get("end_date"); endDateStr != "" {
		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			http.Error(w, "Invalid end_date format. Use YYYY-MM-DD", http.StatusBadRequest)
			return
		}
		// Set to end of day
		endDate = endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
		filter.EndDate = &endDate
	}

	// Parse limit (default 100)
	limit := 100
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit < 1 {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
		limit = parsedLimit
	}
	filter.Limit = limit

	// Parse offset (default 0)
	offset := 0
	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		parsedOffset, err := strconv.Atoi(offsetStr)
		if err != nil || parsedOffset < 0 {
			http.Error(w, "Invalid offset parameter", http.StatusBadRequest)
			return
		}
		offset = parsedOffset
	}
	filter.Offset = offset

	// Get change history from repository
	history, err := h.repo.GetAll(filter)
	if err != nil {
		http.Error(w, "Failed to retrieve change history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

// GetServerChangeHistory retrieves change history for a specific server
func (h *ChangeHistoryHandler) GetServerChangeHistory(w http.ResponseWriter, r *http.Request) {
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

	// Parse limit (default 50)
	limit := 50
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit < 1 {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
		limit = parsedLimit
	}

	history, err := h.repo.GetByServerID(id, limit)
	if err != nil {
		http.Error(w, "Failed to retrieve change history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

// GetChangeHistoryByID retrieves a single change history record by its ID
func (h *ChangeHistoryHandler) GetChangeHistoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		http.Error(w, "Change history ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid change history ID", http.StatusBadRequest)
		return
	}

	record, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, "Failed to retrieve change history record", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}
