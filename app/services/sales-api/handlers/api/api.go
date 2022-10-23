package api

import (
	"context"
	"net/http"

	"github.com/Spuxy/service/foundation/web"
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

	return web.Respond(ctx, w, status, http.StatusOK)
}
