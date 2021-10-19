package storage

import (
	"database/sql"
	"fmt"
)

type dbRepo struct {
	session *sql.DB
}

func NewStorage(session *sql.DB) StorageI {
	return &dbRepo{
		session: session,
	}
}

func (d *dbRepo) Hello(message string) string {
	return fmt.Sprintf("How are you %s. My name is Abrorbek", message)
}
