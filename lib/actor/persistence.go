package actor

import "database/sql"

// Key/Value Storage
type Dictionary interface {
	Put(key string, value interface{}) error
	Has(key string) bool
	Get(key string) interface{}
	Delete(key string) error
	GetDefault(key string, value interface{}) interface{}
}

type RelationalDatabase interface {
	Begin() (*sql.Tx, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

func LimitClause(limit, offset uint) (string, []interface{}) {
	if offset + limit > 0 {
		if 0 == limit {
			return "LIMIT ?, 18446744073709551615", []interface{} {offset}
		} else if offset == 0 {
			return "LIMIT ?", []interface{} {limit}
		} else {
			return "LIMIT ?, ?", []interface{} {offset, limit}
		}
	}

	return "", nil
}
