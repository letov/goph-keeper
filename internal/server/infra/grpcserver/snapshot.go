package grpcserver

import (
	"GophKeeper/internal/common/dto"
	"GophKeeper/proto/compiled/pb"
	"context"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (m *Grpc) SaveSnapshot(ctx context.Context, data *pb.SnapshotDto) (*pb.Response, error) {
	owner, ok := ctx.Value("owner").(string)
	if !ok {
		return nil, status.Error(codes.FailedPrecondition, "owner id missing")
	}

	s := dto.NewSnapshot(data)
	i, _ := strconv.ParseInt(owner, 10, 32)
	err := m.sr.SaveSnapshot(ctx, s, int32(i))
	if err != nil {
		return nil, err
	}
	return &pb.Response{Status: 1}, nil
}

func (m *Grpc) GetSnapshot(ctx context.Context, _ *emptypb.Empty) (*pb.SnapshotDto, error) {
	s, ok := ctx.Value("owner").(string)
	owner, err := strconv.ParseInt(s, 10, 32)
	if !ok || err != nil {
		return nil, status.Error(codes.FailedPrecondition, "owner id missing")
	}

	data, err := m.sr.GetSnapshot(ctx, int32(owner))
	if err != nil {
		return nil, err
	}

	return dto.NewPdSnapshot(data), nil
}
