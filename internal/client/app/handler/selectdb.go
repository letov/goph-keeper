package handler

import (
	"GophKeeper/internal/client/app/repo"
	"GophKeeper/internal/client/app/session"
	"GophKeeper/internal/client/infra/grpcclient"
	"GophKeeper/internal/common/dto"
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	clearDBAction = 0
	loadRemoteDB  = 1
	saveRemoteDB  = 2
)

type SelectDB struct {
	gc *grpcclient.Grpc
	s  *session.Session
	rs repo.Snapshot
}

func (l *SelectDB) GetHandler(
	ctx context.Context,
	errorCallback func(errorMsg string),
) func(action int) error {
	return func(action int) error {
		switch action {
		case clearDBAction:
			if err := l.rs.ClearDB(); err != nil {
				errorCallback(fmt.Sprintf("Clear DB failed!\n%s", err))
				return err
			}
			return nil
		case loadRemoteDB:
			token := l.s.GetToken()
			md := metadata.Pairs("authorization", token)
			ctx = metadata.NewOutgoingContext(ctx, md)

			ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer func() {
				cancel()
			}()

			s, err := l.gc.Client.GetSnapshot(ctxWithTimeout, &emptypb.Empty{})
			if err != nil {
				errorCallback(fmt.Sprintf("Get remote DB failed!\n%s", err))
				return err
			}
			if err = l.rs.SaveSnapshot(dto.NewSnapshot(s)); err != nil {
				errorCallback(fmt.Sprintf("Get remote DB failed!\n%s", err))
				return err
			}
			return nil
		case saveRemoteDB:
			token := l.s.GetToken()
			md := metadata.Pairs("authorization", token)
			ctx = metadata.NewOutgoingContext(ctx, md)

			ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer func() {
				cancel()
			}()

			s, err := l.rs.GetSnapshot()
			if err != nil {
				errorCallback(fmt.Sprintf("Save to remote DB failed!\n%s", err))
				return err
			}
			if _, err = l.gc.Client.SaveSnapshot(ctxWithTimeout, dto.NewPdSnapshot(s)); err != nil {
				errorCallback(fmt.Sprintf("Save to remote DB failed!\n%s", err))
				return err
			}
			return nil
		}

		return errors.New("unhandled action")
	}
}

func NewSelectDB(
	gc *grpcclient.Grpc,
	s *session.Session,
	rs repo.Snapshot,
) *SelectDB {
	return &SelectDB{
		gc,
		s,
		rs,
	}
}
