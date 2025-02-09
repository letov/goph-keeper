package app

import (
	"GophKeeper/internal/server/app/dto"
	"GophKeeper/internal/server/app/repo"
	"context"
)

func Start(ur repo.User) {
	_, _ = ur.Login(context.Background(), dto.LoginUser{
		Email:    "asd",
		PassHash: "sdfd",
	})
}
