package repo

import (
	"time"

	"github.com/jackc/pgx"

)

var (
	_pool *pgx.ConnPool
)

func Init() {
	config := pgx.ConnPoolConfig{}
	config.AcquireTimeout = 2 * time.Second
	config.Database = "yaya"
	config.Host = "192.168.31.60"
	config.Port = 5433
	config.MaxConnections = 5
	config.User = "meidomx"
	config.Password = "meidomx"
	pool, err := pgx.NewConnPool(config)
	if err != nil {
		panic(err)
	}
	_pool = pool
}
