package api

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type Handler struct {
	Log *zap.SugaredLogger
}

func (h Handler) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	statusCode := http.StatusOK
	h.Log.Infow("readiness", "statusCode", statusCode, "method", r.Method, "path", r.URL.Path)

	return json.NewEncoder(w).Encode(status)
}
