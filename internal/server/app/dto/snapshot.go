package dto

import "GophKeeper/proto/compiled/pb"

type Snapshot struct {
	LoginPasswordList []LoginPassword
	BinaryList        []Binary
	BankCardList      []BankCard
}

func NewSnapshot(dto *pb.SnapshotDto) Snapshot {
	return Snapshot{
		LoginPasswordList: NewLoginPasswordList(dto.GetLoginPasswordList()),
		BinaryList:        NewBinaryList(dto.GetBinaryList()),
		BankCardList:      NewBankCardList(dto.GetBankCardList()),
	}
}

func NewPdSnapshot(dto Snapshot) *pb.SnapshotDto {
	return &pb.SnapshotDto{
		LoginPasswordList: NewPbLoginPasswordList(dto.LoginPasswordList),
		BinaryList:        NewPbBinaryList(dto.BinaryList),
		BankCardList:      NewPbBankCardList(dto.BankCardList),
	}
}
