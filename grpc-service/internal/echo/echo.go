package echo

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	echov1 "github.com/stevenweathers/go-templates/grpc-service/gen/go/echo/v1"
)

type Service struct {
	echov1.EchoServiceServer
}

func (es *Service) Echo(ctx context.Context, message *echov1.EchoMessage) (*echov1.EchoMessage, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info(fmt.Sprintf("rpc request Echo(%q)\n", message.Value))
	return message, nil
}

func NewServer() echov1.EchoServiceServer {
	return new(Service)
}
