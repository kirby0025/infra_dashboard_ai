package main

import (
	"log"
	"net/http"
	"time"

	"infra-dashboard/internal/config"
	"infra-dashboard/internal/database"
	"infra-dashboard/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := database.New(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create tables if they don't exist
	if err := db.CreateTablesIfNotExist(); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	// Initialize repository
	serverRepo := database.NewServerRepository(db)

	// Initialize handlers
	serverHandler := handlers.NewServerHandler(serverRepo)

	// Setup router
	router := mux.NewRouter()

	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()

	// Server routes
	api.HandleFunc("/servers", serverHandler.GetServers).Methods("GET")
	api.HandleFunc("/servers", serverHandler.CreateServer).Methods("POST")
	api.HandleFunc("/servers/{id:[0-9]+}", serverHandler.GetServer).Methods("GET")
	api.HandleFunc("/servers/{id:[0-9]+}", serverHandler.UpdateServer).Methods("PUT")
	api.HandleFunc("/servers/{id:[0-9]+}", serverHandler.DeleteServer).Methods("DELETE")

	// Health check
	router.HandleFunc("/health", serverHandler.HealthCheck).Methods("GET")

	// Add CORS middleware
	router.Use(corsMiddleware)

	// Add logging middleware
	router.Use(loggingMiddleware)

	log.Printf("Starting server on port %s", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, router))
}

// corsMiddleware adds CORS headers to responses
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// loggingMiddleware logs HTTP requests in Apache-style format
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()

		// Wrap the ResponseWriter to capture status code
		wrapped := &responseWriter{ResponseWriter: w, statusCode: 200}

		next.ServeHTTP(wrapped, r)

		s := wrapped.statusCode

		log.Printf("%s - - %s \"%s %s %s\" %d %d \"-\" \"%s\" %d\n",
			r.Host,
			t.Format("[02/Jan/2006:15:04:05 -0700]"),
			r.Method,
			r.URL.Path,
			r.Proto,
			s,
			r.ContentLength,
			r.UserAgent(),
			time.Since(t).Milliseconds(),
		)
	})
}
