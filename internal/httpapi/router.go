package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/bricefrisco/journalctl-gui/internal/journal"
	"github.com/bricefrisco/journalctl-gui/internal/util"
)

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/services", func(w http.ResponseWriter, _ *http.Request) {
		services, err := journal.ListServices()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(services)
	})

	mux.HandleFunc("/api/logs", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		limit := util.Atoi(q.Get("limit"))
		if limit <= 0 || limit > 500 {
			limit = 100
		}

		cursor := q.Get("cursor")

		page, err := journal.ListLogsPage(limit, cursor)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(page)
	})

	return withCORS(mux)
}
