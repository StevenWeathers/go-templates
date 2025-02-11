package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/stevenweathers/go-templates/webapp/ui"
)

func (o *Endpoints) handle() {
	uiFS := ui.New(o.config.UseOSFilesystem)
	// k8s healthcheck /healthz as per convention
	o.router.HandleFunc("/healthz", o.handleHealthz).Methods(http.MethodGet)

	// index/static
	o.router.PathPrefix("/").Handler(http.StripPrefix("/", spaHandler(uiFS)))
}

func spaHandler(uiFS http.FileSystem) http.Handler {
	fileServer := http.FileServer(uiFS)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the absolute path to check if file exists
		path := r.URL.Path

		// Try to open the file from the filesystem
		f, err := uiFS.Open(strings.TrimPrefix(path, "/"))
		if err != nil {
			// If file doesn't exist, serve index.html
			if os.IsNotExist(err) {
				// Reset the path to serve index.html
				r.URL.Path = "/"
			}
		} else {
			// Don't forget to close the file
			f.Close()
		}

		// Serve either the requested file or index.html
		fileServer.ServeHTTP(w, r)
	})
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
