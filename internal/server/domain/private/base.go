package private

import "GophKeeper/internal/server/domain"

type BasePrivateSchema struct {
	domain.BaseEntity
	Owner int64
	Meta  string
}
