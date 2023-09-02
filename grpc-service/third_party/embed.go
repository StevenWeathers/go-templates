package third_party

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed swagger-ui
var swagui embed.FS

func New() http.FileSystem {
	fsys, err := fs.Sub(swagui, "swagger-ui")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
