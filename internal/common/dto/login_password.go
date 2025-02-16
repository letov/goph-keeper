package dto

import "GophKeeper/proto/compiled/pb"

type LoginPassword struct {
	Meta     []byte
	Login    []byte
	Password []byte
}

func NewLoginPassword(dto *pb.LoginPasswordDto) LoginPassword {
	return LoginPassword{
		Meta:     dto.GetMeta(),
		Login:    dto.GetLogin(),
		Password: dto.GetPassword(),
	}
}

func NewPbLoginPassword(dto LoginPassword) *pb.LoginPasswordDto {
	return &pb.LoginPasswordDto{
		Meta:     dto.Meta,
		Login:    dto.Login,
		Password: dto.Password,
	}
}

func NewLoginPasswordList(list []*pb.LoginPasswordDto) []LoginPassword {
	r := make([]LoginPassword, 0)

	for _, e := range list {
		r = append(r, NewLoginPassword(e))
	}

	return r
}

func NewPbLoginPasswordList(list []LoginPassword) []*pb.LoginPasswordDto {
	r := make([]*pb.LoginPasswordDto, 0)

	for _, e := range list {
		r = append(r, NewPbLoginPassword(e))
	}

	return r
}
