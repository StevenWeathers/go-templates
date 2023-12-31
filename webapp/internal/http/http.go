package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stevenweathers/go-templates/webapp/ui"
)

func (o *Endpoints) handle() {
	uiFS := ui.New(o.config.UseOSFilesystem)
	// k8s healthcheck /healthz as per convention
	o.router.HandleFunc("/healthz", o.handleHealthz).Methods(http.MethodGet)

	// index/static
	o.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(uiFS)))
}

func (o *Endpoints) ListenAndServe(ctx context.Context) error {
	slog.InfoContext(ctx, fmt.Sprintf("Listening on %s", o.config.ListenAddress))
	return http.ListenAndServe(o.config.ListenAddress, o.router)
}

func New(config Config) *Endpoints {
	router := mux.NewRouter()
	e := &Endpoints{
		config: config,
		router: router,
	}
	e.handle()
	return e
}
