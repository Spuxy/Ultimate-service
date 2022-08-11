package debug

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type Handlers struct {
	Log *zap.SugaredLogger
}

func (h Handlers) Readiness(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Status string
	}{
		Status: "OK",
	}

	statusCode := http.StatusOK

	if err := response(w, data, statusCode); err != nil {
		h.Log.Errorw("readiness", "ERROR", err)
	}

	h.Log.Infow("readiness", "statusCode", statusCode, "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)
}

func response(w http.ResponseWriter, data interface{}, code int) error {
	msg, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if _, err = w.Write(msg); err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return nil
}
