package test

import (
	"GophKeeper/internal/server/app/dto"
	"GophKeeper/internal/server/app/repo"
	"GophKeeper/internal/server/infra/db"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Snapshot(t *testing.T) {
	type args struct {
		SaveUserDto dto.SaveUser
		Snapshot    dto.Snapshot
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "save/get snapshot test",
			args: args{
				SaveUserDto: dto.SaveUser{
					Email:     "Email",
					PassHash:  "PassHash",
					PublicKey: "PublicKey",
				},
				Snapshot: dto.Snapshot{
					LoginPasswordList: []dto.LoginPassword{
						{
							Meta:     []byte("Meta"),
							Login:    []byte("Login"),
							Password: []byte("Password"),
						},
					},
					BinaryList: []dto.Binary{
						{
							Binary: []byte("Text"),
						},
						{
							Binary: []byte("Text"),
						},
						{
							Binary: []byte("Text"),
						},
					},
					BankCardList: []dto.BankCard{
						{
							Meta:   []byte("Meta"),
							Number: []byte("Login"),
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest(t, func(ur repo.User, sr repo.Snapshot, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)

				_ = ur.SaveUser(ctx, tt.args.SaveUserDto)
				owner, _ := ur.LoginUser(ctx, dto.LoginUser{
					Email:    tt.args.SaveUserDto.Email,
					PassHash: tt.args.SaveUserDto.PassHash,
				})
				_ = sr.SaveSnapshot(ctx, tt.args.Snapshot, owner)
				s, _ := sr.GetSnapshot(ctx, owner)
				assert.Equal(
					t,
					s.LoginPasswordList[0].Meta,
					tt.args.Snapshot.LoginPasswordList[0].Meta,
				)
			})
		})
	}
}
