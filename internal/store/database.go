package store

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
)

type Store struct {
	db *sqlx.DB
}

func NewStore() (*Store, error) {
	// Opening a driver typically will not attempt to connect to the database.
	db, err := sqlx.Open("postgres", "dbname=sf6 user=admin password=admin sslmode=disable")
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		logrus.Fatal(err)
	}
	// Confirm a successful connection.
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)
	// Execute database migration scripts
	s := &Store{
		db: db,
	}
	// Execute database migration scripts
	if err = s.migrateDatabase(); err != nil {
		return nil, err
	}
	return s, err
}
