package store

import (
	"database/sql"
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

func saveKeyValueList(tx *sql.Tx, kvm KeyValueMap) error {
	for k, v := range kvm {
		query := `INSERT INTO key_values (value, type, private_id) VALUES (?, ?, (select seq from sqlite_sequence where name='privates'))`
		_, err := tx.Exec(query, v, k)
		if err != nil {
			return err
		}
	}

	return nil
}
