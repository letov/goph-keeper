package repo

import (
	"GophKeeper/internal/server/app/dto"
	"context"
)

type User interface {
	Save(ctx context.Context, u dto.SaveUser) error
	Login(ctx context.Context, l dto.LoginUser) (bool, error)
}
