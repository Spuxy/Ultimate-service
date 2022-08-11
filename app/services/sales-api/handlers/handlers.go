package handlers

import (
	"encoding/json"
	"expvar"
	"github.com/Spuxy/service/app/services/sales-api/handlers/debug"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/dimfeld/httptreemux"
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

func DebugMux(l *zap.SugaredLogger) http.Handler {
	mux := debugStandardLibraryMux()

	dmux := debug.Handlers{
		Log: l,
	}

	mux.HandleFunc("/debug/readiness", dmux.Readiness)

	return mux
}

type APIMuxConfig struct {
	Log      *zap.SugaredLogger
	Shutdown chan os.Signal
}

func APIMux(cfg APIMuxConfig) *httptreemux.ContextMux {
	mx := httptreemux.NewContextMux()

	mx.Handle(http.MethodGet, "/test", func(rw http.ResponseWriter, r *http.Request) {
		status := struct {
			Status string
		}{
			Status: "OK",
		}
		json.NewEncoder(rw).Encode(status)
	})

	return mx
}
