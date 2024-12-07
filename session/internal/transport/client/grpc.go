package client

import (
	"context"

	v1 "session/pkg/api/grpc/v1"

	"github.com/Markuysa/pkg/tracer"
	"google.golang.org/grpc"
)

type sessionTransport struct {
	service service
	v1.UnimplementedSessionServiceServer
}

func New(s service) *sessionTransport {
	return &sessionTransport{service: s}
}

func (s *sessionTransport) RegisterServer(server *grpc.Server) {
	v1.RegisterSessionServiceServer(server, s)
}

func (s *sessionTransport) ClientAuth(ctx context.Context, req *v1.ClientAuthRequest) (*v1.ClientAuthResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	return s.service.ClientAuthorize(ctx, req)
}

func (s *sessionTransport) ClientSetSession(ctx context.Context, req *v1.ClientSetSessionRequest) (*v1.ClientSetSessionResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	return s.service.CreateSession(ctx, req)
}
