package test

import (
	dto2 "GophKeeper/internal/common/dto"
	"GophKeeper/internal/server/app/repo"
	"GophKeeper/internal/server/infra/db"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Snapshot(t *testing.T) {
	type args struct {
		SaveUserDto dto2.SaveUser
		Snapshot    dto2.Snapshot
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "save/get snapshot test",
			args: args{
				SaveUserDto: dto2.SaveUser{
					Email:    "some@email.com",
					PassHash: "pass_hash",
				},
				Snapshot: dto2.Snapshot{
					LoginPasswordList: []dto2.LoginPassword{
						{
							Meta:     []byte("vk"),
							Login:    []byte("some@mail.com"),
							Password: []byte("secret_password"),
						},
						{
							Meta:     []byte("yandex"),
							Login:    []byte("some@mail.com"),
							Password: []byte("secret_password"),
						},
						{
							Meta:     []byte("google"),
							Login:    []byte("some@mail.com"),
							Password: []byte("secret_password"),
						},
					},
					BinaryList: []dto2.Binary{
						{
							Meta:   []byte("file1"),
							Binary: []byte("data1"),
						},
						{
							Meta:   []byte("file2"),
							Binary: []byte("data2"),
						},
						{
							Meta:   []byte("file3"),
							Binary: []byte("data3"),
						},
					},
					BankCardList: []dto2.BankCard{
						{
							Meta:   []byte("sber"),
							Number: []byte("4444 3333 2222 1111"),
							Date:   []byte("10.10.2030"),
							Cvv:    []byte("111"),
						},
						{
							Meta:   []byte("tbank"),
							Number: []byte("4444 3333 2222 1111"),
							Date:   []byte("10.10.2030"),
							Cvv:    []byte("111"),
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
				owner, _ := ur.LoginUser(ctx, dto2.LoginUser{
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
