package database

import (
	"database/sql"
	"log"
)

func NewDB(conn string) (*sql.DB, func(), error) {
	db, err := sql.Open("mssql", conn)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Print(err)
		}
	}
	return db, cleanup, nil
}
