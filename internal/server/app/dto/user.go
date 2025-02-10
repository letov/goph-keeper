package dto

import "GophKeeper/proto/compiled/pb"

type SaveUser struct {
	Email     string
	PassHash  string
	PublicKey string
}

func NewSaveUser(dto *pb.SaveUserDto) SaveUser {
	return SaveUser{
		Email:     dto.GetEmail(),
		PassHash:  dto.GetPassHash(),
		PublicKey: dto.GetPublicKey(),
	}
}

type LoginUser struct {
	Email    string
	PassHash string
}

func NewLoginUser(dto *pb.LoginUserDto) LoginUser {
	return LoginUser{
		Email:    dto.GetEmail(),
		PassHash: dto.GetPassHash(),
	}
}
