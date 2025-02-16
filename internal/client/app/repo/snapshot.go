package repo

import (
	"GophKeeper/internal/common/dto"
)

type Snapshot interface {
	GetSnapshot() (dto.Snapshot, error)
	SaveSnapshot(s dto.Snapshot) error
	ClearDB() error
}
