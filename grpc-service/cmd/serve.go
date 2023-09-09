package cmd

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/stevenweathers/go-templates/grpc-service/db"
	authorsv1 "github.com/stevenweathers/go-templates/grpc-service/gen/go/authors/v1"
	"github.com/stevenweathers/go-templates/grpc-service/internal/authors"

	"github.com/stevenweathers/go-templates/grpc-service/internal/echo"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cobra"
	swaggerui "github.com/stevenweathers/go-templates/grpc-service/third_party"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	echov1 "github.com/stevenweathers/go-templates/grpc-service/gen/go/echo/v1"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches the example webserver on https://localhost:10000",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

// grpcHandlerFunc returns a http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(tamird): point to merged gRPC code rather than a PR.
		// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
func serveSwagger(mux *http.ServeMux) {
	fs := swaggerui.New()
	fileServer := http.FileServer(fs)
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func serve() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dbConn, dbErr := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if dbErr != nil {
		slog.Error(fmt.Sprintf("Unable to connect to database: %v", dbErr))
		os.Exit(1)
	}
	defer dbConn.Close(context.Background())
	authorQueries := db.New(dbConn)

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewClientTLSFromCert(demoCertPool, "localhost:10000")),
	}

	grpcServer := grpc.NewServer(opts...)
	echov1.RegisterEchoServiceServer(grpcServer, echo.NewServer())
	authorsv1.RegisterAuthorsServiceServer(grpcServer, authors.NewServer(authorQueries))
	ctx := context.Background()

	dcreds := credentials.NewTLS(&tls.Config{
		ServerName: demoAddr,
		RootCAs:    demoCertPool,
	})
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}

	mux := http.NewServeMux()
	// serve the echo swagger json
	mux.HandleFunc("/swagger-ui/echo.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/openapiv2/echo/v1/echo.swagger.json")
	})

	gwmux := runtime.NewServeMux()
	err := echov1.RegisterEchoServiceHandlerFromEndpoint(ctx, gwmux, demoAddr, dopts)
	if err != nil {
		slog.Error(fmt.Sprintf("serve: %v", err))
		return
	}

	mux.Handle("/", gwmux)
	serveSwagger(mux)

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    demoAddr,
		Handler: grpcHandlerFunc(grpcServer, mux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*demoKeyPair},
			NextProtos:   []string{"h2"},
		},
	}

	slog.Info("starting grpc and http service", slog.Int("port", port))
	err = srv.Serve(tls.NewListener(conn, srv.TLSConfig))

	if err != nil {
		slog.Error("ListenAndServe: ", err)
		panic(err)
	}

	return
}
