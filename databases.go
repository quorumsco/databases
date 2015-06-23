package databases

import (
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func InitSQLX(engine, source string) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	if db, err = sqlx.Connect(engine, source); err != nil {
		return nil, err
	}

	db.Ping()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	return db, nil
}

func InitGORM(engine, source string) (*gorm.DB, error) {
	db, err := gorm.Open(engine, source)
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(false)

	return &db, nil
}
