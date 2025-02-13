package store

import (
	"GophKeeper/internal/server/infra/db"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type RepoDB struct {
	pool *pgxpool.Pool
	log  zap.SugaredLogger
}

//func getSeqStr(cnt int64) string {
//	if cnt <= 0 {
//		panic("Incorrect cnt value")
//	}
//	res := make([]string, 0)
//	s := "currval('key_values_id_seq') - %v"
//	for i := cnt - 1; i >= 0; i-- {
//		res = append(res, fmt.Sprintf(s, i))
//	}
//	return strings.Join(res, ",")
//}

func NewRepoDB(db *db.DB, log zap.SugaredLogger) *RepoDB {
	return &RepoDB{
		db.GetPool(),
		log,
	}
}
