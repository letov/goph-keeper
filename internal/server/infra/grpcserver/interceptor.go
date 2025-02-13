package grpcserver

import (
	"GophKeeper/internal/server/infra/auth"
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	auth *auth.Service
}

func NewAuthInterceptor(auth *auth.Service) *AuthInterceptor {
	return &AuthInterceptor{auth: auth}
}

func (i *AuthInterceptor) UnaryAuthMiddleware(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	var ignoreMethod = []string{
		"/proto.GophKeeperService/SaveUser",
		"/proto.GophKeeperService/LoginUser",
	}
	method, _ := grpc.Method(ctx)
	for _, i := range ignoreMethod {
		if method == i {
			return handler(ctx, req)
		}
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	token := md["authorization"]
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "authorization token is not provided")
	}

	trimed := strings.TrimPrefix(token[0], "Bearer ")
	userID, err := i.auth.ValidateToken(trimed)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("invalid token: %v", err))
	}

	ctx = context.WithValue(ctx, "owner", userID)

	return handler(ctx, req)
}
