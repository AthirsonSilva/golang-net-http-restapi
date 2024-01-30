package database

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Holds the database connection pool
type Database struct {
	SQL *sql.DB
}

var dbConn = &Database{}

const (
	maxOpenDbConn = 10
	maxDbLifeTime = 5 * time.Minute
	maxIdleDbConn = 5
)

// Connect to database
func ConnectSQL(dsn string) (*Database, error) {
	db, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifeTime)
	dbConn.SQL = db

	err = testConnection(db)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func testConnection(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// Create new database connection
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Close database connection
func (db *Database) CloseDatabaseConnection() {
	dbConn.SQL.Close()
}
