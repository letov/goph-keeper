package repo

import (
	"GophKeeper/internal/common/dto"
	"context"
)

type Snapshot interface {
	GetSnapshot(ctx context.Context, owner int32) (dto.Snapshot, error)
	SaveSnapshot(ctx context.Context, s dto.Snapshot, owner int32) error
}
