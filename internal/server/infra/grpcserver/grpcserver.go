package grpcserver

import (
	"GophKeeper/internal/server/app/dto"
	"GophKeeper/internal/server/app/repo"
	"GophKeeper/internal/server/infra/config"
	"GophKeeper/proto/compiled/pb"
	"context"
	"net"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Grpc struct {
	ur repo.User
	pb.UnimplementedGophKeeperServiceServer
}

func (m *Grpc) SaveUser(ctx context.Context, data *pb.SaveUserDto) (*pb.Empty, error) {
	d := dto.NewSaveUser(data)
	err := m.ur.Save(ctx, d)
	return &pb.Empty{}, err
}

func (m *Grpc) LoginUser(ctx context.Context, data *pb.LoginUserDto) (*pb.LoginResponse, error) {
	d := dto.NewLoginUser(data)
	status, err := m.ur.Login(ctx, d)
	return &pb.LoginResponse{Status: status}, err
}

func NewGrpcServer(
	lc fx.Lifecycle,
	log zap.SugaredLogger,
	c config.Config,
	ur repo.User,
) *grpc.Server {
	srv := grpc.NewServer()
	grpcServer := &Grpc{ur: ur}
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
