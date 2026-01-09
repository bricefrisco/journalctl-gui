package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/bricefrisco/journalctl-gui/internal/journal"
)

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

	return mux
}
