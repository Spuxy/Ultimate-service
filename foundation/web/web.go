package web

import (
	"os"

	"github.com/dimfeld/httptreemux"
)

type App struct {
	*httptreemux.ContextMux
	Shutdown chan os.Signal
}

func NewApp(sd chan os.Signal) *App {
	return &App{
		ContextMux: httptreemux.NewContextMux(),
		Shutdown:   sd,
	}
}
