// Connects to databases and return the client
package databases

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/quorumsco/logs"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// timeout defines the sleep duration between each connection retry.
const timeout = 5 * time.Second

// retry defines the number of times to retry.
const retry = 3

// Return a sqlx client
func InitSQLX(dialect, args string) (*sqlx.DB, error) {
	var (
		db  *sqlx.DB
		err error
		i   int
	)

retry:
	for {
		db, err = sqlx.Connect(dialect, args)
		switch {
		case err == nil:
			break retry
		case i >= retry:
			return nil, err
		default:
			logs.Error(err)
			i++
		}
		time.Sleep(timeout)
	}

	if db.Ping() != nil {
		return db, err
	}

	return db, nil
}

/*
InitGORM returns a gorm client.

It wraps gorm.Open(dialect, args ...interface{}) in order to allows retrying connection on failure. Example:

	db, err := InitGORM("postgres", "postgres://user:password@/dbname?charset=utf8&parseTime=True&loc=Local")

Number of retries depends on the `retry` package constant, separated by a `timeout` time constant.
*/
func InitGORM(dialect, args string) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
		i   int
	)

retry:
	for {
		db, err = gorm.Open(dialect, args)
		switch {
		case err == nil:
			break retry
		case i >= retry:
			return nil, err
		default:
			logs.Error(err)
			i++
		}
		time.Sleep(timeout)
	}

	if db.DB().Ping() != nil {
		return db, err
	}

	return db, nil
}
