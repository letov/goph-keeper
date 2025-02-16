package store

import (
	"GophKeeper/internal/common/dto"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepoDB) getBankCardList(ctx context.Context, owner int32) ([]dto.BankCard, error) {
	rows, err := r.getPrivateList(ctx, BankCard, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		id    int32
		tp    KeyValueType
		value []byte
	)
	m := make(map[int32]dto.BankCard)
	for rows.Next() {
		err = rows.Scan(&id, &tp, &value)
		if err != nil {
			return nil, err
		}
		d, ok := m[id]
		if !ok {
			m[id] = dto.BankCard{}
		}
		switch tp {
		case Meta:
			d.Meta = value
		case Number:
			d.Number = value
		case Date:
			d.Date = value
		case Cvv:
			d.Cvv = value
		}
		m[id] = d
	}

	return mapToArr(m), nil
}

func updateBankCardList(ctx context.Context, tx pgx.Tx, l []dto.BankCard, owner int32) error {
	kvms := make([]KeyValueMap, 0)
	for _, e := range l {
		kvms = append(kvms, KeyValueMap{
			Meta:   e.Meta,
			Number: e.Number,
			Date:   e.Date,
			Cvv:    e.Cvv,
		})
	}

	if err := savePrivateList(ctx, tx, BankCard, kvms, owner); err != nil {
		return err
	}

	return nil
}
