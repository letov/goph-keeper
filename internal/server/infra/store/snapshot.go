package store

import (
	"GophKeeper/internal/server/app/dto"
	"context"
)

func (r *RepoDB) GetSnapshot(ctx context.Context, owner int32) (dto.Snapshot, error) {
	l, err := r.getLoginPasswordList(ctx, owner)
	if err != nil {
		return dto.Snapshot{}, err
	}

	b, err := r.getBinaryList(ctx, owner)
	if err != nil {
		return dto.Snapshot{}, err
	}

	bc, err := r.getBankCardList(ctx, owner)
	if err != nil {
		return dto.Snapshot{}, err
	}

	return dto.Snapshot{
		LoginPasswordList: l,
		BinaryList:        b,
		BankCardList:      bc,
	}, nil
}

func (r *RepoDB) SaveSnapshot(ctx context.Context, s dto.Snapshot, owner int32) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}

	if err = deletePrivates(ctx, tx, owner); err != nil {
		return err
	}

	if err = updateLoginPasswordList(ctx, tx, s.LoginPasswordList, owner); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	if err = updateBinaryList(ctx, tx, s.BinaryList, owner); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	if err = updateBankCardList(ctx, tx, s.BankCardList, owner); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}
