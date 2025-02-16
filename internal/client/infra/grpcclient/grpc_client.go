package grpcclient

import (
	"GophKeeper/internal/client/infra/config"
	"GophKeeper/proto/compiled/pb"
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Grpc struct {
	Client pb.GophKeeperServiceClient
}

func NewGrpcClient(
	lc fx.Lifecycle,
	log zap.SugaredLogger,
	c config.Config,
) *Grpc {
	conn, err := grpc.NewClient(c.GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})

	client := pb.NewGophKeeperServiceClient(conn)
	return &Grpc{Client: client}
}
