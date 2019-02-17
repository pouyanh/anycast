package mysql

import (
	"database/sql"

	"github.com/pouyanh/anycast/lib/actor"
)

func NewRelationalSqlDatabase(dsn string) (actor.RelationalDatabase, error) {
	return sql.Open("mysql", dsn)
}
