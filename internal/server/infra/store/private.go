package store

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5"
)

type PrivateType = string

const (
	LoginPassword PrivateType = "login_password"
	Binary        PrivateType = "binary"
	BankCard      PrivateType = "bank_card"
)

func deletePrivates(ctx context.Context, tx pgx.Tx, owner int32) error {
	query := `DELETE FROM privates WHERE owner = @owner`
	args := pgx.NamedArgs{"owner": owner}
	_, err := tx.Exec(ctx, query, args)
	return err
}

func mapToArr[T interface{}](m map[int32]T) []T {
	l := make([]T, 0)
	for _, v := range m {
		l = append(l, v)
	}
	return l
}

func (r *RepoDB) getPrivateList(
	ctx context.Context,
	tp PrivateType,
	owner int32,
) (pgx.Rows, error) {
	query := strings.Join([]string{
		"SELECT kv.private_id as private_id, kv.type as type, kv.value as value",
		"FROM privates p",
		"LEFT JOIN key_values kv ON (p.id = kv.private_id)",
		"WHERE owner = @owner AND p.type = @type",
	}, " ")
	args := pgx.NamedArgs{
		"owner": owner,
		"type":  tp,
	}

	return r.pool.Query(ctx, query, args)
}

func savePrivateList(
	ctx context.Context,
	tx pgx.Tx,
	pt PrivateType,
	kvms []KeyValueMap,
	owner int32,
) error {
	for _, kvm := range kvms {
		query := `INSERT INTO privates (owner, type) VALUES (@owner, @type)`
		args := pgx.NamedArgs{
			"owner": owner,
			"type":  pt,
		}
		_, err := tx.Exec(ctx, query, args)
		if err != nil {
			return err
		}

		if err = saveKeyValueList(ctx, tx, kvm); err != nil {
			return err
		}
	}
	return nil
}
