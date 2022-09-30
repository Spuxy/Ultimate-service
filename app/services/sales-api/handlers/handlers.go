package handlers

import (
	"expvar"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/Spuxy/service/app/services/sales-api/handlers/debug"
	v1 "github.com/Spuxy/service/app/services/sales-api/handlers/v1"
	"github.com/Spuxy/service/foundation/web"

	"go.uber.org/zap"
)

// DebugStandardLibraryMux registers all the debug routes from the standard library
// into a new mux bypassing the use of the DefaultServerMux. Using the
// DefaultServerMux would be a security risk since a dependency could inject a
// handler into our service without us knowing it.
func debugStandardLibraryMux() *http.ServeMux {
	mux := http.NewServeMux()

	// Register all the standard library debug endpoints.
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())

	return mux
}

func DebugMux(build string, l *zap.SugaredLogger) http.Handler {
	mux := debugStandardLibraryMux()

	dmux := debug.Handlers{
		Build: build,
		Log:   l,
	}

	mux.HandleFunc("/debug/readiness", dmux.Readiness)
	mux.HandleFunc("/debug/liveness", dmux.Liveness)

	return mux
}

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

func APIMux(cfg APIMuxConfig) *web.App {
	mx := web.NewApp(cfg.Shutdown)

	ah := v1.Handler{
		Log: cfg.Log,
	}

	mx.Handle(http.MethodGet, "/test", ah.Test)

	return mx
}
