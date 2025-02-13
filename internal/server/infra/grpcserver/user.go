package grpcserver

import (
	"GophKeeper/internal/server/app/dto"
	"GophKeeper/proto/compiled/pb"
	"context"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *Grpc) SaveUser(ctx context.Context, data *pb.SaveUserDto) (*pb.Response, error) {
	d := dto.NewSaveUser(data)
	err := m.ur.SaveUser(ctx, d)
	return &pb.Response{Status: 1}, err
}

func (m *Grpc) LoginUser(ctx context.Context, data *pb.LoginUserDto) (*pb.UserResponse, error) {
	d := dto.NewLoginUser(data)
	id, err := m.ur.LoginUser(ctx, d)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "wrong login/password")
	}
	if id > 0 {
		jwt, e := m.auth.IssueToken(strconv.FormatInt(int64(id), 10))
		if e != nil {
			return nil, e
		}
		return &pb.UserResponse{
			Jwt: jwt,
		}, nil
	}
	return nil, status.Error(codes.Unauthenticated, "wrong login/password")
}
