package dto

import "GophKeeper/proto/compiled/pb"

type BankCard struct {
	Meta   []byte
	Number []byte
	Date   []byte
	Cvv    []byte
}

func NewBankCard(dto *pb.BankCardDto) BankCard {
	return BankCard{
		Meta:   dto.GetMeta(),
		Number: dto.GetNumber(),
		Date:   dto.GetDate(),
		Cvv:    dto.GetCvv(),
	}
}

func NewPbBankCard(dto BankCard) *pb.BankCardDto {
	return &pb.BankCardDto{
		Meta:   dto.Meta,
		Number: dto.Number,
		Date:   dto.Date,
		Cvv:    dto.Cvv,
	}
}

func NewBankCardList(list []*pb.BankCardDto) []BankCard {
	r := make([]BankCard, 0)

	for _, e := range list {
		r = append(r, NewBankCard(e))
	}

	return r
}

func NewPbBankCardList(list []BankCard) []*pb.BankCardDto {
	r := make([]*pb.BankCardDto, 0)

	for _, e := range list {
		r = append(r, NewPbBankCard(e))
	}

	return r
}
