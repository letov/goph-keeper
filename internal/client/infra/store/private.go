package store

import (
	"database/sql"
	"strings"
)

type PrivateType = string

const (
	LoginPassword PrivateType = "login_password"
	Binary        PrivateType = "binary"
	BankCard      PrivateType = "bank_card"
)

func deletePrivates(tx *sql.Tx) error {
	query := strings.Join([]string{
		"DELETE FROM privates;",
		"DELETE FROM key_values;",
		"DELETE FROM sqlite_sequence WHERE name='privates'",
	}, "\n")
	_, err := tx.Exec(query)
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
	tp PrivateType,
) (*sql.Rows, error) {
	query := strings.Join([]string{
		"SELECT kv.private_id as private_id, kv.type as type, kv.value as value",
		"FROM privates p",
		"LEFT JOIN key_values kv ON (p.id = kv.private_id)",
		"WHERE p.type = ?",
	}, " ")

	return r.db.Query(query, tp)
}

func savePrivateList(
	tx *sql.Tx,
	pt PrivateType,
	kvms []KeyValueMap,
) error {
	for _, kvm := range kvms {
		query := `INSERT INTO privates (type) VALUES (?)`
		_, err := tx.Exec(query, pt)
		if err != nil {
			return err
		}

		if err = saveKeyValueList(tx, kvm); err != nil {
			return err
		}
	}
	return nil
}
