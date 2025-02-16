package dto

import "GophKeeper/proto/compiled/pb"

type SaveUser struct {
	Email    string
	PassHash string
}

func NewSaveUser(dto *pb.SaveUserDto) SaveUser {
	return SaveUser{
		Email:    dto.GetEmail(),
		PassHash: dto.GetPassHash(),
	}
}

func NewPbSaveUser(email string, password string) *pb.SaveUserDto {
	return &pb.SaveUserDto{
		Email:    email,
		PassHash: password,
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

func NewPbLoginUser(email string, password string) *pb.LoginUserDto {
	return &pb.LoginUserDto{
		Email:    email,
		PassHash: password,
	}
}
