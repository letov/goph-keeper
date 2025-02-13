package store

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type KeyValueType = string

const (
	Meta        KeyValueType = "meta"
	Login       KeyValueType = "login"
	Password    KeyValueType = "password"
	BinaryValue KeyValueType = "binary"
	Number      KeyValueType = "number"
	Date        KeyValueType = "date"
	Cvv         KeyValueType = "cvv"
)

type KeyValueMap = map[KeyValueType][]byte

func saveKeyValueList(ctx context.Context, tx pgx.Tx, kvm KeyValueMap) error {
	for k, v := range kvm {
		query := `INSERT INTO key_values (value, type, private_id) VALUES (@value, @type, currval('privates_id_seq'))`
		args := pgx.NamedArgs{
			"value": v,
			"type":  k,
		}
		_, err := tx.Exec(ctx, query, args)
		if err != nil {
			return err
		}
	}

	return nil
}
