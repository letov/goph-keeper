package app

import (
	"GophKeeper/internal/client/app/view"
	"GophKeeper/internal/client/infra/db"
)

func Start(db *db.DB, _ *view.Root) {
	print(db)
}
