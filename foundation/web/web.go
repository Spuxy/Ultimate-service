package web

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dimfeld/httptreemux"
)

// A Handler is a type that handles a http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*httptreemux.ContextMux
	Shutdown chan os.Signal
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal) *App {
	return &App{
		ContextMux: httptreemux.NewContextMux(),
		Shutdown:   shutdown,
	}
}

// Handle sets a handler function for a given HTTP method and path pair
// to the application server mux.
func (a *App) Handle(method string, group string, path string, handler Handler) {

	h := func(w http.ResponseWriter, r *http.Request) {

		// PRE CODE PROCCESSING
		fmt.Fprintln(w, "PRE")

		if err := handler(r.Context(), w, r); err != nil {
			panic(err)
		}

		// POST CODE PROCCESSING
		fmt.Fprintln(w, "POST")

	}

	finalPath := path
	if group != "" {
		finalPath = "/" + group + path
	}

	a.ContextGroup.Handle(method, finalPath, h)
}
