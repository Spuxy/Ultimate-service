package handlers

import (
	"expvar"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/Spuxy/service/app/services/sales-api/handlers/api"
	"github.com/Spuxy/service/app/services/sales-api/handlers/debug"
	"github.com/Spuxy/service/business/web/mid"
	"github.com/Spuxy/service/foundation/web"
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

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *httptreemux.ContextMux {

	// Construct the web.App which holds all routes as well as common Middleware.
	mx := web.NewApp(
		cfg.Shutdown,
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Panics(),
	)

	// Load the v1 routes.
	v1(mx, cfg)

	return mx.ContextMux
}

// Routes binds all the version 1 routes.
func v1(app *web.App, cfg APIMuxConfig) {
	const version = "v1"

	ah := api.Handler{
		Log: cfg.Log,
	}

	app.Handle(http.MethodGet, version, "/test", ah.Test)
}
