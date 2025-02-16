package store

import (
	"GophKeeper/internal/common/dto"
)

func (r *RepoDB) GetSnapshot() (dto.Snapshot, error) {
	l, err := r.getLoginPasswordList()
	if err != nil {
		return dto.Snapshot{}, err
	}

	b, err := r.getBinaryList()
	if err != nil {
		return dto.Snapshot{}, err
	}

	bc, err := r.getBankCardList()
	if err != nil {
		return dto.Snapshot{}, err
	}

	return dto.Snapshot{
		LoginPasswordList: l,
		BinaryList:        b,
		BankCardList:      bc,
	}, nil
}

func (r *RepoDB) SaveSnapshot(s dto.Snapshot) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if err = deletePrivates(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = updateLoginPasswordList(tx, s.LoginPasswordList); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = updateBinaryList(tx, s.BinaryList); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = updateBankCardList(tx, s.BankCardList); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *RepoDB) ClearDB() error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if err = deletePrivates(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
