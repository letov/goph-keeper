package test

import (
	"GophKeeper/internal/common/dto"
	"GophKeeper/internal/server/app/repo"
	"GophKeeper/internal/server/infra/db"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_User(t *testing.T) {
	type args struct {
		SaveUserDto dto.SaveUser
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "save/login user test",
			args: args{
				SaveUserDto: dto.SaveUser{
					Email:    "some@email.com",
					PassHash: "pass_hash",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest(t, func(ur repo.User, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)

				_ = ur.SaveUser(ctx, tt.args.SaveUserDto)
				id, _ := ur.LoginUser(ctx, dto.LoginUser{
					Email:    tt.args.SaveUserDto.Email,
					PassHash: tt.args.SaveUserDto.PassHash,
				})
				assert.Greater(t, id, int32(0))

				res, _ := ur.LoginUser(ctx, dto.LoginUser{
					Email:    tt.args.SaveUserDto.Email,
					PassHash: tt.args.SaveUserDto.PassHash,
				})
				assert.Greater(t, res, int32(0))
			})
		})
	}
}
