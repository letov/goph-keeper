package handler

import (
	"GophKeeper/internal/client/app/session"
	"GophKeeper/internal/client/infra/grpcclient"
	"GophKeeper/internal/common/dto"
	"context"
	"fmt"
	"time"
)

type Register struct {
	gc *grpcclient.Grpc
	s  *session.Session
}

func (l *Register) GetHandler(
	ctx context.Context,
	errorCallback func(errorMsg string),
) func(email string, password string) error {
	return func(email string, password string) error {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer func() {
			cancel()
		}()

		_, err := l.gc.Client.SaveUser(ctxWithTimeout, dto.NewPbSaveUser(email, password))
		if err != nil {
			errorCallback(fmt.Sprintf("Register failed!\n%s", err))
			return err
		}
		return nil
	}
}

func NewRegister(
	gc *grpcclient.Grpc,
	s *session.Session,
) *Register {
	return &Register{
		gc,
		s,
	}
}
