package store

import (
	"GophKeeper/internal/common/dto"
	"database/sql"
)

func (r *RepoDB) getLoginPasswordList() ([]dto.LoginPassword, error) {
	rows, err := r.getPrivateList(LoginPassword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		id    int32
		tp    KeyValueType
		value []byte
	)
	m := make(map[int32]dto.LoginPassword)
	for rows.Next() {
		err = rows.Scan(&id, &tp, &value)
		if err != nil {
			return nil, err
		}
		d, ok := m[id]
		if !ok {
			m[id] = dto.LoginPassword{}
		}
		switch tp {
		case Meta:
			d.Meta = value
		case Login:
			d.Login = value
		case Password:
			d.Password = value
		}
		m[id] = d
	}

	return mapToArr(m), nil
}

func updateLoginPasswordList(tx *sql.Tx, l []dto.LoginPassword) error {
	kvms := make([]KeyValueMap, 0)
	for _, e := range l {
		kvms = append(kvms, KeyValueMap{
			Meta:     e.Meta,
			Login:    e.Login,
			Password: e.Password,
		})
	}

	if err := savePrivateList(tx, LoginPassword, kvms); err != nil {
		return err
	}

	return nil
}
