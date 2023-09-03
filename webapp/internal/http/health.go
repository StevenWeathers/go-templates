package http

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

// handleHealthz returns a json response for healthcheck with a key of `alive` set to true
func (o *Endpoints) handleHealthz(w http.ResponseWriter, r *http.Request) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	err := sendJsonResponse(w, `{"alive": true}`)
	if err != nil {
		logger.ErrorContext(r.Context(), fmt.Sprintf("error sending json response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
