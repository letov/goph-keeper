package repo

import (
	"GophKeeper/internal/common/dto"
	"context"
)

type User interface {
	SaveUser(ctx context.Context, u dto.SaveUser) error
	LoginUser(ctx context.Context, l dto.LoginUser) (int32, error)
}
