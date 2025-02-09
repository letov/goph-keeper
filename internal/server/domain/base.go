package domain

import "time"

type BaseEntity struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
