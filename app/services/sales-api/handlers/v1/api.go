package v1

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type Handler struct {
	Log *zap.SugaredLogger
}

func (Handler) Test(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	json.NewEncoder(w).Encode(status)
}
