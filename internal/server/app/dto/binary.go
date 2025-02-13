package dto

import "GophKeeper/proto/compiled/pb"

type Binary struct {
	Meta   []byte
	Binary []byte
}

func NewBinary(dto *pb.BinaryDto) Binary {
	return Binary{
		Meta:   dto.GetMeta(),
		Binary: dto.GetBinary(),
	}
}

func NewPbBinary(dto Binary) *pb.BinaryDto {
	return &pb.BinaryDto{
		Meta:   dto.Meta,
		Binary: dto.Binary,
	}
}

func NewBinaryList(list []*pb.BinaryDto) []Binary {
	r := make([]Binary, 0)

	for _, e := range list {
		r = append(r, NewBinary(e))
	}

	return r
}

func NewPbBinaryList(list []Binary) []*pb.BinaryDto {
	r := make([]*pb.BinaryDto, 0)

	for _, e := range list {
		r = append(r, NewPbBinary(e))
	}

	return r
}
