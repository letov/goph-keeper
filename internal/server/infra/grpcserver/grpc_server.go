package grpcserver

import (
	"GophKeeper/internal/server/app/repo"
	"GophKeeper/internal/server/infra/auth"
	"GophKeeper/internal/server/infra/config"
	"GophKeeper/proto/compiled/pb"
	"context"
	"net"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Grpc struct {
	ur   repo.User
	sr   repo.Snapshot
	auth *auth.Service
	pb.UnimplementedGophKeeperServiceServer
}

func NewGrpcServer(
	lc fx.Lifecycle,
	log zap.SugaredLogger,
	c config.Config,
	ur repo.User,
	sr repo.Snapshot,
	ai *AuthInterceptor,
	auth *auth.Service,
) *grpc.Server {
	srv := grpc.NewServer(grpc.UnaryInterceptor(ai.UnaryAuthMiddleware))
	grpcServer := &Grpc{ur: ur, sr: sr, auth: auth}
	pb.RegisterGophKeeperServiceServer(srv, grpcServer)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			addr := c.GrpcAddress
			ln, err := net.Listen("tcp", addr)
			if err != nil {
				return err
			}
			log.Info("Starting GRPC server: ", addr)
			go func() {
				err := srv.Serve(ln)
				if err != nil {
					log.Error(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Stop()
			return nil
		},
	})

	return srv
}
