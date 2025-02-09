package test

import (
	"GophKeeper/internal/server/app/dto"
	"GophKeeper/internal/server/app/repo"
	"GophKeeper/internal/server/infra/db"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SaveUser(t *testing.T) {
	type args struct {
		Email     string
		PassHash  string
		PublicKey string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "save/login user test",
			args: args{
				Email:     "email",
				PassHash:  "password_hash",
				PublicKey: "public_key",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest(t, func(ur repo.User, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)

				_ = ur.Save(ctx, dto.SaveUser(tt.args))
				res, _ := ur.Login(ctx, dto.LoginUser{
					Email:    tt.args.Email,
					PassHash: tt.args.PassHash,
				})
				assert.True(t, res)
			})
		})
	}
}
