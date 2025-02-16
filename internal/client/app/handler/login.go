package handler

import (
	"GophKeeper/internal/client/app/session"
	"GophKeeper/internal/client/infra/grpcclient"
	"GophKeeper/internal/common/dto"
	"context"
	"fmt"
	"time"
)

type Login struct {
	gc *grpcclient.Grpc
	s  *session.Session
}

func (l *Login) GetHandler(
	ctx context.Context,
	errorCallback func(errorMsg string),
) func(email string, password string) error {
	return func(email string, password string) error {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer func() {
			cancel()
		}()

		r, err := l.gc.Client.LoginUser(ctxWithTimeout, dto.NewPbLoginUser(email, password))
		if err != nil {
			errorCallback(fmt.Sprintf("Login failed!\n%s", err))
			return err
		}
		l.s.SetPassword(password)
		l.s.SetToken(r.GetJwt())
		return nil
	}
}

func NewLogin(
	gc *grpcclient.Grpc,
	s *session.Session,
) *Login {
	return &Login{
		gc,
		s,
	}
}
