package store

import (
	"GophKeeper/internal/common/dto"
	"database/sql"
)

func (r *RepoDB) getBinaryList() ([]dto.Binary, error) {
	rows, err := r.getPrivateList(Binary)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		id    int32
		tp    KeyValueType
		value []byte
	)
	m := make(map[int32]dto.Binary)
	for rows.Next() {
		err = rows.Scan(&id, &tp, &value)
		if err != nil {
			return nil, err
		}
		d, ok := m[id]
		if !ok {
			m[id] = dto.Binary{}
		}
		switch tp {
		case Meta:
			d.Meta = value
		case BinaryValue:
			d.Binary = value
		}
		m[id] = d
	}

	return mapToArr(m), nil
}

func updateBinaryList(tx *sql.Tx, l []dto.Binary) error {
	kvms := make([]KeyValueMap, 0)
	for _, e := range l {
		kvms = append(kvms, KeyValueMap{
			Meta:        e.Meta,
			BinaryValue: e.Binary,
		})
	}

	if err := savePrivateList(tx, Binary, kvms); err != nil {
		return err
	}

	return nil
}
